// You can edit this code!
// Click here and start typing.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Operation struct {
	iaOperands [2]int
	sOperation string
}

var (
	// Значения bNumberFormat: 0 - по умолчанию, 1 - арабские, 2 - римские
	bNumberFormat byte = 0
	mErrors            = map[int]string{
		1: "Введённый запрос не соответствует необходимому формату.",
		2: "Введённые числа заданы в разных системах счисления.",
		3: "Одно или оба вводимых числа не входят в промежуток от 1 до 10 включительно.",
		4: "Ответ в римской системе счисления может быть только натуральным.",
		5: "Что-то пошло не так при конвертации, пните разраба! :)",
	}

	mRomanToArabic = map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	mArabicToRoman = map[int]string{
		1:   "I",
		4:   "IV",
		5:   "V",
		9:   "IX",
		10:  "X",
		40:  "XL",
		50:  "L",
		90:  "XC",
		100: "C",
	}
)

func ExitOnError(iErrorCode int) {
	fmt.Println("ОШИБКА (" + strconv.Itoa(iErrorCode) + "): " + mErrors[iErrorCode])
	os.Exit(0)
}

func ConvertNum(sInput string) int {
	iResult, _er := strconv.Atoi(sInput)
	if _er == nil {
		if (iResult >= 1) && (iResult <= 10) {
			switch bNumberFormat {
			case 2:
				ExitOnError(2)
			default:
				bNumberFormat = 1
				return iResult
			}
		} else {
			ExitOnError(3)
		}
	} else {
		iResult, _boo := ConvertRomanToInt(sInput)
		if _boo {
			if (iResult >= 1) && (iResult <= 10) {
				switch bNumberFormat {
				case 1:
					ExitOnError(2)
				default:
					bNumberFormat = 2
					return iResult
				}
			} else {
				ExitOnError(3)
			}
		} else {
			ExitOnError(1)
		}
	}
	return 0
}

func ConvertIntToRoman(iInput int) string {
	sResult, _boo := mArabicToRoman[iInput]
	if _boo {
		return sResult
	} else {
		sResult = ""
		if iInput/90 == 1 {
			sResult += "XC"
			iInput -= 90
		}
		if iInput/50 == 1 {
			sResult += "L"
			iInput -= 50
		}
		if iInput/40 == 1 {
			sResult += "XL"
			iInput -= 40
		}
		if iInput/10 > 0 {
			for _i := 0; _i < iInput/10; _i++ {
				sResult += "X"
			}
			iInput -= (iInput / 10) * 10
		}
		if iInput/9 == 1 {
			sResult += "IX"
			iInput -= 9
		}
		if iInput/5 == 1 {
			sResult += "V"
			iInput -= 5
		}
		if iInput/4 == 1 {
			sResult += "IV"
			iInput -= 4
		}
		if iInput > 0 {
			for _i := 0; _i < iInput; _i++ {
				sResult += "I"
			}
		}
		return sResult
	}

}

func ConvertRomanToInt(sInput string) (int, bool) {
	iResult, iTopNum := 0, 0
	for _i := len(sInput) - 1; _i >= 0; _i-- {
		iCurNum, _boo := mRomanToArabic[rune(sInput[_i])]
		if _boo {
			if iCurNum >= iTopNum {
				iTopNum = iCurNum
				iResult += iCurNum
				continue
			}
		} else {
			return 0, false
		}
		iResult -= iCurNum
	}
	return iResult, true
}

func Calc(sInput string) {
	bNumberFormat = 0
	sResult, iResult := "", 0
	opFunc := Operation{[2]int{0, 0}, ""}
	sArray := strings.Split(sInput, " ")
	if len(sArray) == 3 {
		for _i := 0; _i <= 2; _i += 2 {
			opFunc.iaOperands[_i/2] = ConvertNum(sArray[_i])
		}
		switch sArray[1] {
		case "+":
			iResult = opFunc.iaOperands[0] + opFunc.iaOperands[1]
		case "-":
			iResult = opFunc.iaOperands[0] - opFunc.iaOperands[1]
		case "*":
			iResult = opFunc.iaOperands[0] * opFunc.iaOperands[1]
		case "/":
			iResult = opFunc.iaOperands[0] / opFunc.iaOperands[1]
		default:
			ExitOnError(1)
		}
	} else {
		ExitOnError(1)
	}
	switch bNumberFormat {
	case 1:
		sResult = strconv.Itoa(iResult)
	case 2:
		if iResult > 0 {
			sResult = ConvertIntToRoman(iResult)
		} else {
			ExitOnError(4)
		}
	default:
		ExitOnError(5)
	}
	fmt.Println(sResult)
}

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	sInitInput, _ := inputReader.ReadString('\n')
	sInitInput = strings.TrimSuffix(sInitInput, "\n")
	Calc(sInitInput)
}
