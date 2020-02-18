package main

import (
	"fmt"
	"regexp"
)

func main() {
	// /* 基础 */
	// true, err := regexp.Match("AB", []byte("AB"))               //连接
	// ture, err = regexp.Match("AB|ACD", []byte("ACD"))           //或
	// true, err = regexp.Match("AB*", []byte("ABBBBBBBB"))        //*前面的一位重复N次，包括0.
	// true, err = regexp.Match("(AB)*", []byte("ABABABABABABBB")) //()将内部试做一个整体
	// true, err = regexp.Match("A(B|CD)*", []byte("A"))           //整合，ACDCD，A，ABBB，AB

	// /* 通配符 */
	// true, err = regexp.Match("A.B", []byte("ACB"))   //.通配符表示任意一个字符
	// false, err = regexp.Match("A.B", []byte("ACDB")) //但是只能表示一个，并且和周围元素形成连接关系。
	// true, err = regexp.Match("A.*B", []byte("ACDB")) //N个通配符

	// /* []表示范围 */
	// true, err = regexp.Match("[abc]", []byte("a"))         //abc中的一个
	// true, err = regexp.Match("[a-z][0-9]", []byte("a1"))   //a-z中的一个加0-9中的一个
	// false, err = regexp.Match("[^a-z][0-9]", []byte("a1")) //a-z的补集加0-9中的一个

	// /* 其他 */
	// true, err = regexp.Match("(ab)+", []byte("ab"))       //+表示一个或多个
	// true, err = regexp.Match("(ab)?", []byte(""))         //？表示0个或1个
	// true, err = regexp.Match("(ab){3}", []byte("ababab")) //{3}三个
	// true, err = regexp.Match("(ab){0,3}", []byte(""))     //{0,3}0-3个

	matched, _ := regexp.MatchString("(ab)", "ab")
	fmt.Println(matched)
	reg := regexp.MustCompile("(ab)|(cd)")
	bs := reg.Find([]byte("cdabcd"))
	fmt.Println(bs)
	bss := reg.FindAll([]byte("abcdab"), -1)
	fmt.Println(bss)
	index := reg.FindIndex([]byte("abcdab"))
	fmt.Println(index)
	indexs := reg.FindAllIndex([]byte("abcdab"), -1)
	fmt.Println(indexs)
	str := reg.FindString("abcdab")
	fmt.Println(str)
	strs := reg.FindAllString("abcdab", -1)
	fmt.Println(strs)
}
