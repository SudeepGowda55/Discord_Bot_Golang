package main

import (
	"log"
	"os"
	"os/exec"
)

func commandExecutor(command string, cmd_args []string) (err error) {
	arg := cmd_args
	cmd_struct := exec.Command(command, arg...)
	cmd_struct.Stdout = os.Stdout
	cmd_struct.Stderr = os.Stderr
	cmd_struct.Stdin = os.Stdin

	err = cmd_struct.Run()

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func main() {
	commandExecutor("sudo", []string{"ifconfig", "wlan0", "down"})
	commandExecutor("sudo", []string{"ifconfig", "wlan0", "hw", "ether", "0a:eb:32:55:ef:f2"})
	commandExecutor("sudo", []string{"ifconfig", "wlan0", "up"})
}
