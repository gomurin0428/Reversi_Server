package reversi

import (
	"strings"
	"testing"
)

func TestBitStringTo64BitInt_WhenOutputIs63(t *testing.T) {
	var input string = `
	111111
	000000
	000000
	000000
	000000
	000000`
	input = strings.ReplaceAll(input, "\n", "")
	input = strings.ReplaceAll(input, "\t", "")
	var output int64 = bitStringTo64BitInt(input)
	if output != 63 {
		t.Fatalf("failed Test Output is %d", output)
	}
}

func TestBitStringTo64BitInt_WhenOutputIs127(t *testing.T) {
	var input string = `
	111111
	100000
	000000
	000000
	000000
	000000`
	input = strings.ReplaceAll(input, "\n", "")
	input = strings.ReplaceAll(input, "\t", "")
	var output int64 = bitStringTo64BitInt(input)
	if output != 127 {
		t.Fatalf("failed Test Output is %d", output)
	}
}

func TestBitStringTo64BitInt_WhenOutputIs1073741951(t *testing.T) {
	var input string = `
	111111
	100000
	000000
	000000
	000000
	100000`
	input = strings.ReplaceAll(input, "\n", "")
	input = strings.ReplaceAll(input, "\t", "")
	var output int64 = bitStringTo64BitInt(input)
	if output != 1073741951 {
		t.Fatalf("failed Test Output is %d", output)
	}
}

func TestBitStringTo64BitInt_WhenOutputIs34359738431(t *testing.T) {
	var input string = `
	111111
	000000
	000000
	000000
	000000
	000001`
	input = strings.ReplaceAll(input, "\n", "")
	input = strings.ReplaceAll(input, "\t", "")
	var output int64 = bitStringTo64BitInt(input)
	if output != 34359738431 {
		t.Fatalf("failed Test Output is %d", output)
	}
}

func TestGetLeftContinuousStone_WhenInitPattern(t *testing.T) {

	var player string = `
	000000
	000000
	000100
	001000
	000000
	000000`
	player = removeNot0Or1String(player)

	var opponent string = `
	000000
	000000
	001000
	000100
	000000
	000000
	`
	opponent = removeNot0Or1String(opponent)

	var output int64 = getLeftContinuousStone(bitStringTo64BitInt(player), bitStringTo64BitInt(opponent))

	var expectation string = `
	000000
	000000
	001000
	000000
	000000
	000000`

	expectation = removeNot0Or1String(expectation)

	if sixtyFourBitIntToBitString(output) != expectation {
		t.Fatalf("failed Test Output is %s", sixtyFourBitIntToBitString(output))
	}

}

func TestGetLeftContinuousStone_WhenMiddlePattern1(t *testing.T) {

	var player string = `
	000000
	000000
	000100
	001000
	000000
	000000`
	player = removeNot0Or1String(player)

	var opponent string = `
	000000
	000000
	011000
	000100
	000000
	000000
	`
	opponent = removeNot0Or1String(opponent)

	var output int64 = getLeftContinuousStone(bitStringTo64BitInt(player), bitStringTo64BitInt(opponent))

	var expectation string = `
	000000
	000000
	011000
	000000
	000000
	000000`

	expectation = removeNot0Or1String(expectation)

	if sixtyFourBitIntToBitString(output) != expectation {
		t.Fatalf("failed Test Output is %s", sixtyFourBitIntToBitString(output))
	}

}

func TestGetLeftContinuousStone_WhenMiddlePattern2(t *testing.T) {

	var player string = `
	000000
	000000
	000100
	001000
	000000
	000000`
	player = removeNot0Or1String(player)

	var opponent string = `
	000000
	000000
	111000
	000100
	000000
	000000
	`
	opponent = removeNot0Or1String(opponent)

	var output int64 = getLeftContinuousStone(bitStringTo64BitInt(player), bitStringTo64BitInt(opponent))

	var expectation string = `
	000000
	000000
	011000
	000000
	000000
	000000`

	expectation = removeNot0Or1String(expectation)

	if sixtyFourBitIntToBitString(output) != expectation {
		t.Fatalf("failed Test Output is %s", sixtyFourBitIntToBitString(output))
	}

}

