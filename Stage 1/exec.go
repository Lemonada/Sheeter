package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

func execute(method string, command string) (string, string, error) {
	// Doing some clean up, this is wrong and only a temporary fix....
	newCommand := strings.Replace(command, "\\\"", "\"", -1)
	newCommand = strings.Replace(newCommand, "\\\\", "\\", -1)
	newCommand = "(" + newCommand + ")"
	cmd := exec.Command("cmd.exe", "/C", newCommand)
	if method == "powershell" {
		cmd = exec.Command("powershell.exe", "/C", newCommand)
	} else if method != "cmd" {
		return "", "Failed to decide how to run", fmt.Errorf("Failed to decide how to run")
	}
	cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000}
	var out bytes.Buffer
	var outErr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &outErr
	cmd.Start()

	// Use a channel to signal completion so we can use a select statement
	done := make(chan error)
	go func() { done <- cmd.Wait() }()

	// Start a timer
	timeout := time.After(commandTimeout)

	select {
	case <-timeout:
		// Timeout happened first, kill the process and print a message.
		cmd.Process.Kill()
		return "Command timed out", " ", nil
	case err := <-done:
		// Command completed before timeout. Print output and error if it exists.
		if err != nil {
			return "", outErr.String(), err
		}
		return out.String(), " ", nil
	}
}
