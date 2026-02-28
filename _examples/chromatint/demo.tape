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

Hide
Type "go run *.go"
Sleep 500ms
Enter
Show
Sleep 1s
Right@600ms 90
Sleep 1s