func TestGetLeftContinuousStone_WhenMiddlePattern3(t *testing.T) {

	var player string = `
	000000
	000000
	000100
	001000
	000000
	000000`
	player = removeNot0Or1String(player)

	var opponent string = `
	000000
	000000
	011000
	010100
	000000
	000000
	`
	opponent = removeNot0Or1String(opponent)

	var output int64 = getLeftContinuousStone(bitStringTo64BitInt(player), bitStringTo64BitInt(opponent))

	var expectation string = `
	000000
	000000
	011000
	010000
	000000
	000000`

	expectation = removeNot0Or1String(expectation)

	if sixtyFourBitIntToBitString(output) != expectation {
		t.Fatalf("failed Test Output is %s", sixtyFourBitIntToBitString(output))
	}

}

func TestGetRightContinuousStone_WhenInitPattern(t *testing.T) {

	var player string = `
	000000
	000000
	000100
	001000
	000000
	000000`
	player = removeNot0Or1String(player)

	var opponent string = `
	000000
	000000
	001000
	000100
	000000
	000000
	`
	opponent = removeNot0Or1String(opponent)

	var output int64 = getRightContinuousStone(bitStringTo64BitInt(player), bitStringTo64BitInt(opponent))

	var expectation string = `
	000000
	000000
	000000
	000100
	000000
	000000`

	expectation = removeNot0Or1String(expectation)

	if sixtyFourBitIntToBitString(output) != expectation {
		t.Fatalf("failed Test Output is %s", sixtyFourBitIntToBitString(output))
	}

}

func TestGetDownContinuousStone_WhenInitPattern(t *testing.T) {

	var player string = `
	000000
	000000
	000100
	001000
	000000
	000000`
	player = removeNot0Or1String(player)

	var opponent string = `
	000000
	000000
	001000
	000100
	000000
	000000
	`
	opponent = removeNot0Or1String(opponent)

	var output int64 = getDownContinuousStone(bitStringTo64BitInt(player), bitStringTo64BitInt(opponent))

	var expectation string = `
	000000
	000000
	000000
	000100
	000000
	000000`

	expectation = removeNot0Or1String(expectation)

	if sixtyFourBitIntToBitString(output) != expectation {
		t.Fatalf("failed Test Output is %s", sixtyFourBitIntToBitString(output))
	}

}

func TestGetDownContinuousStone_WhenMiddlePattern1(t *testing.T) {

	var player string = `
	000000
	000000
	000100
	001000
	000000
	000000`
	player = removeNot0Or1String(player)

	var opponent string = `
	000000
	000000
	001000
	000100
	000100
	000000
	`
	opponent = removeNot0Or1String(opponent)

	var output int64 = getDownContinuousStone(bitStringTo64BitInt(player), bitStringTo64BitInt(opponent))

	var expectation string = `
	000000
	000000
	000000
	000100
	000100
	000000`

	expectation = removeNot0Or1String(expectation)

	if sixtyFourBitIntToBitString(output) != expectation {
		t.Fatalf("failed Test Output is %s", sixtyFourBitIntToBitString(output))
	}

}

func TestGetDownContinuousStone_WhenMiddlePattern2(t *testing.T) {

	var player string = `
	000000
	000000
	000100
	001000
	000000
	000000`
	player = removeNot0Or1String(player)

	var opponent string = `
	000000
	000000
	001000
	000100
	000100
	000100
	`
	opponent = removeNot0Or1String(opponent)

	var output int64 = getDownContinuousStone(bitStringTo64BitInt(player), bitStringTo64BitInt(opponent))

	var expectation string = `
	000000
	000000
	000000
	000100
	000100
	000000`

	expectation = removeNot0Or1String(expectation)

	if sixtyFourBitIntToBitString(output) != expectation {
		t.Fatalf("failed Test Output is %s", sixtyFourBitIntToBitString(output))
	}

}

