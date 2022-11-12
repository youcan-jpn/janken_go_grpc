package pkg

import "github.com/youcan-jpn/janken_go_grpc/pb"

func EncodeHandShapes(n int32) pb.HandShape {
	switch n {
	case 1:
		return pb.HandShape_ROCK
	case 2:
		return pb.HandShape_PAPER
	case 3:
		return pb.HandShape_SCISSORS
	default:
		return pb.HandShape_HAND_SHAPES_UNKNOWN
	}
}