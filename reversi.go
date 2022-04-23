package reversi

import (
	"fmt"
	"strconv"
)

const horizontalMask string = "011110011110011110011110011110011110"
const verticalMask string = "000000111111111111111111111111000000"
const allSideMask string = "000000011110011110011110011110000000"

/*
置ける場所を列挙する関数

引数は、自分の石の配置状況、敵の石の配置状況
戻り値は置ける場所に1の立った整数
*/
func GetReversibleStone(player int64, opponent int64) int64 {
	var output int64 = 0
	output = output | getLeftReversibleStone(player, opponent)
	output = output | getRightReversibleStone(player, opponent)
	output = output | getUpReversibleStone(player, opponent)
	output = output | getDownReversibleStone(player, opponent)
	output = output | getUpRightReversibleStone(player, opponent)
	output = output | getUpLeftReversibleStone(player, opponent)
	output = output | getDownLeftReversibleStone(player, opponent)
	output = output | getDownRightReversibleStone(player, opponent)
	return output
}

/*
置いた結果左側の石をひっくり返せる場所を列挙する関数

引数は、自分の石の配置状況、敵の石の配置状況
戻り値は、置いた結果左側の石をひっくり返すことの出来る全ての場所にビットが立った数。
処理は、
1.左に連続する石を求める関数を実行する。
2.1の戻り値を右にビットシフトして、自分の石も敵の石も置かれていない所にビットが立っている数字とANDする。
*/
func getLeftReversibleStone(player int64, opponent int64) int64 {
	var continuousStone int64 = getLeftContinuousStone(player, opponent)
	var reverseBits int64
	reverseBits, _ = strconv.ParseInt("0000000000000000000000000000111111111111111111111111111111111111", 2, 64)
	var output int64 = ((player | opponent) ^ reverseBits) & (continuousStone >> 1)
	return output
}

/*
置いた結果右側の石をひっくり返せる場所を列挙する関数

置いた結果左側の石をひっくり返せる場所を列挙する関数の、左側を自分の方向に読み替えて実装する。
*/
func getRightReversibleStone(player int64, opponent int64) int64 {
	var continuousStone int64 = getRightContinuousStone(player, opponent)
	var reverseBits int64
	reverseBits, _ = strconv.ParseInt("0000000000000000000000000000111111111111111111111111111111111111", 2, 64)
	var output int64 = ((player | opponent) ^ reverseBits) & (continuousStone << 1)
	return output
}

/*
置いた結果下側の石をひっくり返せる場所を列挙する関数

置いた結果左側の石をひっくり返せる場所を列挙する関数の、左側を自分の方向に読み替えて実装する。
*/
func getDownReversibleStone(player int64, opponent int64) int64 {
	var continuousStone int64 = getDownContinuousStone(player, opponent)
	var reverseBits int64
	reverseBits, _ = strconv.ParseInt("0000000000000000000000000000111111111111111111111111111111111111", 2, 64)
	var output int64 = ((player | opponent) ^ reverseBits) & (continuousStone << 6)
	return output
}

/*
置いた結果上側の石をひっくり返せる場所を列挙する関数

置いた結果左側の石をひっくり返せる場所を列挙する関数の、左側を自分の方向に読み替えて実装する。
*/
func getUpReversibleStone(player int64, opponent int64) int64 {
	var continuousStone int64 = getUpContinuousStone(player, opponent)
	var reverseBits int64
	reverseBits, _ = strconv.ParseInt("0000000000000000000000000000111111111111111111111111111111111111", 2, 64)
	var output int64 = ((player | opponent) ^ reverseBits) & (continuousStone >> 6)
	return output
}

/*
置いた結果右上側の石をひっくり返せる場所を列挙する関数

置いた結果左側の石をひっくり返せる場所を列挙する関数の、左側を自分の方向に読み替えて実装する。
*/
func getUpRightReversibleStone(player int64, opponent int64) int64 {
	var continuousStone int64 = getUpRightContinuousStone(player, opponent)
	var reverseBits int64
	reverseBits, _ = strconv.ParseInt("0000000000000000000000000000111111111111111111111111111111111111", 2, 64)
	var output int64 = ((player | opponent) ^ reverseBits) & (continuousStone >> 5)
	return output
}

/*
置いた結果左上側の石をひっくり返せる場所を列挙する関数

置いた結果左側の石をひっくり返せる場所を列挙する関数の、左側を自分の方向に読み替えて実装する。
*/
func getUpLeftReversibleStone(player int64, opponent int64) int64 {
	var continuousStone int64 = getUpLeftContinuousStone(player, opponent)
	var reverseBits int64
	reverseBits, _ = strconv.ParseInt("0000000000000000000000000000111111111111111111111111111111111111", 2, 64)
	var output int64 = ((player | opponent) ^ reverseBits) & (continuousStone >> 7)
	return output
}

/*
置いた結果左下側の石をひっくり返せる場所を列挙する関数

置いた結果左側の石をひっくり返せる場所を列挙する関数の、左側を自分の方向に読み替えて実装する。
*/
func getDownLeftReversibleStone(player int64, opponent int64) int64 {
	var continuousStone int64 = getDownLeftContinuousStone(player, opponent)
	var reverseBits int64
	reverseBits, _ = strconv.ParseInt("0000000000000000000000000000111111111111111111111111111111111111", 2, 64)
	var output int64 = ((player | opponent) ^ reverseBits) & (continuousStone << 5)
	return output
}

