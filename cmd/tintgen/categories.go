// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"context"
	"encoding/json/jsontext"
	"encoding/json/v2"
	"errors"
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
	"strings"

	jev "github.com/Gaurav-Gosain/jev-go"
)

const (
	categoriesFileName = "categories.json"
	categoriesSchema   = "./categories.schema.json"

	// toneUncategorized is cached and offered to the classifier. It is not a
	// public Tone: an empty Tone means the same thing.
	toneUncategorized = "uncategorized"

	toneQuestionID     = "tone"
	toneMinProbability = 0.5
	toneConcurrency    = 4

	jevModel          = "jev-1.14"
	openRouterBaseURL = "https://openrouter.ai/api"

	toneInstructions = "Classify the aesthetic tone of this terminal color theme. Judge the hex colors. Treat the theme name as a hint only."

	uncategorizedCriteria = "None of the tones fit."
)

// tone is one aesthetic bucket. Order is the public Tone order and the
// categories.json key order. Criteria is the classifier description.
type tone struct {
	Name     string
	Criteria string
}

// tones is the tone catalog. Adding a tone is one entry here, then regenerate.
var tones = []tone{
	{Name: "muted", Criteria: "Grayed, low saturation. Muted is grayed, while pastel is light and chalky."},
	{Name: "pastel", Criteria: "Soft and chalky. Pastel is light and chalky, while muted is grayed."},
	{Name: "vivid", Criteria: "Strongly saturated, still a normal theme. Vivid is merely saturated, while neon is fluorescent."},
	{Name: "neon", Criteria: "Fluorescent, glow, or synthwave. Neon is fluorescent, while vivid is merely saturated."},
	{Name: "retro", Criteria: "Vintage hardware: CRT phosphor, amber terminal, CGA, or Apple II. Retro beats monochrome for hardware palettes."},
	{Name: "monochrome", Criteria: "One ink color, and not a vintage-hardware palette. Retro beats monochrome for hardware palettes."},
	{Name: "high-contrast", Criteria: "The palette's purpose is stark foreground/background separation. Use this only when that separation is the point of the theme."},
}

func toneNames() []string {
	names := make([]string, len(tones))
	for i, spec := range tones {
		names[i] = spec.Name
	}
	return names
}

func cacheToneNames() []string {
	return append(toneNames(), toneUncategorized)
}

func toneCriteria() map[string]string {
	out := make(map[string]string, len(tones)+1)
	for _, spec := range tones {
		out[spec.Name] = spec.Criteria
	}
	out[toneUncategorized] = uncategorizedCriteria
	return out
}

func knownTone(name string) bool {
	return slices.Contains(cacheToneNames(), name)
}

// toneFile is the committed cache. Hue is not stored: it is computed from the
// palette on every generate.
type toneFile struct {
	Schema string    `json:"$schema"`
	Tone   toneLists `json:"tone"`
}

// toneLists is the tone object in categories.json. Keys are written in
// catalog order, including buckets that are empty.
type toneLists map[string][]string

type themeState struct {
	Name    string       `json:"name"`
	Author  string       `json:"author,omitempty"`
	Dark    bool         `json:"dark"`
	Palette themePalette `json:"palette"`
}

type themePalette struct {
	Foreground    string `json:"foreground,omitempty"`
	Background    string `json:"background,omitempty"`
	Selection     string `json:"selection,omitempty"`
	SelectionText string `json:"selection_text,omitempty"`
	Bold          string `json:"bold,omitempty"`
	Cursor        string `json:"cursor,omitempty"`
	CursorText    string `json:"cursor_text,omitempty"`
	Black         string `json:"black,omitempty"`
	Red           string `json:"red,omitempty"`
	Green         string `json:"green,omitempty"`
	Yellow        string `json:"yellow,omitempty"`
	Blue          string `json:"blue,omitempty"`
	Purple        string `json:"purple,omitempty"`
	Cyan          string `json:"cyan,omitempty"`
	White         string `json:"white,omitempty"`
	BrightBlack   string `json:"brightBlack,omitempty"`
	BrightRed     string `json:"brightRed,omitempty"`
	BrightGreen   string `json:"brightGreen,omitempty"`
	BrightYellow  string `json:"brightYellow,omitempty"`
	BrightBlue    string `json:"brightBlue,omitempty"`
	BrightPurple  string `json:"brightPurple,omitempty"`
	BrightCyan    string `json:"brightCyan,omitempty"`
	BrightWhite   string `json:"brightWhite,omitempty"`
}

// toneOutcome is one completed classification. ok is false when the request
// failed or the response was malformed, which leaves the name out of the cache
// so the next run retries it.
type toneOutcome struct {
	bucket string
	ok     bool
}

