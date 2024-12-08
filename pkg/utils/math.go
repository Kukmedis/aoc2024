package utils

func Abs(num int) int {
	absNum := num
	if num < 0 {
		absNum = -num
	}
	return absNum
}
