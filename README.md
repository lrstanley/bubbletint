<!-- template:define:options
{
  "nodescription": true
}
-->
<img title="Logo" src="./_examples/_images/logo.png" width="961">

<!-- template:begin:header -->
<!-- do not edit anything in this "template" block, its auto-generated -->

<p align="center">
  <a href="https://github.com/lrstanley/bubbletint/releases">
    <img title="Release Downloads" src="https://img.shields.io/github/downloads/lrstanley/bubbletint/total?style=flat-square">
  </a>
  <a href="https://github.com/lrstanley/bubbletint/tags">
    <img title="Latest Semver Tag" src="https://img.shields.io/github/v/tag/lrstanley/bubbletint?style=flat-square">
  </a>
  <a href="https://github.com/lrstanley/bubbletint/commits/master">
    <img title="Last commit" src="https://img.shields.io/github/last-commit/lrstanley/bubbletint?style=flat-square">
  </a>



  <a href="https://github.com/lrstanley/bubbletint/actions?query=workflow%3Atest+event%3Apush">
    <img title="GitHub Workflow Status (test @ master)" src="https://img.shields.io/github/actions/workflow/status/lrstanley/bubbletint/test.yml?branch=master&label=test&style=flat-square">
  </a>




  <a href="https://codecov.io/gh/lrstanley/bubbletint">
    <img title="Code Coverage" src="https://img.shields.io/codecov/c/github/lrstanley/bubbletint/master?style=flat-square">
  </a>

  <a href="https://pkg.go.dev/github.com/lrstanley/bubbletint/v2">
    <img title="Go Documentation" src="https://pkg.go.dev/badge/github.com/lrstanley/bubbletint/v2?style=flat-square">
  </a>
</p>
<p align="center">
  <a href="https://github.com/lrstanley/bubbletint/issues?q=is:open+is:issue+label:bug">
    <img title="Bug reports" src="https://img.shields.io/github/issues/lrstanley/bubbletint/bug?label=issues&style=flat-square">
  </a>
  <a href="https://github.com/lrstanley/bubbletint/issues?q=is:open+is:issue+label:enhancement">
    <img title="Feature requests" src="https://img.shields.io/github/issues/lrstanley/bubbletint/enhancement?label=feature%20requests&style=flat-square">
  </a>
  <a href="https://github.com/lrstanley/bubbletint/pulls">
    <img title="Open Pull Requests" src="https://img.shields.io/github/issues-pr/lrstanley/bubbletint?label=prs&style=flat-square">
  </a>
  <a href="https://github.com/lrstanley/bubbletint/releases">
    <img title="Latest Semver Release" src="https://img.shields.io/github/v/release/lrstanley/bubbletint?style=flat-square">
    <img title="Latest Release Date" src="https://img.shields.io/github/release-date/lrstanley/bubbletint?label=date&style=flat-square">
  </a>
  <a href="https://github.com/lrstanley/bubbletint/discussions/new?category=q-a">
    <img title="Ask a Question" src="https://img.shields.io/badge/support-ask_a_question!-blue?style=flat-square">
  </a>
  <a href="https://liam.sh/chat"><img src="https://img.shields.io/badge/discord-bytecord-blue.svg?style=flat-square" title="Discord Chat"></a>
</p>
<!-- template:end:header -->

