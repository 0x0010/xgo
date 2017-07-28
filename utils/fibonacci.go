package utils

import "github.com/0x0010/xgo/stringutil"

// Fibonacci generate fibonacci sequence
// no limit
func Fibonacci() func() string {
	fa := []string{"0", "1"}
	flen := len(fa)
	pre, prepre, latest, idx := "", "", "", 0
	return func() string {
		if idx < flen {
			latest, idx = fa[idx], idx+1
		} else {
			latest = sumLargeNum(pre, prepre)
		}
		prepre, pre = pre, latest
		return latest
	}
}

// SumLargeNum sum two large numbers format by string
func sumLargeNum(n1, n2 string) string {
	large, small := whichIsLarge([]rune(n1), []rune(n2))
	sum := []int32{}
	var shift int32
	var bitSum int32
	for i, lLen, sLen := 0, len(large), len(small); i < lLen; i++ {
		// 遍历至较小的那个数的第一位
		if sLen-i-1 >= 0 {
			bitSum = large[lLen-i-1] + small[sLen-i-1] - 96 + shift
		} else {
			// 超出较小的那个数的第一位时，直接取大数作为和
			bitSum = large[lLen-i-1] + shift - 48
		}
		// 重置进位标识
		shift = 0
		// 进位条件是大于等于10
		if bitSum >= 10 {
			bitSum -= 10
			// 如果当前位求和出现进位，则标记shift
			shift = 1
		}
		sum = append(sum, bitSum+48)
	}
	// 遍历结束时，如果进位符是1，则在最终结果的高位补1
	if shift == 1 {
		sum = append(sum, 49)
	}
	return stringutil.Reverse(string(sum))
}

func whichIsLarge(r1 []rune, r2 []rune) ([]rune, []rune) {
	if len(r1) > len(r2) {
		return r1, r2
	}
	return r2, r1
}
