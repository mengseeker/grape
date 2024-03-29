package arrays

import "strconv"

func JoinInt(a []int, sep string) string {
	bs := make([]byte, 0, len(a)*2)
	sepb := []byte(sep)
	if len(a) > 0 {
		bs = strconv.AppendInt(bs, int64(a[0]), 10)
	}
	for _, v := range a[1:] {
		bs = append(bs, sepb...)
		bs = strconv.AppendInt(bs, int64(v), 10)
	}
	return string(bs)
}

func FindStr(a []string, str string) int {
	for i, val := range a {
		if val == str {
			return i
		}
	}
	return -1
}

func ContainsStr(a []string, str string) bool {
	return FindStr(a, str) > 0
}
