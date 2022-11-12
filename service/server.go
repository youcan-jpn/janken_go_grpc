package service

import (
	"context"
	"math/rand"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/youcan-jpn/janken_go_grpc/pb"
	"github.com/youcan-jpn/janken_go_grpc/pkg"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type RockPaperScissorsService struct {
	numberOfGames int32
	numberOfWins int32
	matchResults []*pb.MatchResult
}

func NewRockPaperScissorsService() *RockPaperScissorsService {
	return &RockPaperScissorsService{
		numberOfGames: 0,
		numberOfWins: 0,
		matchResults: make([]*pb.MatchResult, 0),
	}
}

func (s *RockPaperScissorsService) PlayGame(ctx context.Context, req *pb.PlayRequest) (*pb.PlayResponse, error) {
	if req.Handshape == pb.HandShape_HAND_SHAPES_UNKNOWN {
		return nil, status.Errorf(codes.InvalidArgument, "Choose Rock, Paper, or Scissors.")
	}

	opponentHandShape := pkg.EncodeHandShapes(int32(rand.Intn(3) + 1))

	var result pb.Result
	if req.Handshape == opponentHandShape {
		result = pb.Result_DRAW
	} else if (req.Handshape.Number()-opponentHandShape.Number()+3)%3 == 1 {
		result = pb.Result_WIN
	} else {
		result = pb.Result_LOSE
	}
	now := time.Now().String()
	matchResult := &pb.MatchResult{
		YourHandShape: req.Handshape,
		OpponentHandShape: opponentHandShape,
		Result: result,
		CreatedAt: now,
	}
	s.numberOfGames++
	if result == pb.Result_WIN {
		s.numberOfWins++
	}
	s.matchResults = append(s.matchResults, matchResult)

	return &pb.PlayResponse{
		MatchResult: matchResult,
	}, nil
}

func (s *RockPaperScissorsService) ShowHistory(ctx context.Context, req *pb.ShowRequest) (*pb.ShowResponse, error) {
	return &pb.ShowResponse{
		History: &pb.History{
			NumberOfGames: s.numberOfGames,
			NumberOfWins: s.numberOfWins,
			MatchResults: s.matchResults,
		},
	}, nil
}
