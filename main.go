package main

import (
	"reversi_server/reversi"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	PutStoneSiteAndPlayerAndOpponent string `json:"PutStoneSiteAndPlayerAndOpponent"`
}

type MyResponse struct {
	PuttedPlayerAndOpponent string `json:"PuttedPlayerAndOpponent"`
}

func hello(event MyEvent) (MyResponse, error) {
	var putStoneSite int64
	var player int64
	var opponent int64
	putStoneSite, _ = strconv.ParseInt(strings.Split(event.PutStoneSiteAndPlayerAndOpponent, ";")[0], 10, 64)
	player, _ = strconv.ParseInt(strings.Split(event.PutStoneSiteAndPlayerAndOpponent, ";")[1], 10, 64)
	opponent, _ = strconv.ParseInt(strings.Split(event.PutStoneSiteAndPlayerAndOpponent, ";")[2], 10, 64)

	var outputPlayer int64
	var outputOpponent int64
	outputPlayer, outputOpponent = reversi.PutStoneToSpecifiedSiteAndReverse(putStoneSite, player, opponent)

	return MyResponse{PuttedPlayerAndOpponent: strconv.FormatInt(outputPlayer, 10) + ";" + strconv.FormatInt(outputOpponent, 10)}, nil
}

func main() {
	lambda.Start(hello)
}
