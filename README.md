# ClipBoardVirus

### Simple Virus Created with GOLang for fun

This dosnt allow you to copy. When you trying to paste it, it paste the text that you used in script

Test on Windows 10 only

- If you have installed GO on your machine
- Build project using `go build ClipBoardVirus.go `. This command will generate executable. But when when you run it it opens a window
- Use `go build -ldflags -H=windowsgui ClipBoardVirus.go ` to build executable. This executable will run without Window

#### Change the string that you wanted to paste by editing following variable value

`var str string = "String you want to paste"`

<img src="https://user-images.githubusercontent.com/75155192/200337939-1070ec0e-126e-415e-8524-bb1653b087c6.png"
     alt="Text to Change"/>