// classifyTones loads the tone cache, classifies themes that are not cached,
// and rewrites the cache. The returned map is the tone to emit: uncategorized
// and never-classified themes are absent. A non-nil error means at least one
// theme must be retried; themes that did get a normal answer are still cached.
func classifyTones(ctx context.Context, themes []Tint, path, token string, reclassify bool) (map[string]string, error) {
	token = strings.TrimSpace(token)
	if reclassify && token == "" {
		return nil, errors.New("reclassify requires an openrouter token")
	}

	cached, err := loadToneCache(path)
	if err != nil {
		return nil, err
	}

	names := make([]string, len(themes))
	for i := range themes {
		names[i] = themes[i].Name
	}

	pending := uncachedThemes(themes, cached, reclassify)
	var outcomes map[string]toneOutcome
	asked := false
	if len(pending) > 0 && token != "" {
		asked = true
		var askErr error
		outcomes, askErr = askTones(ctx, token, pending)
		if askErr != nil {
			return nil, askErr
		}
	} else if len(pending) > 0 {
		logger.InfoContext(ctx, "skipping tone classification", "missing", len(pending))
	}

	kept, failed := resolveTones(names, cached, reclassify, asked, outcomes)
	if writeErr := writeToneCache(path, kept); writeErr != nil {
		return nil, writeErr
	}
	if len(failed) > 0 {
		logger.WarnContext(ctx, "tone classification failed", "count", len(failed))
		return generatedTones(kept), fmt.Errorf("tone classification failed for %d themes", len(failed))
	}
	return generatedTones(kept), nil
}

func uncachedThemes(themes []Tint, cached map[string]string, reclassify bool) []Tint {
	pending := make([]Tint, 0, len(themes))
	for i := range themes {
		if !reclassify {
			if _, ok := cached[themes[i].Name]; ok {
				continue
			}
		}
		pending = append(pending, themes[i])
	}
	sort.Slice(pending, func(i, j int) bool {
		return pending[i].Name < pending[j].Name
	})
	return pending
}

// resolveTones merges a tone cache with classification outcomes.
// Names absent from current are dropped. When reclassify is set, cached
// buckets are ignored. When asked is false, names missing from the cache stay
// missing. When asked is true, a failed or missing outcome is omitted so the
// next run retries it.
func resolveTones(current []string, cached map[string]string, reclassify, asked bool, outcomes map[string]toneOutcome) (kept map[string]string, failed []string) {
	kept = make(map[string]string, len(current))
	if !reclassify {
		currentSet := make(map[string]struct{}, len(current))
		for _, name := range current {
			currentSet[name] = struct{}{}
		}
		for name, bucket := range cached {
			if _, ok := currentSet[name]; ok {
				kept[name] = bucket
			}
		}
	}

	if !asked {
		return kept, nil
	}

	failed = make([]string, 0)
	for _, name := range current {
		if _, ok := kept[name]; ok {
			continue
		}
		outcome, ok := outcomes[name]
		if !ok || !outcome.ok {
			failed = append(failed, name)
			continue
		}
		kept[name] = outcome.bucket
	}
	sort.Strings(failed)
	return kept, failed
}

func generatedTones(kept map[string]string) map[string]string {
	out := make(map[string]string, len(kept))
	for name, bucket := range kept {
		if bucket == toneUncategorized {
			continue
		}
		out[name] = bucket
	}
	return out
}

