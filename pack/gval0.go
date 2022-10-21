package pack

import (
	"fmt"
	"os"
	"strings"

	"mvdan.cc/sh/v3/syntax"
)

func syntax_test() {
	r := strings.NewReader("echo 'rm -f'") //字符串读入Reader结构体中
	//fmt.Println(reflect.TypeOf(r))
	f, err := syntax.NewParser().Parse(r, "") //分配一个新的解析器
	if err != nil {
		return
	}
	syntax.NewPrinter().Print(os.Stdout, f) //打印语法树node
	//syntax.DebugPrint(os.Stdout, f) //打印语法树
}

func quote_test() {
	//Quote 返回输入字符串的带引号版本，
	//Quote returns a quoted version of the input string, so that the quoted version is expanded
	//or interpreted as the original string in the given language variant.
	//An error is returned when a string cannot be quoted for a variant

	//Some strings do not require any quoting and are returned unchanged.
	//Those strings can be directly surrounded in single quotes as well.

	for _, s := range []string{
		"foo", //定义字符串切片
		"bar $baz",
		`"won't"`,
		"~/home",
		"#1304",
		"name=value",
		"for",
		"glob-*",
		"invalid-\xe2'",
		"nonprint-\x0b\x1b",

		"cd", `"cd"`, //字符串中含单引号返回"字符串"
		"poweroff", "poweroff -f",
		"echo `poweroff`", `echo "poweroff"`,
	} {
		quoted, err := syntax.Quote(s, syntax.LangBash)
		if err != nil {
			fmt.Printf("%q cannot be quoted: %v\n", s, err)
		} else {
			fmt.Printf("Quote(%17q): %s\n", s, quoted)
		}
	}

}
