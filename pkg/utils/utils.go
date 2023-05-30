package utils

import (
	"fmt"
	"strconv"
)

func GetMouth(mouth int) string {
	curmoth := strconv.Itoa(mouth)
	if mouth < 10 {
		curmoth = fmt.Sprintf("0%v", mouth)
	}
	return curmoth
}

func ConvertDate(year, mouth, day int) string {
	date := fmt.Sprintf("%v-%s-%v", year, GetMouth(mouth), day)
	return date
}
