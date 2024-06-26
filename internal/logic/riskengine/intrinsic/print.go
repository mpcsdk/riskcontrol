package intrinsic

import (
	"fmt"
)

func Println(s ...interface{}) {
	fmt.Println(s...)
}

type Intrinsic struct {
}
