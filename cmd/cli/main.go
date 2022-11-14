package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/gookit/color"

	"github.com/youcan-jpn/janken_go_grpc/service"
)

func main() {
	fmt.Println("start Rock-Paper-Scissors game.")
	scanner := bufio.NewScanner(os.Stdin)

	loop:
	for {
		fmt.Println("1: Play game")
		fmt.Println("2: Show history")
		fmt.Println("3: exit")
		fmt.Print("please enter > ")

		scanner.Scan()
		in := scanner.Text()

		switch in {
		case "1":
			color.Println("Please enter <red>Rock</>, <cyan>Paper</>, or <green>Scissors</>.")
			color.Red.Println("1: Rock")
			color.Cyan.Println("2: Paper")
			color.Green.Println("3: Scissors")
			fmt.Print("please enter >")

			scanner.Scan()
			hand_str := scanner.Text()
			switch hand_str {
			case "1", "2", "3":
				hand_int, _ := strconv.Atoi(hand_str)
				fmt.Println("")
				service.PlayGame(int32(hand_int))
			default:
				fmt.Println("Invalid command.")
				continue
			}
		case "2":
			fmt.Println("Here are your match results.")
			service.ShowMatchResults()
			continue
		case "3":
			fmt.Println("bye.")
			break loop
		default:
			fmt.Println("Invalid command")
			continue
		}
	}
}