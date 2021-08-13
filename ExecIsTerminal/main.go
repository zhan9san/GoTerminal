package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("../IsTerminal/IsTerminal")
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Starting command")
	cmd.Start()
	fmt.Println("out:", outb.String())
	fmt.Println("err:", errb.String())
}
