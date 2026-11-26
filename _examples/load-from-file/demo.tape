# reel run demo.tape
Output demo.webp
Require go

Set Theme "Dracula+"
Set FontFamily "JetBrains Mono"
Set FontSize 18
Set Columns 136
Set Rows 17
Set Padding 24
Set Margin 40
Set MarginFill "135deg #674EFF #FF6B9D"
Set Shadow "0 20 48 #00000080"
Set BorderRadius 14
Set WindowBar Colorful
Set Title "bubbletint · load-from-file"
Set KeyOverlay true
Env PS1 "$ "

Hide
Type "go run ." Enter
Wait+Screen@2m "Type something here"
Show
Sleep 1s
Type "Hello from bubbletint!"
Enter
Type " ~ lrstanley"
Sleep 2s
