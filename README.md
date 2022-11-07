# ClipBoardVirus

### Simple Virus Created with Golang for fun

This doesn't allow you to copy. When you trying to paste, it pastes the text that you used in script

Tested on Windows 10 only

- If you have installed GO on your machine
- Build project using `go build ClipBoardVirus.go `. This command will generate executable. But when you run it open a window
- Use `go build -ldflags -H=windowsgui ClipBoardVirus.go ` to build executable. This executable will run without a window

#### Change the string, that you wanted to paste by editing following variable value

`var str string = "String you want to paste"`

<img src="https://user-images.githubusercontent.com/75155192/200337939-1070ec0e-126e-415e-8524-bb1653b087c6.png"
     alt="Text to Change"/>

> If you want to close the programme. Just close the console window.
> If used windowless build then open task manager. Then search for name of exe file that you run and end the task.
> <img src="https://user-images.githubusercontent.com/75155192/200343500-9803b85a-d350-4abd-ab63-69b816500ea4.png" alt="End the Task"/>
