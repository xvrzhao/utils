package strings

import (
	"math/rand"
	"time"
)

// RandLetterNum 生成包含大小写字母和数字的随机字符串，并发安全
func RandLetterNum(len int) string {
	const (
		numLetters    = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		numLettersLen = 62
		bits          = 6        // 62 的二进制位数，111110 共 6 位
		mask          = 0b111111 // 6 位掩码
		times         = 10       // 生成一个 63 位随机数，每次取 6 位，可以取 10 次
	)

	res := make([]byte, len)
	rand.Seed(time.Now().UnixNano())
	asset, remain := rand.Int63(), times

	for i := 0; i < len; i++ {
		if remain == 0 {
			asset, remain = rand.Int63(), times
		}
		if bits6 := asset & mask; bits6 < numLettersLen {
			res[i] = numLetters[bits6]
		} else {
			i--
		}
		asset >>= bits
		remain--
	}

	return string(res)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
