// func Match(pattern, name string) (matched bool, err error)
// 参数列表
// pattern 匹配的模式
// name 原始字符串
// 返回值：
// 返回bool error
// 功能说明：
// 这个函数主要是文件名匹配，只有完全匹配则返回true，nil
// 官方文档：
// Match reports whether name matches the shell file name pattern. The pattern syntax is:
// pattern:
// 	{ term }
// term:
// 	'*'         matches any sequence of non-/ characters
// 	'?'         matches any single non-/ character
// 	'[' [ '^' ] { character-range } ']'
// 	            character class (must be non-empty)
// 	c           matches character c (c != '*', '?', '\\', '[')
// 	'\\' c      matches character c

// character-range:
// 	c           matches character c (c != '\\', '-', ']')
// 	'\\' c      matches character c
// 	lo '-' hi   matches character c for lo <= c <= hi
// Match requires pattern to match all of name, not just a substring.
// The only possible returned error is ErrBadPattern, when pattern is malformed.

package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Match("*", "alll"))  //true nil
	fmt.Println(path.Match("*", "a/lll")) //false nil
	fmt.Println(path.Match("?", "alll"))  //false nil
	fmt.Println(path.Match("?", "a"))     //true nil
	fmt.Println(path.Match("a", "a"))     //true nil
	fmt.Println(path.Match("\\a", "a"))   //true nil
}
