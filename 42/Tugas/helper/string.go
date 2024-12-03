package helper

import "strconv"

func StringToBool(str string) bool {
	convBool, _ := strconv.ParseBool(str)
	return convBool
}

func StringToInt(num string) int {
	convInt, _ := strconv.Atoi(num)
	return convInt
}
