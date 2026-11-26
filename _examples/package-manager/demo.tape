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
Set Title "bubbletint · package-manager"
Set KeyOverlay false
Env PS1 "$ "

# build once while hidden, so the recording doesn't sit on the compile.
Hide
Type "go build -o /dev/null . && clear" Enter
Wait@2m /\$\s*$/
Show
Sleep 500ms
Type "go run ." Enter
Wait+Screen@1m "Done!"
Sleep 2s
