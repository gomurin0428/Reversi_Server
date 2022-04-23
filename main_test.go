package main

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

func removeNot0Or1String(input string) string {
	var output string = input
	output = strings.ReplaceAll(output, "\n", "")
	output = strings.ReplaceAll(output, "\t", "")
	return output
}