func loadToneCache(path string) (map[string]string, error) {
	body, err := os.ReadFile(path) //nolint:gosec // Path is the generator output directory.
	if errors.Is(err, os.ErrNotExist) {
		return map[string]string{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("read tone cache: %w", err)
	}

	var file toneFile
	if err = json.Unmarshal(body, &file, json.RejectUnknownMembers(true)); err != nil {
		return nil, fmt.Errorf("parse tone cache: %w", err)
	}
	assigned, err := file.Tone.assignments()
	if err != nil {
		return nil, fmt.Errorf("parse tone cache: %w", err)
	}
	return assigned, nil
}

// assignments maps a theme name to its tone. A name listed under two tones
// keeps the earlier catalog entry. Unknown keys are an error.
func (l toneLists) assignments() (map[string]string, error) {
	for bucket := range l {
		if !knownTone(bucket) {
			return nil, fmt.Errorf("unknown tone %q", bucket)
		}
	}
	out := map[string]string{}
	for _, bucket := range cacheToneNames() {
		for _, name := range l[bucket] {
			if _, exists := out[name]; exists {
				continue
			}
			out[name] = bucket
		}
	}
	return out, nil
}

func (l toneLists) MarshalJSONTo(enc *jsontext.Encoder) error {
	if err := enc.WriteToken(jsontext.BeginObject); err != nil {
		return err
	}
	for _, bucket := range cacheToneNames() {
		names := l[bucket]
		if names == nil {
			names = []string{}
		}
		if err := enc.WriteToken(jsontext.String(bucket)); err != nil {
			return err
		}
		if err := json.MarshalEncode(enc, names); err != nil {
			return err
		}
	}
	return enc.WriteToken(jsontext.EndObject)
}

func writeToneCache(path string, assignments map[string]string) error {
	lists := make(toneLists, len(cacheToneNames()))
	for _, bucket := range cacheToneNames() {
		lists[bucket] = []string{}
	}
	for name, bucket := range assignments {
		names, ok := lists[bucket]
		if !ok {
			return fmt.Errorf("unknown tone %q for %s", bucket, name)
		}
		lists[bucket] = append(names, name)
	}
	for bucket, names := range lists {
		sort.Strings(names)
		lists[bucket] = names
	}

	body, err := json.Marshal(toneFile{Schema: categoriesSchema, Tone: lists}, jsontext.WithIndent("  "))
	if err != nil {
		return fmt.Errorf("encode tone cache: %w", err)
	}
	if len(body) == 0 || body[len(body)-1] != '\n' {
		body = append(body, '\n')
	}
	if err = os.WriteFile(path, body, 0o644); err != nil { //nolint:gosec // categories.json is repository source.
		return fmt.Errorf("write tone cache: %w", err)
	}
	return nil
}

func askTones(ctx context.Context, token string, themes []Tint) (map[string]toneOutcome, error) {
	client, err := jev.New(
		jev.WithAPIKey(token),
		jev.WithBaseURL(openRouterBaseURL),
		jev.WithModel(jevModel),
	)
	if err != nil {
		return nil, fmt.Errorf("jev client: %w", err)
	}

	results, stats := jev.Batch(ctx, client, themes, toneRequest, jev.BatchOptions{
		Concurrency: toneConcurrency,
		OnProgress: func(done, total int) {
			if done%25 == 0 || done == total {
				logger.InfoContext(ctx, "classified theme tones", "done", done, "total", total)
			}
		},
	})
	logger.InfoContext(ctx, "tone classification finished",
		"ok", stats.Succeeded,
		"failed", stats.Failed,
		"input_tokens", stats.InputTokens,
	)

	outcomes := make(map[string]toneOutcome, len(results))
	var first error
	var malformed int
	for i := range results {
		if results[i].Err != nil && first == nil {
			first = results[i].Err
		}
		outcome := outcomeFromResult(results[i])
		if results[i].Err == nil && !outcome.ok {
			malformed++
		}
		outcomes[results[i].Item.Name] = outcome
	}
	if first != nil {
		logger.WarnContext(ctx, "tone request failed", "error", first, "failed", stats.Failed)
	}
	if malformed > 0 {
		logger.WarnContext(ctx, "tone response malformed", "count", malformed)
	}
	return outcomes, nil
}

func toneRequest(t Tint) (any, jev.Questions) {
	return newThemeState(t), jev.Questions{
		toneQuestionID: jev.Choice{
			Instructions: toneInstructions,
			Criteria:     toneCriteria(),
		},
	}
}

func newThemeState(t Tint) themeState {
	return themeState{
		Name:   t.Name,
		Author: t.Author,
		Dark:   t.IsDark(),
		Palette: themePalette{
			Foreground:    t.Fg,
			Background:    t.Bg,
			Selection:     t.SelectionBg,
			SelectionText: t.SelectionText,
			Bold:          t.Bold,
			Cursor:        t.Cursor,
			CursorText:    t.CursorText,
			Black:         t.Black,
			Red:           t.Red,
			Green:         t.Green,
			Yellow:        t.Yellow,
			Blue:          t.Blue,
			Purple:        t.Purple,
			Cyan:          t.Cyan,
			White:         t.White,
			BrightBlack:   t.BrightBlack,
			BrightRed:     t.BrightRed,
			BrightGreen:   t.BrightGreen,
			BrightYellow:  t.BrightYellow,
			BrightBlue:    t.BrightBlue,
			BrightPurple:  t.BrightPurple,
			BrightCyan:    t.BrightCyan,
			BrightWhite:   t.BrightWhite,
		},
	}
}

func outcomeFromResult(result jev.Result[Tint]) toneOutcome {
	if result.Err != nil || result.Response == nil {
		return toneOutcome{}
	}
	choice, err := result.Response.Answers.Choice(toneQuestionID)
	if err != nil {
		return toneOutcome{}
	}
	bucket, ok := bucketFromChoice(choice.Selected, choice.Probabilities)
	if !ok {
		return toneOutcome{}
	}
	return toneOutcome{bucket: bucket, ok: true}
}

// bucketFromChoice applies the tone policy to one winning option.
// uncategorized is kept. A catalog tone below toneMinProbability is stored
// as uncategorized. Anything outside the catalog is malformed.
func bucketFromChoice(selected string, probabilities map[string]float64) (string, bool) {
	probability, ok := probabilities[selected]
	if !ok || selected == "" || math.IsNaN(probability) || probability < 0 || probability > 1 {
		return "", false
	}
	if selected == toneUncategorized {
		return toneUncategorized, true
	}
	if !knownTone(selected) {
		return "", false
	}
	if probability < toneMinProbability {
		return toneUncategorized, true
	}
	return selected, true
}
