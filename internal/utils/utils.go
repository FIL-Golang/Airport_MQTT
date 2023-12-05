package utils

func StrToInt(str string) int {
	res := 0
	for _, c := range str {
		res = res*10 + int(c-'0')
	}
	return res
}