func TestGetUpContinuousStone_WhenInitPattern(t *testing.T) {

	var player string = `
	000000
	000000
	000100
	001000
	000000
	000000`
	player = removeNot0Or1String(player)

	var opponent string = `
	000000
	000000
	001000
	000100
	000000
	000000
	`
	opponent = removeNot0Or1String(opponent)

	var output int64 = getUpContinuousStone(bitStringTo64BitInt(player), bitStringTo64BitInt(opponent))

	var expectation string = `
	000000
	000000
	001000
	000000
	000000
	000000`

	expectation = removeNot0Or1String(expectation)

	if sixtyFourBitIntToBitString(output) != expectation {
		t.Fatalf("failed Test Output is %s", sixtyFourBitIntToBitString(output))
	}

}

func TestGetUpContinuousStone_WhenMiddlePattern1(t *testing.T) {

	var player string = `
	000000
	000000
	000100
	001000
	000000
	000000`
	player = removeNot0Or1String(player)

	var opponent string = `
	000000
	001000
	001000
	000100
	000000
	000000
	`
	opponent = removeNot0Or1String(opponent)

	var output int64 = getUpContinuousStone(bitStringTo64BitInt(player), bitStringTo64BitInt(opponent))

	var expectation string = `
	000000
	001000
	001000
	000000
	000000
	000000`

	expectation = removeNot0Or1String(expectation)

	if sixtyFourBitIntToBitString(output) != expectation {
		t.Fatalf("failed Test Output is %s", sixtyFourBitIntToBitString(output))
	}

}

func TestGetUpRightContinuousStone_WhenInitPattern(t *testing.T) {

	var player string = `
	000000
	000000
	000100
	001000
	000000
	000000`
	player = removeNot0Or1String(player)

	var opponent string = `
	000000
	000000
	001000
	000100
	000000
	000000
	`
	opponent = removeNot0Or1String(opponent)

	var output int64 = getUpRightContinuousStone(bitStringTo64BitInt(player), bitStringTo64BitInt(opponent))

	var expectation string = `
	000000
	000000
	000000
	000000
	000000
	000000`

	expectation = removeNot0Or1String(expectation)

	if sixtyFourBitIntToBitString(output) != expectation {
		t.Fatalf("failed Test Output is %s", sixtyFourBitIntToBitString(output))
	}

}

func TestGetUpRightContinuousStone_WhenMiddlePattern1(t *testing.T) {

	var player string = `
	000000
	000000
	000100
	001000
	000000
	000000`
	player = removeNot0Or1String(player)

	var opponent string = `
	000000
	000010
	001000
	000100
	000000
	000000
	`
	opponent = removeNot0Or1String(opponent)

	var output int64 = getUpRightContinuousStone(bitStringTo64BitInt(player), bitStringTo64BitInt(opponent))

	var expectation string = `
	000000
	000010
	000000
	000000
	000000
	000000`

	expectation = removeNot0Or1String(expectation)

	if sixtyFourBitIntToBitString(output) != expectation {
		t.Fatalf("failed Test Output is %s", sixtyFourBitIntToBitString(output))
	}

}

func TestGetUpRightContinuousStone_WhenMiddlePattern2(t *testing.T) {

	var player string = `
	000000
	000000
	000000
	000000
	000000
	100000`
	player = removeNot0Or1String(player)

	var opponent string = `
	000000
	000010
	000100
	001000
	010000
	000000
	`
	opponent = removeNot0Or1String(opponent)

	var output int64 = getUpRightContinuousStone(bitStringTo64BitInt(player), bitStringTo64BitInt(opponent))

	var expectation string = `
	000000
	000010
	000100
	001000
	010000
	000000`

	expectation = removeNot0Or1String(expectation)

	if sixtyFourBitIntToBitString(output) != expectation {
		t.Fatalf("failed Test Output is %s", sixtyFourBitIntToBitString(output))
	}

}

