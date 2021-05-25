package core

import (
	"fmt"
)

func isNum(char byte) bool {
	if char >= '0' && char <= '9' {
		return true
	}
	return false
}
func GetResult(str string) (result int) {
	fmt.Println("getResult", str)
	preSign := '+'
	num := 0
	stack := []int{}
	for i, char := range str {
		flag := isNum(byte(char))
		if flag {
			num = num*10 + int(char-'0')
		}
		if !flag && char != ' ' || i == len(str)-1 {
			switch preSign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			default:
				panic("字符串格式不正确:" + str)
			}
			num = 0
			preSign = char
		}
	}
	for _, v := range stack {
		result += v
	}
	return
}
