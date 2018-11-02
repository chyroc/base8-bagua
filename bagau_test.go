package base8_bagua

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	as := assert.New(t)

	for i := 0; i < 255; i++ {
		s := Encode([]byte{byte(i)})
		b, err := Decode(s)
		as.Nil(err)
		as.Equal([]byte{byte(i)}, b)
	}

	for i := 0; i < 255; i++ {
		s := Encode([]byte(strconv.Itoa(i) + " - test"))
		b, err := Decode(s)
		as.Nil(err)
		as.Equal([]byte(strconv.Itoa(i)+" - test"), b)
	}
}

func TestEncode(t *testing.T) {
	as := assert.New(t)

	m1 := map[byte]string{
		0:  "☰☰☰",
		1:  "☰☰☲",
		2:  "☰☰☴",
		3:  "☰☰☶",
		4:  "☰☱☰",
		5:  "☰☱☲",
		6:  "☰☱☴",
		7:  "☰☱☶",
		8:  "☰☲☰",
		9:  "☰☲☲",
		10: "☰☲☴",
		11: "☰☲☶",
		12: "☰☳☰",
		13: "☰☳☲",
		14: "☰☳☴",
		15: "☰☳☶",
		16: "☰☴☰",
		17: "☰☴☲",
		18: "☰☴☴",
		19: "☰☴☶",
		20: "☰☵☰",
	}

	for k, v := range m1 {
		as.Equal(v, Encode([]byte{k}))
	}

	m2 := map[string]string{
		"0":             "☱☴☰",
		"test":          "☳☵☰☶☲☵☶☳☳☵☰",
		"bagua":         "☳☰☴☶☰☵☴☷☳☵☲☶☰☴",
		"base8 - bagua": "☳☰☴☶☰☵☶☳☳☱☲☳☴☰☴☰☱☳☲☲☰☱☴☲☳☰☲☶☳☵☶☵☳☰☲",
	}

	for k, v := range m2 {
		as.Equal(v, Encode([]byte(k)))
	}
}
