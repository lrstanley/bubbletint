# go run github.com/charmbracelet/vhs@latest demo.tape
Output demo.gif
Require go

Set Theme "Dracula+"
Set FontSize 20
Set Width 1600
Set Height 1040
Set Margin 40
Set Padding 30
Set MarginFill "#674EFF"
Set BorderRadius 14
Env PS1 "$ "

Type "go run *.go"
Sleep 500ms
Enter
Sleep 1s
Right@400ms 217
Sleep 1s
