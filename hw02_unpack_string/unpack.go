package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {

	var currCharIsNum, nextCharIsNum bool
	var err error
	var sb strings.Builder

	strArr := []rune(str)
	lenStr := len(strArr)

	for num := 0; num < lenStr; num++ {
		currCharIsNum = unicode.IsNumber(strArr[num])
		if num == lenStr-1 {
			sb.WriteRune(strArr[num])
			break
		}

		nextCharIsNum = unicode.IsNumber(strArr[num+1])

		switch {
		case !currCharIsNum && nextCharIsNum: // буква + цифра --> выводим N букв
			{
				sb.WriteString(strings.Repeat(string(strArr[num]), int(strArr[num+1]-'0')))
				num++ // так как мы использовали 2 символа из строки
			}
		case !currCharIsNum && !nextCharIsNum: // буква + буква --> выводим букву
			{
				sb.WriteRune(strArr[num])
			}
		case currCharIsNum: // ошибка
			{
				sb.Reset()
				num = lenStr
				err = ErrInvalidString
			}
		}
	}
	return sb.String(), err
}