func TestGetUpLeftContinuousStone_WhenMiddlePattern1(t *testing.T) {

	var player string = `
	000000
	000000
	000100
	001000
	000000
	000000`
	player = removeNot0Or1String(player)

	var opponent string = `
	000000
	001000
	001000
	000100
	000000
	000000
	`
	opponent = removeNot0Or1String(opponent)

	var output int64 = getUpLeftContinuousStone(bitStringTo64BitInt(player), bitStringTo64BitInt(opponent))

	var expectation string = `
	000000
	001000
	000000
	000000
	000000
	000000`

	expectation = removeNot0Or1String(expectation)

	if sixtyFourBitIntToBitString(output) != expectation {
		t.Fatalf("failed Test Output is %s", sixtyFourBitIntToBitString(output))
	}

}

func TestGetUpLeftContinuousStone_WhenMiddlePattern2(t *testing.T) {

	var player string = `
	000000
	000000
	000000
	000000
	000000
	000001`
	player = removeNot0Or1String(player)

	var opponent string = `
	000000
	010000
	001000
	000100
	000010
	000000
	`
	opponent = removeNot0Or1String(opponent)

	var output int64 = getUpLeftContinuousStone(bitStringTo64BitInt(player), bitStringTo64BitInt(opponent))

	var expectation string = `
	000000
	010000
	001000
	000100
	000010
	000000`

	expectation = removeNot0Or1String(expectation)

	if sixtyFourBitIntToBitString(output) != expectation {
		t.Fatalf("failed Test Output is %s", sixtyFourBitIntToBitString(output))
	}

}

func TestGetDownLeftContinuousStone_WhenMiddlePattern2(t *testing.T) {

	var player string = `
	000001
	000000
	000000
	000000
	000000
	000000`
	player = removeNot0Or1String(player)

	var opponent string = `
	000000
	000010
	000100
	001000
	010000
	000000
	`
	opponent = removeNot0Or1String(opponent)

	var output int64 = getDownLeftContinuousStone(bitStringTo64BitInt(player), bitStringTo64BitInt(opponent))

	var expectation string = `
	000000
	000010
	000100
	001000
	010000
	000000`

	expectation = removeNot0Or1String(expectation)

	if sixtyFourBitIntToBitString(output) != expectation {
		t.Fatalf("failed Test Output is %s", sixtyFourBitIntToBitString(output))
	}

}

func TestGetDownRightContinuousStone_WhenMiddlePattern2(t *testing.T) {

	var player string = `
	100000
	000000
	000000
	000000
	000000
	000000`
	player = removeNot0Or1String(player)

	var opponent string = `
	000000
	010000
	001000
	000100
	000010
	000000
	`
	opponent = removeNot0Or1String(opponent)

	var output int64 = getDownRightContinuousStone(bitStringTo64BitInt(player), bitStringTo64BitInt(opponent))

	var expectation string = `
	000000
	010000
	001000
	000100
	000010
	000000`

	expectation = removeNot0Or1String(expectation)

	if sixtyFourBitIntToBitString(output) != expectation {
		t.Fatalf("failed Test Output is %s", sixtyFourBitIntToBitString(output))
	}

}

func TestGetLeftReversibleStone(t *testing.T) {
	var player string = `
	000100
	000100
	000001
	000001
	000010
	000010`
	player = removeNot0Or1String(player)

	var opponent string = `
	001000
	011000
	011110
	000110
	000100
	001100
	`
	opponent = removeNot0Or1String(opponent)

	var output int64 = getLeftReversibleStone(bitStringTo64BitInt(player), bitStringTo64BitInt(opponent))

	var expectation string = `
	010000
	100000
	100000
	001000
	001000
	010000`

	expectation = removeNot0Or1String(expectation)

	if sixtyFourBitIntToBitString(output) != expectation {
		t.Fatalf("failed Test Output is %s", sixtyFourBitIntToBitString(output))
	}
}

func TestGetRightReversibleStone(t *testing.T) {
	var player string = `
	100000
	010000
	001000
	000100
	000010
	010000`
	player = removeNot0Or1String(player)

	var opponent string = `
	011110
	001000
	000010
	000010
	000001
	001111
	`
	opponent = removeNot0Or1String(opponent)

	var output int64 = getRightReversibleStone(bitStringTo64BitInt(player), bitStringTo64BitInt(opponent))

	var expectation string = `
	000001
	000100
	000000
	000001
	000000
	000000`

	expectation = removeNot0Or1String(expectation)

	if sixtyFourBitIntToBitString(output) != expectation {
		t.Fatalf("failed Test Output is %s", sixtyFourBitIntToBitString(output))
	}
}

