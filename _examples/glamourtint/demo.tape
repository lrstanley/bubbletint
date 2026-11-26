# reel run demo.tape
Output demo.webp
Require go

Set Theme "Dracula+"
Set FontFamily "JetBrains Mono"
Set FontSize 20
Set Width 1600
Set Height 1040
Set Padding 24
Set Margin 56
Set MarginFill "135deg #674EFF #FF6B9D"
Set Shadow "0 20 48 #00000080"
Set BorderRadius 14
Set WindowBar Colorful
Set Title "bubbletint · glamourtint"
Set KeyOverlay true
Env PS1 "$ "

Hide
Type "go run ." Enter
Wait+Screen@2m "scroll • q: quit"
Show
Sleep 1s
Right@600ms 90
Sleep 1s
