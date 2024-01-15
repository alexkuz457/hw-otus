package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	// Place your code here.
	var currChar, nextChar string
	var prevCharIsNum, currCharIsNum, nextCharIsNum bool
	var resStr string
	var err error

	strArr := []rune(str)
	lenStr := len(strArr)

	for num := 0; num < lenStr; num++ {
		currChar = string(strArr[num])
		currCharIsNum = unicode.IsNumber(strArr[num])
		if num == lenStr-1 {
			resStr = resStr + currChar
			break
		}
		nextCharIsNum = unicode.IsNumber(strArr[num+1])
		nextChar = string(strArr[num+1])
		switch {
		case !currCharIsNum && nextCharIsNum: // буква + цифра --> выводим N букв
			{
				countRep, _ := strconv.Atoi(nextChar)
				resStr += strings.Repeat(currChar, countRep)
				num++ // так как мы использовали 2 символа из строки
			}
		case !prevCharIsNum && !currCharIsNum: // буква + буква --> выводим букву
			{
				resStr += currChar
			}
		case currCharIsNum: // ошибка
			{
				resStr = "invalid string"
				num = lenStr
				err = ErrInvalidString
			}
		}
	}
	return resStr, err
}
