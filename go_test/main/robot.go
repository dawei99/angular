package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {

	inputReader:=bufio.NewReader(os.Stdin)
	fmt.Printf("请作答：")
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		name := input[:len(input)-1]
		switch name {
		case "":
			continue
		case "bye":
			fmt.Println("bye")
			os.Exit(1)
		default:
			fmt.Printf("继续说：")
		}

	}
}