func TestGetDownReversibleStone(t *testing.T) {
	var player string = `
	100000
	010000
	001000
	000100
	000010
	000001`
	player = removeNot0Or1String(player)

	var opponent string = `
	000000
	100000
	100000
	111000
	111100
	000010
	`
	opponent = removeNot0Or1String(opponent)

	var output int64 = getDownReversibleStone(bitStringTo64BitInt(player), bitStringTo64BitInt(opponent))

	var expectation string = `
	000000
	000000
	000000
	000000
	000000
	101100`

	expectation = removeNot0Or1String(expectation)

	if sixtyFourBitIntToBitString(output) != expectation {
		t.Fatalf("failed Test Output is %s", sixtyFourBitIntToBitString(output))
	}
}

func TestGetUpReversibleStone(t *testing.T) {
	var player string = `
	000000
	000010
	010100
	011000
	110000
	100000`
	player = removeNot0Or1String(player)

	var opponent string = `
	000000
	111000
	101010
	100100
	000000
	000000
	`
	opponent = removeNot0Or1String(opponent)

	var output int64 = getUpReversibleStone(bitStringTo64BitInt(player), bitStringTo64BitInt(opponent))

	var expectation string = `
	111000
	000000
	000000
	000000
	000000
	000000`

	expectation = removeNot0Or1String(expectation)

	if sixtyFourBitIntToBitString(output) != expectation {
		t.Fatalf("failed Test Output is %s", sixtyFourBitIntToBitString(output))
	}
}

func TestGetUpRightReversibleStone(t *testing.T) {
	var player string = `
	000000
	000000
	001000
	000000
	100000
	100000`
	player = removeNot0Or1String(player)

	var opponent string = `
	000000
	000110
	000100
	011000
	010000
	000000
	`
	opponent = removeNot0Or1String(opponent)

	var output int64 = getUpRightReversibleStone(bitStringTo64BitInt(player), bitStringTo64BitInt(opponent))

	var expectation string = `
	000011
	000000
	000000
	000000
	000000
	000000`

	expectation = removeNot0Or1String(expectation)

	if sixtyFourBitIntToBitString(output) != expectation {
		t.Fatalf("failed Test Output is %s", sixtyFourBitIntToBitString(output))
	}
}

func TestGetUpLeftReversibleStone(t *testing.T) {
	var player string = `
	000000
	000000
	000100
	000000
	000001
	000001`
	player = removeNot0Or1String(player)

	var opponent string = `
	000000
	011000
	001000
	000110
	000010
	000000
	`
	opponent = removeNot0Or1String(opponent)

	var output int64 = getUpLeftReversibleStone(bitStringTo64BitInt(player), bitStringTo64BitInt(opponent))

	var expectation string = `
	110000
	000000
	000000
	000000
	000000
	000000`

	expectation = removeNot0Or1String(expectation)

	if sixtyFourBitIntToBitString(output) != expectation {
		t.Fatalf("failed Test Output is %s", sixtyFourBitIntToBitString(output))
	}
}

func TestGetDownLeftReversibleStone(t *testing.T) {
	var player string = `
	000001
	000001
	000000
	000100
	000000
	000000`
	player = removeNot0Or1String(player)

	var opponent string = `
	000000
	000010
	000100
	001000
	011010
	000000
	`
	opponent = removeNot0Or1String(opponent)

	var output int64 = getDownLeftReversibleStone(bitStringTo64BitInt(player), bitStringTo64BitInt(opponent))

	var expectation string = `
	000000
	000000
	000000
	000000
	000000
	110000`

	expectation = removeNot0Or1String(expectation)

	if sixtyFourBitIntToBitString(output) != expectation {
		t.Fatalf("failed Test Output is %s", sixtyFourBitIntToBitString(output))
	}
}

