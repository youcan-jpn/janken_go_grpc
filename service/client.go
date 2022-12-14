package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	"github.com/gookit/color"

	"github.com/youcan-jpn/janken_go_grpc/pb"
	"github.com/youcan-jpn/janken_go_grpc/pkg"
)

func PlayGame(handShape int32) {
	address := "localhost:50051"
	conn, err := grpc.Dial(
		address,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second,
	)
	defer cancel()

	client := pb.NewRockPaperScissorsServiceClient(conn)
	playRequest := pb.PlayRequest{
		Handshape: pkg.EncodeHandShapes(handShape),
	}

	reply, err := client.PlayGame(ctx, &playRequest)
	if err != nil {
		log.Fatal("Request failed.")
		return
	}
	matchResult := reply.GetMatchResult()
	fmt.Println("***********************************")
	printResult(
		matchResult.YourHandShape.String(),
		matchResult.OpponentHandShape.String(),
		matchResult.Result.String())
	fmt.Println("***********************************")
	fmt.Println()
}

func printResult(yourHand string, opponentHand string, result string) {
	color.Printf("Your hand shape: %s\n", colorfulHand(yourHand))
	color.Printf("Opponent hand shape: %s\n", colorfulHand(opponentHand))
	color.Printf("Result: %s\n", colorfulResult(result))
}

func colorfulHand(hand string) string {
	switch hand {
	case "ROCK":
		return "<red>Rock</>"
	case "PAPER":
		return "<cyan>Paper</>"
	case "SCISSORS":
		return "<green>Scissors</>"
	default:
		return hand
	}
}

func colorfulResult(result string) string {
	switch result {
	case "WIN":
		return "<red>WIN</>"
	case "DRAW":
		return "<default>DRAW</>"
	case "LOSE":
		return "<blue>LOSE</>"
	default:
		return "UNKNOWN RESULT"
	}
}

func ShowMatchResults() {
	address := "localhost:50051"
	conn, err := grpc.Dial(
		address,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second,
	)
	defer cancel()

	client := pb.NewRockPaperScissorsServiceClient(conn)
	showRequest := pb.ShowRequest{}

	reply, err := client.ShowHistory(ctx, &showRequest)
	if err != nil {
		log.Fatal("Request failed.")
		return
	}

	report := reply.GetHistory()
	if len(report.MatchResults) == 0 {
		fmt.Println("***********************************")
		fmt.Println("There are no match results.")
		fmt.Println("***********************************")
		fmt.Println()
		return
	}

	fmt.Println("***********************************")
	for k, v := range report.MatchResults {
		fmt.Println(k + 1)
		fmt.Printf("Your hand shapes: %s \n", v.YourHandShape.String())
		fmt.Printf("Opponent hand shapes: %s \n", v.OpponentHandShape.String())
		fmt.Printf("Result: %s \n", v.Result.String())
		fmt.Printf("Datetime of match: %s \n", v.CreatedAt)
		fmt.Println()
	}

	fmt.Printf("Number of games: %d \n", reply.GetHistory().NumberOfGames)
	fmt.Printf("Number of wins: %d \n", reply.GetHistory().NumberOfWins)
	fmt.Println("***********************************")
	fmt.Println()
}