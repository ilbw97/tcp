package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("/bin/bash", "-c", "`df -lh`")

	//Create get command output pipeline
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error: can not obtain stdout pipe for command:%s\n", err)
		return
	}

	//Execute command
	if err := cmd.Start(); err != nil {
		fmt.Println("Error: The command is err,", err)
		return
	}
	//Using a buffered reader
	outputBuf := bufio.NewReader(stdout)

	for {
		//Get one line at a time, and check whether the current line has been read
		output, _, err := outputBuf.ReadLine()
		if err != nil {
			//Determine whether the end of the file is reached, otherwise an error will occur
			if err.Error() != "EOF" {
				fmt.Printf("Error :%s\n", err)
			}
			return
		}
		fmt.Printf("%s\n", string(output))
	}

	//The wait method blocks until the command to which it belongs runs completely
	if err := cmd.Wait(); err != nil {
		fmt.Println("wait:", err.Error())
		return
	}
}
