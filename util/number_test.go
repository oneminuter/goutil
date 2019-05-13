package ut

import (
	"fmt"
	"testing"
)

func TestIsNumeric(t *testing.T) {
	var a = "23434234"
	var b = "0XAB001" //十六进制
	var c = "2432sdfsdfdsfd"
	var d = "sdfsd2432fdsfd"
	var e = "sdfsd"
	var f = "1.345793457e8" //科学计数法

	fmt.Printf("a 是否是数字：%v \n", IsNumeric(a))
	fmt.Printf("b 是否是数字：%v \n", IsNumeric(b))
	fmt.Printf("c 是否是数字：%v \n", IsNumeric(c))
	fmt.Printf("d 是否是数字：%v \n", IsNumeric(d))
	fmt.Printf("e 是否是数字：%v \n", IsNumeric(e))
	fmt.Printf("f 是否是数字：%v \n", IsNumeric(f))
}
