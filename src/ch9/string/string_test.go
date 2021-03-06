package string_test

import "testing"

func TestString(t *testing.T) {
	var s string
	t.Log(s) // 初始化为默认值""
	
	s = "hello"
	t.Log(len(s))
	// s[1] = '3'
	// string 是不可变的 byte slice

	s = "\xE4\xB8\xA5"
	t.Log(s)
	t.Log(len(s))

	s = "\xE4\xB8\xAD"
	t.Log(s)

	s = "中"
	t.Log(len(s)) // byte 数

	c := []rune(s)
	t.Log("rune c", len(c))

	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}


func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _,  c:= range s {
		t.Logf("%[1]c, %[1]d, %[1]x", c)
	}
}