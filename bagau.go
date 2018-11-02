package base8_bagua

import (
	"errors"
	"strconv"
	"strings"
)

/*
base8-bagua
2^3 = 8
每3个b为一组，一个字节8个b，公约数24（24个b），可以表示3个字节，用8个组（一卦）来表示

算法

Encode
输入：[]byte
3个byte一组，转成二进制后每3个b切分，即组成8进制数，再转成八卦
如果少了1-2byte，补0
返回八卦

Decode
输入八卦，按照8进制处理，转成二进制，然后每8个b组成一个byte，最后的全0的b丢弃
返回 []byte
*/

const (
	qian = "☰" // 乾
	dui  = "☱" // 兑
	li   = "☲" // 离
	zhen = "☳" // 震
	xun  = "☴" // 巽
	kan  = "☵" // 坎
	gen  = "☶" // 艮
	kun  = "☷" // 坤
)

var m1 = map[int]string{
	0: qian,
	1: dui,
	2: li,
	3: zhen,
	4: xun,
	5: kan,
	6: gen,
	7: kun,
}

var m2 = map[string][3]int{
	qian: {0, 0, 0},
	dui:  {0, 0, 1},
	li:   {0, 1, 0},
	zhen: {0, 1, 1},
	xun:  {1, 0, 0},
	kan:  {1, 0, 1},
	gen:  {1, 1, 0},
	kun:  {1, 1, 1},
}

func Encode(src []byte) string {
	bs := make([]int, len(src)*8)
	bl := len(bs)
	for k, v := range src {
		byteTo2(int(v), bs[k*8:k*8+8])
	}

	buf := make([]string, (bl+2)/3)
	for i := 0; i*3+2 < len(bs); i++ {
		buf[i] = m1[bs[i*3]<<2+bs[i*3+1]<<1+bs[i*3+2]]
	}

	switch bl % 3 {
	case 1:
		buf[(bl+2)/3-1] = m1[bs[bl-1]<<2]
	case 2:
		buf[(bl+2)/3-1] = m1[bs[bl-2]<<2+bs[bl-1]<<1]
	}

	return strings.Join(buf, "")
}

func Decode(s string) ([]byte, error) {
	if s == "" {
		return nil, nil
	}

	sl := len(s)

	is := make([]int, sl)
	for i := 0; i < sl/3; i++ {
		b, ok := m2[s[i*3:i*3+3]]
		if !ok {
			return nil, errors.New("invalid string, cur: " + strconv.Itoa(i))
		}
		copy(is[i*3:i*3+3], b[:])
	}

	buf := make([]byte, sl/8)
	for i := 0; i < sl/8; i++ {
		buf[i] = b8ToByte(is[i*8 : i*8+8])
	}

	return buf, nil
}

func b8ToByte(b []int) byte {
	return byte(b[0]<<7 + b[1]<<6 + b[2]<<5 + b[3]<<4 + b[4]<<3 + b[5]<<2 + b[6]<<1 + b[7])
}

func byteTo2(byt int, dst []int) {
	var i = 7
	for byt != 0 {
		dst[i] = byt % 2
		byt = byt >> 1
		i--
	}
	return
}
