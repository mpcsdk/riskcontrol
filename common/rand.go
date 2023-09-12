package common

import (
	"strconv"

	"github.com/dchest/captcha"
)

func RandomDigits(n int) string {
	d := captcha.RandomDigits(6)
	code := ""
	for _, b := range d {
		code += strconv.Itoa(int(b))
	}

	return code
}