func TestGetDownRightReversibleStone(t *testing.T) {
	var player string = `
	100000
	100000
	000000
	001000
	000000
	000000`
	player = removeNot0Or1String(player)

	var opponent string = `
	000000
	010000
	011000
	000100
	000110
	000000
	`
	opponent = removeNot0Or1String(opponent)

	var output int64 = getDownRightReversibleStone(bitStringTo64BitInt(player), bitStringTo64BitInt(opponent))

	var expectation string = `
	000000
	000000
	000000
	000000
	000000
	000011`

	expectation = removeNot0Or1String(expectation)

	if sixtyFourBitIntToBitString(output) != expectation {
		t.Fatalf("failed Test Output is %s", sixtyFourBitIntToBitString(output))
	}
}

func TestGetReversibleStone(t *testing.T) {
	//1048576
	var player string = `
	000000
	000000
	000000
	001000
	000000
	000000`
	player = removeNot0Or1String(player)

	//237559808
	var opponent string = `
	000000
	000000
	011100
	010100
	011100
	000000
	`
	opponent = removeNot0Or1String(opponent)

	var output int64 = GetReversibleStone(bitStringTo64BitInt(player), bitStringTo64BitInt(opponent))

	//22553036096
	var expectation string = `
	000000
	101010
	000000
	100010
	000000
	101010`

	expectation = removeNot0Or1String(expectation)

	if sixtyFourBitIntToBitString(output) != expectation {
		t.Fatalf("failed Test Output is %s", sixtyFourBitIntToBitString(output))
	}
}

func TestGetOpponentsStoneSite(t *testing.T) {
	var player string = `
	000000
	000000
	000000
	001000
	000000
	000000`
	player = removeNot0Or1String(player)

	var opponent string = `
	000000
	000000
	011100
	010100
	011100
	000000
	`
	opponent = removeNot0Or1String(opponent)

	var output int64 = GetOpponentsStoneSite(bitStringTo64BitInt(player), bitStringTo64BitInt(opponent))

	var expectation string = `
	000000
	101010
	000000
	100010
	000000
	101010`

	expectation = removeNot0Or1String(expectation)

	var expectationInt64 = bitStringTo64BitInt(expectation)

	if (expectationInt64 & (output)) == 0 {
		t.Fatalf("failed Test Output is %s", sixtyFourBitIntToBitString(output))
	}

}

func TestPutStoneToSpecifiedSiteAndReverse(t *testing.T) {

	//256
	var putStoneSite string = `
	000000
	001000
	000000
	000000
	000000
	000000
	`

	putStoneSite = removeNot0Or1String(putStoneSite)
	//1048576
	var player string = `
	000000
	000000
	000000
	001000
	000000
	000000`
	player = removeNot0Or1String(player)

	//237559808
	var opponent string = `
	000000
	000000
	011100
	010100
	011100
	000000
	`
	opponent = removeNot0Or1String(opponent)

	var outputPlayer int64
	var outputOpponent int64

	outputPlayer, outputOpponent = PutStoneToSpecifiedSiteAndReverse(bitStringTo64BitInt(putStoneSite), bitStringTo64BitInt(player), bitStringTo64BitInt(opponent))

	//1065216
	var expectedPlayer string = `
	000000
	001000
	001000
	001000
	000000
	000000
	`

	expectedPlayer = removeNot0Or1String(expectedPlayer)

	//237543424
	var expectedOpponent string = `
	000000
	000000
	010100
	010100
	011100
	000000
	`

	expectedOpponent = removeNot0Or1String(expectedOpponent)

	if !(bitStringTo64BitInt(expectedPlayer) == outputPlayer && bitStringTo64BitInt(expectedOpponent) == outputOpponent) {
		t.Fatalf("failed;Expected Output Player is %s but %s ; Expected Output Opponent is %s but %s", (expectedPlayer), sixtyFourBitIntToBitString(outputPlayer), (expectedOpponent), sixtyFourBitIntToBitString(outputOpponent))
	}

}

func removeNot0Or1String(input string) string {
	var output string = input
	output = strings.ReplaceAll(output, "\n", "")
	output = strings.ReplaceAll(output, "\t", "")
	return output
}
