package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"ag/agent"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	a := agent.New()

	fmt.Println("🚀 ag (Auto Gemma v5)")

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "/exit" {
			break
		}

		fmt.Println(a.Handle(input))
	}
}
