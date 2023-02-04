package main

import (
	_ "assignment/Api"
	"bufio"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-delve/delve/service/api"
)

func main() {

outer:
	for {
		fmt.Println("====================")
		fmt.Println("Purchase History\n",
			"1.Get All Purchase History\n",
			"2.View all purchase History\n",
			"9.Quit\n")
		fmt.Println("====================")

		fmt.Print("Enter an option: ")
		var choice string
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(input)

		switch choice {
		case "1":
			fmt.Println("launching case 1")
		case "2":
			fmt.Println("launching case 2")
		case "9":
			fmt.Println("You are now exiting the app")
			break outer
		}
	}
}
