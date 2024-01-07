package utils

func StrToInt(str string) int {
	res := 0
	for _, c := range str {
		res = res*10 + int(c-'0')
	}
	return res
}

func ByteToFloat32(b []byte) float32 {
	res := 0.0
	for _, c := range b {
		res = res*10 + float64(c-'0')
	}
	return float32(res)
}