/*
置いた結果右下側の石をひっくり返せる場所を列挙する関数

置いた結果左側の石をひっくり返せる場所を列挙する関数の、左側を自分の方向に読み替えて実装する。
*/
func getDownRightReversibleStone(player int64, opponent int64) int64 {
	var continuousStone int64 = getDownRightContinuousStone(player, opponent)
	var reverseBits int64
	reverseBits, _ = strconv.ParseInt("0000000000000000000000000000111111111111111111111111111111111111", 2, 64)
	var output int64 = ((player | opponent) ^ reverseBits) & (continuousStone << 7)
	return output
}

//左に連続する石を求める関数
func getLeftContinuousStone(player int64, opponent int64) int64 {
	var maskedBoard int64 = opponent & bitStringTo64BitInt(horizontalMask)
	var tmp int64
	tmp = player >> 1 & maskedBoard

	for i := 0; i < 3; i++ {
		tmp = tmp | ((tmp >> 1) & maskedBoard)
	}

	return tmp
}

/*
左に連続する石を求める関数

引数は、自分の石の配置状況、敵の石の配置状況
戻り値は、自分のすべての石の左に連続して存在する全ての敵の石にビットが立った数
1.まず初期状態の盤面を用意する。
2.次に、敵の石の配置状況に対し、horizontal maskをANDして、マスクを適用する。
3.自分の石の配置状況を右にシフトし、マスクされた盤面とANDしてtmpに代入。
4.以後、tmpをビットシフトしてマスクされた盤面とANDしてビットシフトする前のtmpとORする事を3回繰り返す。
5.tmpをreturnする。
*/
func getRightContinuousStone(player int64, opponent int64) int64 {
	var maskedBoard int64 = opponent & bitStringTo64BitInt(horizontalMask)
	var tmp int64
	tmp = player << 1 & maskedBoard

	for i := 0; i < 3; i++ {
		tmp = tmp | ((tmp << 1) & maskedBoard)
	}

	return tmp
}

//下に連続する石を求める関数
func getDownContinuousStone(player int64, opponent int64) int64 {
	var maskedBoard int64 = opponent & bitStringTo64BitInt(verticalMask)
	var tmp int64
	tmp = player << 6 & maskedBoard

	for i := 0; i < 3; i++ {
		tmp = tmp | ((tmp << 6) & maskedBoard)
	}

	return tmp
}

//上に連続する石を求める関数
func getUpContinuousStone(player int64, opponent int64) int64 {
	var maskedBoard int64 = opponent & bitStringTo64BitInt(verticalMask)
	var tmp int64
	tmp = player >> 6 & maskedBoard

	for i := 0; i < 3; i++ {
		tmp = tmp | ((tmp >> 6) & maskedBoard)
	}

	return tmp
}

//右上に連続する石を求める関数
func getUpRightContinuousStone(player int64, opponent int64) int64 {
	var maskedBoard int64 = opponent & bitStringTo64BitInt(allSideMask)
	var tmp int64
	tmp = player >> 5 & maskedBoard

	for i := 0; i < 3; i++ {
		tmp = tmp | ((tmp >> 5) & maskedBoard)
	}

	return tmp
}

//左上に連続する石を求める関数
func getUpLeftContinuousStone(player int64, opponent int64) int64 {
	var maskedBoard int64 = opponent & bitStringTo64BitInt(allSideMask)
	var tmp int64
	tmp = player >> 7 & maskedBoard

	for i := 0; i < 3; i++ {
		tmp = tmp | ((tmp >> 7) & maskedBoard)
	}

	return tmp
}

//左下に連続する石を求める関数
func getDownLeftContinuousStone(player int64, opponent int64) int64 {
	var maskedBoard int64 = opponent & bitStringTo64BitInt(allSideMask)
	var tmp int64
	tmp = player << 5 & maskedBoard

	for i := 0; i < 3; i++ {
		tmp = tmp | ((tmp << 5) & maskedBoard)
	}

	return tmp
}

//右下に連続する石を求める関数
func getDownRightContinuousStone(player int64, opponent int64) int64 {
	var maskedBoard int64 = opponent & bitStringTo64BitInt(allSideMask)
	var tmp int64
	tmp = player << 7 & maskedBoard

	for i := 0; i < 3; i++ {
		tmp = tmp | ((tmp << 7) & maskedBoard)
	}

	return tmp
}

//36文字の0と1からなる6x6のオセロ盤を表す文字列を符号付64bit整数にして返す。
func bitStringTo64BitInt(input string) int64 {
	var output int64
	output = 0
	for i := 0; i < len(input); i++ {
		var str string
		str = input[i : i+1]
		if str == "1" {
			output += 1 << i
		}
	}
	return output
}

//符号付64bit整数を、36文字の0と1からなる6x6のオセロ盤を表す文字列にして返す。
func sixtyFourBitIntToBitString(input int64) string {
	var tmp string
	tmp = fmt.Sprintf("%036b", input)

	var output string
	for i := 0; i < len(tmp); i++ {
		output = output + tmp[len(tmp)-i-1:len(tmp)-i]
	}

	return output
}