<!-- template:begin:toc -->
<!-- do not edit anything in this "template" block, its auto-generated -->
## :link: Table of Contents

  - [Features](#sparkles-features)
  - [Usage](#gear-usage)
    - [Global registry](#global-registry)
    - [Custom Registry](#custom-registry)
    - [Filter by hue and tone](#filter-by-hue-and-tone)
    - [Use specific tints](#use-specific-tints)
    - [React to changes](#react-to-changes)
    - [Follow the terminal's theme](#follow-the-terminals-theme)
    - [Mix colors and check contrast](#mix-colors-and-check-contrast)
    - [Semantic roles](#semantic-roles)
  - [Examples](#clap-examples)
    - [Package manager example](#package-manager-example)
    - [Complex example](#complex-example)
    - [Loading tints from a file](#loading-tints-from-a-file)
    - [Chroma Support](#chroma-support)
    - [Glamour Support](#glamour-support)
    - [Semantic roles example](#semantic-roles-example)
    - [Themed app example](#themed-app-example)
    - [Follow the terminal](#follow-the-terminal)
  - [Changes in v2](#rocket-changes-in-v2)
    - [No More Bloated Interfaces](#no-more-bloated-interfaces)
    - [Chroma Support](#chroma-support-1)
    - [JSON Serializable](#json-serializable)
  - [Support &amp; Assistance](#raising_hand_man-support--assistance)
  - [Contributing](#handshake-contributing)
  - [License](#balance_scale-license)
<!-- template:end:toc -->

## :sparkles: Features

- :heavy_check_mark: Utilize 770+ built-in tints. [Check them all out here!](https://lrstanley.github.io/bubbletint/)
  We're taking advantage of the great work others have done
  with [iTerm2-Color-Schemes](https://github.com/mbadolato/iTerm2-Color-Schemes),
  and making them more accessible in a programmatic fashion.
- :heavy_check_mark: Support for dynamically changing tints at runtime, to give
  immediate feedback on which tints might look the best for your application and
  terminal setup.
- :heavy_check_mark: Piecemeal support for specific tints. Really like one specific
  tint, and want to use that as the standard for your TUI? No problem!
- :heavy_check_mark: Filter built-in tints by hue (red, blue, and the rest of the
  spectrum) and tone (muted, retro, and the other aesthetic labels).
- :heavy_check_mark: Extendible with your own custom tints, in addition to marshalling
  and unmarshalling from JSON, if you wanted to support reading/writing tints to
  and from a file.
- :heavy_check_mark: Works with `lipgloss` and similar tools which support the
  `color.Color` interface from **stdlib**!
- :heavy_check_mark: Map a tint onto a [Glamour](https://github.com/charmbracelet/glamour)
  stylesheet through the `glamourtint` sub-package.
- :heavy_check_mark: Get notified when tints are added, replaced, removed, or when the
  current tint changes, with `Subscribe`.
- :heavy_check_mark: Follow the theme of the terminal your app runs in (kitty, Ghostty,
  Alacritty, foot, Windows Terminal, X resources, and iTerm2 color presets), on Linux,
  macOS and Windows, through the opt-in `sources` sub-package.
- :heavy_check_mark: Turn any tint into ready-to-use UI colors (panels, dialogs,
  buttons, selected rows, muted text, status pills) through the `roles` sub-package, so
  your components look right on every tint, without picking colors yourself.
- :heavy_check_mark: Perceptual color helpers with no dependencies: OKLab mixing, WCAG
  contrast, and contrast guards that keep a color's hue.

---

## :gear: Usage

```console
go get -u github.com/lrstanley/bubbletint/v2@latest
```

[Take a look at all tints here](https://lrstanley.github.io/bubbletint/).

BubbleTint supports three different modes of usage, as shown below:

### Global registry

```go
package main

import (
	// [...]
	tint "github.com/lrstanley/bubbletint/v2"
)


func main() {
	// [...]
	tint.NewDefaultRegistry()
	tint.SetTint(tint.TintDraculaPlus) // Set a specific tint.
	tint.SetTintID("dracula_plus") // Or by ID (this aligns with a Tint's ID field).

	// You can now use the current tint from the registry:
	style := lipgloss.NewStyle().
		SetForeground(tint.Current().Fg).
		Background(tint.Current().BrightGreen)
}
```

### Custom Registry

```go
package main

import (
	// [...]
	tint "github.com/lrstanley/bubbletint/v2"
)


func main() {
	// [...]
	theme := tint.NewRegistry(
		// Default tint.
		tint.TintDraculaPlus,
		// All others you want to register.
		tint.TintCatppuccinFrappe,
		tint.TintCatppuccinMocha,
		tint.TintDraculaPlus,
		tint.TintMonokaiPro,
		tint.TintTinaciousDesignDark,
	}

	theme.Register(tint.TintNord) // Register additional tints on the custom registry.
	theme.SetTint(tint.TintNord) // Can still set a tint on a custom registry.

	// Can also paginate through tints, using PreviousTint/NextTint
	theme.NextTint()

	// You can now use the current tint from the registry:
	style := lipgloss.NewStyle().
		SetForeground(theme.Current().Fg).
		Background(theme.Current().BrightGreen)
}
```

### Filter by hue and tone

Each built-in tint has a hue computed from its ANSI palette, and may have a
tone such as muted or retro. An empty tone means the tint is uncategorized or
has not been classified yet.

```go
package main

import (
	"fmt"

	tint "github.com/lrstanley/bubbletint/v2"
)

func main() {
	// A registry of one hue.
	blues := tint.NewRegistry(nil, tint.TintsByHue(tint.HueBlue)...)

	// Filter that registry by tone.
	for _, theme := range blues.TintsByTone(tint.ToneMuted) {
		fmt.Println(theme.DisplayName)
	}
}
```

### Use specific tints

```go
package main

import (
	// [...]
	tint "github.com/lrstanley/bubbletint/v2"
)

var (
	statusbarStyle = lipgloss.NewStyle().
		SetForeground(tint.TintDraculaPlus.Fg).
		Background(tint.TintDraculaPlus.BrightGreen)
)

// Or:
var (
	theme = tint.TintDraculaPlus
	statusbarStyle = lipgloss.NewStyle().
		SetForeground(theme.Fg).
		Background(theme.BrightGreen)
)
```

### React to changes

`Subscribe` calls your function whenever a registry changes: a tint being added,
replaced or removed, or the current tint changing. `CurrentChanged` is the one to
listen for if all you need to know is when to rebuild your styles. It is also sent
when the current tint is replaced in-place, and when the current tint is removed
and the registry falls back to another tint.

```go
cancel := tint.Subscribe(func(ev tint.Event) {
	if ev.Kind == tint.CurrentChanged {
		p.Send(rebuildStylesMsg{}) // Forward it to your Bubble Tea program.
	}
})
defer cancel()
```

Your function is called synchronously, on the goroutine that made the change, after
the registry's lock is released: it can safely call back into the registry, but it
should return quickly. Subscriptions belong to the registry they were made on, so if
you replace `tint.DefaultRegistry` after subscribing, subscribe again on the new one.

### Follow the terminal's theme

Tints can also come from outside of bubbletint, through a `tint.Source`, which is added
to a registry with `Attach`. Sources are opt-in, and live in their own module, so the
core stays free of dependencies:

```go
import (
	"github.com/lrstanley/bubbletint/sources/v2/terminal"
	tint "github.com/lrstanley/bubbletint/v2"
)

src, err := terminal.Detect() // The source for the terminal we're running in.
if err == nil {
	detach, err := tint.Attach(ctx, src)
	// [...]
}
```

- The source's tints are registered with their ID prefixed by the source's name, e.g.
  `kitty:current`, so they never collide with your own tints.
- By default the registry follows the source: `kitty:current` becomes the current tint,
  and stays up to date as the user changes their terminal's theme (subscribers are told
  about it). Use `tint.WithFollow(false)` to only register the tints.
- Check for `tint.ErrNotFound` (with `errors.Is`) to handle terminals that aren't
  supported, or have no config.
- Ghostty (`theme = light:X,dark:Y`) and Windows Terminal can name a theme for each of
  light and dark mode, and switch with the system's appearance. The source can't see which
  the terminal is in, but your app can, from the terminal's background: tell it with
  `src.SetDark(msg.IsDark())` on Bubble Tea's `tea.BackgroundColorMsg`, and the registry
  follows the other theme. Both are registered too, as `<source>:current_dark` and
  `<source>:current_light`.

| Package                                   | Reads                                                                 | Platforms                       |
| ----------------------------------------- | --------------------------------------------------------------------- | ------------------------------- |
| `sources/kitty`                           | `kitty.conf` and its includes (e.g. `current-theme.conf`)             | Linux, macOS, BSDs              |
| `sources/ghostty`                         | `config`, its `config-file` includes, and the `theme` it names        | Linux, macOS                    |
| `sources/alacritty`                       | `alacritty.toml` and its `import`s                                    | Linux, macOS, Windows, BSDs     |
| `sources/foot`                            | `foot.ini` and its `include`s                                         | Linux, BSDs                     |
| `sources/windowsterminal`                 | `settings.json` color schemes, and the active profile's scheme        | Windows (and WSL)               |
| `sources/xresources`                      | `~/.Xresources` (xterm, urxvt, st, ...)                               | Linux, BSDs                     |
| `sources/itermcolors`                     | `.itermcolors` files, or a directory of them                          | Anywhere (path must be given)   |
| `sources/terminal`                        | Picks one of the above, from the environment                          | Anywhere                        |

Config files are polled (once a second by default, see `sources.WithPollInterval`)
rather than watched with filesystem notifications, so watching behaves the same on every
platform, and keeps working with editors that replace files and with symlinked dotfiles.
Use `sources.WithPath` to read a config from somewhere other than the default locations.
Writing your own source only takes implementing `Name()` and `Load()` (and `Watch()`,
if it can tell when it changes).

... and that's it!

### Mix colors and check contrast

Beside `Darken` and `Lighten`, colors can be mixed in OKLab, where halfway looks
halfway and hue is kept, and checked or guarded for contrast:

```go
t := tint.Current()

hover := t.Bg.Mix(*t.Fg, 0.12)                // the background, 12% toward the text
ratio := t.BrightBlack.Contrast(*t.Bg)        // WCAG contrast ratio, 1 to 21
hint := t.BrightBlack.EnsureContrast(*t.Bg, 3) // lighter or darker, same hue, until it reads
```

The 16 ANSI colors can also be indexed by `tint.Slot`, with `t.ANSI(tint.SlotBrightRed)`
or `t.ANSI16()`.

### Semantic roles

A tint is a palette: a background, a text color, and 16 ANSI colors. A UI needs more than
that. It needs to know what to draw a dialog in, what a selected row looks like, how a
button changes under the pointer, and which gray is quiet enough for a hint but still
readable. Picking those from a tint by hand only works for the tint you picked them on.

The `roles` sub-package does it for you. You ask for a color by **what it's for** (its
role), such as `Raised` for a dialog, `Muted` for secondary text or `Danger` for an error,
and get a color that works on whatever tint is in use: dark, light, low contrast, or a
single hue.

**Use it when** you're building components or an app that should look right on any tint,
including ones your users pick or bring from their terminal. If you only want a tint's
own colors, use the tint directly.

```console
go get -u github.com/lrstanley/bubbletint/roles/v2@latest
```

**1. Derive once**, when the tint or the terminal changes, not on every render:

```go
// profile is the terminal's color support, such as Bubble Tea's tea.ColorProfileMsg.
f := roles.Derive(tint.Current()).Fit(roles.Target{Profile: profile})
```

**2. Ask for styles by role.** They're Lip Gloss styles, so they drop into what you have:

```go
title := f.Text(roles.Accent).Bold(true)
hint := f.Text(roles.Subtle)
errPill := f.Fill(roles.Danger)        // filled with danger, with text that reads on it
errTag := f.Fill(roles.Danger.Soft())  // a quieter, faded version
panel := f.Fill(roles.Surface)
```

**3. Draw on things.** `f.On(role)` gives you the roles for drawing *on* a surface, so text
and controls stay readable there:

```go
dialog := f.Fill(roles.Raised)
inDialog := f.On(roles.Raised)
note := inDialog.Text(roles.Muted)
cancel := inDialog.Fill(roles.Control) // a neutral button, a step off the dialog
input := inDialog.Fill(roles.Sunken)   // an input, sunk into the dialog
```

**4. Add states.** Any role can be hovered, pressed or selected:

```go
row := f.Fill(roles.Bg.Selected())
menuItem := f.Fill(roles.Overlay.Hover())
pressed := inDialog.Fill(roles.Control.Pressed())
```

What it takes care of, so you don't have to:

- **Readability.** Text, lines and status colors are kept readable on whatever they're
  drawn on, without losing the tint's hues.
- **Every tint.** The same code gives sensible surfaces and distinct status colors on dark,
  light and single-hue tints. Where hue can't tell success from danger, weight does.
- **Every terminal.** Colors are fitted to true color, 256, 16 or no colors. Neighboring
  surfaces and states stay apart, and where a color can't be drawn, an attribute
  (faint, underline, reverse) stands in for it.
- **Your own roles.** Add app-specific roles with `roles.Define`, and they're handled like
  the built-in ones.

For a larger app, keep all of this in one place: a theme package that's the only code
importing `roles`, and that gives components what they draw by intent (a panel, a selected
row, a danger button). The app holds one current theme and builds a new one when the
tint, a setting or the terminal changes. The [themed app example](#themed-app-example)
is laid out that way.

<details>
<summary>More detail</summary>

- The layers (`Bg`, `Surface`, `Raised`, `Overlay`) are the same everywhere. Text, lines,
  `Control`, `Sunken` and soft fills are worked out from what they're drawn on, so text on
  a light accent pill turns dark.
- Contrast guards move a color's lightness but keep its hue. Turn them off with
  `roles.WithoutGuards()` for a hand-made tint whose author wants exactly their colors.
- `f.Text` sets no background, for renderers that keep the background under text;
  `f.Span` adds it, for Lip Gloss strings rendered on their own and joined onto a fill.
  Outside Lip Gloss, `f.Color(role)` and `f.Attr(role)` are the color and what to add to
  it, and `f.Fit` fits a color of your own, such as a gradient's step.
- `f.Distinct(roles.Raised, roles.Bg)` reports whether two roles look different on the
  terminal, to choose a design by, such as outlining dialogs where their surface won't
  show.
- `roles.Lerp` crossfades between two fitted palettes, dark to light included.
- `roles.AccentChoice` and `roles.ColorLimit` parse the accent and color settings an app
  reads from its config file, with JSON schemas to match.
- Role names and the API follow semver; the exact colors derived may be tuned in minor
  releases.

</details>

## :clap: Examples

### Package manager example

- Uses a static theme with a specific palette for the entire application.
- [Example source](./_examples/package-manager/main.go)

![package-manager example](./_examples/package-manager/demo.webp)

### Complex example

- Uses the default registry (with all tints), and has keybinds to paginate through each tint.
- Note that only a portion of this example is styled (e.g. not a full background color, which
  may make sense for some of the provided tints).
- [Example source](./_examples/complex).

![complex example](./_examples/complex/demo.webp)

### Loading tints from a file

- Uses the `json` package to load a tint from a file.
- Unmarshalling supports both `RGBA` objects and hex strings.
  - Example: `{"r": 255, "g": 0, "b": 0, "a": 255}` or `"#ff0000"`
  - Hex strings would be easier for end users in most cases.
- Marshalling supports RGBA. This is to ensure that we don't lose the alpha channel
  of the color when marshalling to JSON. Alpha channels default to `255` with all
  of the default tints, but if you use custom tints, this will help keep that info.
- [Example source](./_examples/load-from-file).

![load-from-file example](./_examples/load-from-file/demo.webp)

### Chroma Support

- Uses the `chromatint` sub-package to render a code example as a `chroma.StyleEntries`,
  using the specified tint.
- [Example source](./_examples/chromatint).

![chromatint example](./_examples/chromatint/demo.webp)

### Glamour Support

- Uses the `glamourtint` sub-package to build a Glamour stylesheet from the current tint.
- Fenced code blocks use the same palette roles as `chromatint`.
- [Example source](./_examples/glamourtint).

![glamourtint example](./_examples/glamourtint/demo.webp)

```go
import (
	"charm.land/glamour/v2"
	"github.com/lrstanley/bubbletint/glamourtint/v2"
	tint "github.com/lrstanley/bubbletint/v2"
)

renderer, err := glamour.NewTermRenderer(
	glamour.WithStyles(glamourtint.StyleConfig(tint.Current())),
	glamour.WithWordWrap(80),
)
```

---

### Semantic roles example

- Uses the `roles` sub-package to style a small dashboard (panels, a selected row, a
  dialog with an input and buttons, status pills) by role, with no colors picked by hand.
  Press `←`/`→` to see it hold up on every tint, and `d` to see it on 256, 16 and no
  colors.
- [Example source](./_examples/roles).

### Themed app example

- Shows how to structure an app around `roles`: a `theme` package is the only place colors
  come from, components ask it for what they draw, and the app holds one current theme,
  rebuilding it when anything changes.
- The tint follows the terminal's background through a dark/light `tint.Pair`, and ←/→
  crossfades to another. Settings (accent, color limit, flat surfaces) load from an
  optional JSON config and change live from a menu (`s`). Rows and buttons respond to
  the mouse.
- [Example source](./_examples/roles-app).

### Follow the terminal

- Uses the `sources/terminal` sub-package to follow the theme of the terminal it runs in,
  and `Subscribe` to be told when it changes. Change your terminal's theme while it runs.
  It asks the terminal for its background, so a light/dark theme pair follows the mode
  the terminal is in.
- [Example source](./_examples/follow-terminal).

## :rocket: Changes in v2

> [!CAUTION]
> At a high-level, **there are more breaking changes with v2 of bubbletint, than
> not**. You will want to review the new API.

### Remove Reliance on Lipgloss for `color.Color`

Refactored the library as a whole to support the `color.Color` interface directly,
without relying on lipgloss, for more agnostic usage. It should now be possible to
use it with many other libraries without pulling in lipgloss (and other associated
dependencies), though you may still need down-sampling if you're not using
bubbletea/lipgloss, to support non-true-color (or 256bit) terminals.

- This also makes it easier to use with lipgloss v2, as less conversions need to happen now. Example:

  ```go
  theme := tint.TintDraculaPlus
  s := lipgloss.NewStyle().Foreground(theme.Fg)
  ```

- This should make things more performant.
- Should make hex codes more consistent (some themes from our datasources had **4**
  length hex colors, in addition to **7** length, but that should now be normalized as **7** length.
- **Reduces total external dependencies for the main library down to 1**.

### No More Bloated Interfaces

Removed the tint interface with a huge list of both global functions, as well as
methods for querying colors, and have instead opted for a more straightforward
struct for storing color values.

- This has lowered the overall amount of code generated.
- Callers of the library can still implement their own interfaces.
- The global registry still works the same as before, so you don't have to manage
  passing the theme state across many different components.
- Rather than `Registry.Blue()`, you'd do one of:

  ```go
  // Using custom registry. Done during initialization.
  registry := tint.NewRegistry(
      // Default tint.
      tint.TintDraculaPlus,
      // All others you want to register.
      tint.TintCatppuccinFrappe,
      tint.TintCatppuccinMocha,
      tint.TintDraculaPlus,
      tint.TintMonokaiPro,
      tint.TintTinaciousDesignDark,
  )

  // Use it.
  s := lipgloss.NewStyle().
      Foreground(registry.Current().BrightBlue)
  ```

  ```go
  // Using global registry. Done during initialization.
  tint.SetTint(tint.TintDraculaPlus)

  // Use it.
  s := lipgloss.NewStyle().
      Foreground(tint.Current().BrightBlue)
  ```

  ```go
  // Pick a specific tint:
  theme := tint.TintDraculaPlus // Set globally somewhere.

  // Use it.
  s := lipgloss.NewStyle().
      Foreground(theme.BrightBlue)
  ```

### Chroma Support

We now natively support [github.com/alecthomas/chroma](https://github.com/alecthomas/chroma),
through a new sub-package, `github.com/lrstanley/bubbletint/chromatint/v2`. See
[_examples/chromatint](_examples/chromatint). This allows natively mapping bubbletint's
tints to rendered versions of various programming/markup languages.

### JSON Serializable

Tints are now fully JSON serializable.

- When marshalled, they will be stored as RGBA (to not lose the alpha value, if your tints support that).
- When unmarshalled, we support both RGBA (`{"r": 255, "g": 0, "b": 0, "a": 255}`), in addition to hex strings (`"#ff0000"`), which makes it easier for end users to add their own tints.

### `panic()`'s OH MY!

The code now has more resilient fallbacks, reducing the potential chance that developer-error results in a panic.

---

<!-- template:begin:support -->
<!-- do not edit anything in this "template" block, its auto-generated -->
## :raising_hand_man: Support & Assistance

* :heart: Please review the [Code of Conduct](https://github.com/lrstanley/bubbletint?tab=coc-ov-file) for
     guidelines on ensuring everyone has the best experience interacting with
     the community.
* :raising_hand_man: Take a look at the [support](https://github.com/lrstanley/.github/blob/HEAD/.github/SUPPORT.md) document on
     guidelines for tips on how to ask the right questions.
* :lady_beetle: For all features/bugs/issues/questions/etc, [head over here](https://github.com/lrstanley/bubbletint/issues/new/choose).
<!-- template:end:support -->

<!-- template:begin:contributing -->
<!-- do not edit anything in this "template" block, its auto-generated -->
## :handshake: Contributing

* :heart: Please review the [Code of Conduct](https://github.com/lrstanley/bubbletint?tab=coc-ov-file) for guidelines
     on ensuring everyone has the best experience interacting with the
    community.
* :clipboard: Please review the [contributing](https://github.com/lrstanley/bubbletint?tab=contributing-ov-file) doc for submitting
     issues/a guide on submitting pull requests and helping out.
* :old_key: For anything security related, please review this repositories [security policy](https://github.com/lrstanley/bubbletint/security/policy).
<!-- template:end:contributing -->

<!-- template:begin:license -->
<!-- do not edit anything in this "template" block, its auto-generated -->
## :balance_scale: License

```
MIT License

Copyright (c) 2022 Liam Stanley <liam@liam.sh>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

_Also located [here](LICENSE)_
<!-- template:end:license -->
