package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {

	file, err := os.Open("dns.txt")
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		title(scanner.Text())
		cmd := exec.Command("nslookup", scanner.Text())
		cmdOutput := &bytes.Buffer{}
		cmd.Stdout = cmdOutput
		err1 := cmd.Run()
		if err1 != nil {
			os.Stderr.WriteString(err1.Error())
		}
		fmt.Print(string(cmdOutput.Bytes()))
		spaceLine()
	}

}

func title(serverName string) {
	fmt.Println("**********************")
	fmt.Println("****", serverName, "****")
	fmt.Println("**********************")

}

func spaceLine() {
	fmt.Println("")
	fmt.Println("")

}
