package main
import (
	"fmt"
	"strings"
)
func main () {
	TestCompare()
	TestEqualFold()
	TestRepeat()
	TestReplace()
	TestJoin()
}
//按字典顺序比较a和b字符串大小
func TestCompare() {
	fmt.Println(strings.Compare("abc","bcd"))
	fmt.Println("abc" < "bcd")
}
//判断s和t两个utf-8字符串是否相等，忽略大小写
func TestEqualFold() {
	fmt.Println(strings.EqualFold("Go","go"))
}
//将字符串s重复count次返回
func TestRepeat() {
	fmt.Println("g" + strings.Repeat("o",8) + "le")
}
//替换字符串s中old字符为new字符并返回，n<0时替换所有old字符串
func TestReplace() {
	fmt.Println(strings.Replace("王老大 王老二 王老三", "王", "张", 2))
	fmt.Println(strings.Replace("王老大 王老二 王老三", "王", "张", -1))
}
//将a中的所有字符串连接成一个字符串，使用字符串sep作为分隔符
func TestJoin() {
	s := []string{"abc", "ABC", "123"}
	fmt.Println(strings.Join(s, ","))
	fmt.Println(strings.Join(s,""))
}
