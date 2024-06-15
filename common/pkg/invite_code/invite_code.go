package invite_code

import "strings"

// 邀请码生成规则
const (
	base       = "PLM9KN1IBJF8HU3VGW2EY7TCQ6XZ4ASD5R" // 把0剔除，当做补位符号，比如小于四位的邀请码在高位补0，把O也去掉防止混淆，这样36进制就变成了34进制。
	suffixChar = "0"                                  // 补位字符，不能与自定义重复
	binLen     = 34                                   // 进制长度：len(Base)
	codeLen    = 6                                    // 生成邀请码最小长度
	baseNum    = 100000
)

func strPadLeft(input string, padLength int, padString string) string {
	output := padString
	for padLength > len(output) {
		output += output
	}
	if len(input) >= padLength {
		return input
	}
	return output[:padLength-len(input)] + input
}

func strRev(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, v := range s {
		n--
		runes[n] = v
	}
	return string(runes[n:])
}

func pow(a, b int64) int64 {
	var p int64
	p = 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}

		b >>= 1
		a *= a
	}
	return p
}

func Encode(id int64) string {
	code := ""
	id += baseNum
	for id > 0 {
		mod := id % binLen
		id = (id - mod) / binLen
		code = string(base[mod]) + code
	}
	//if len(code) < codeLen {
	//	code = strPadLeft(code, codeLen, suffixChar)
	//}
	return code
}

func Decode(code string) int64 {
	var num int64
	//index := strings.LastIndex(code, "0")
	//if index != -1 {
	//	code = code[index+1:]
	//}
	code = strRev(code)
	for i := 0; i < len(code); i++ {
		num += int64(strings.Index(base, string(code[i]))) * pow(binLen, int64(i))
	}
	num -= baseNum
	return num
}
