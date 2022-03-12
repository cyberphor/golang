package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("powershell")

	stdin, err := cmd.StdinPipe()

	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		fmt.Fprintln(stdin, "Import-Module soap")
		fmt.Fprintln(stdin, "$FilterHashTable = @{ LogName = 'Microsoft-Windows-Sysmon/Operational' }")
		fmt.Fprintln(stdin, "Get-WinEvent -FilterHashTable $FilterHashTable -MaxEvents 1 | Read-WinEvent")
	}()

	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", output)
}
