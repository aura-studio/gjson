package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tidwall/gjson"
)

func main() {
	jsonStr := flag.String("json", "", "要解析的JSON字符串")
	path := flag.String("path", "", "gjson路径表达式")
	flag.Parse()

	if *jsonStr == "" || *path == "" {
		fmt.Fprintln(os.Stderr, "用法: man -json '<json字符串>' -path '<gjson路径>'")
		os.Exit(1)
	}

	result := gjson.Get(*jsonStr, *path)
	switch {
	case result.Type == gjson.String:
		fmt.Println(result.String())
	case result.Type == gjson.Number:
		fmt.Println(result.Num)
	case result.Type == gjson.True:
		fmt.Println(true)
	case result.Type == gjson.False:
		fmt.Println(false)
	case result.Type == gjson.JSON:
		fmt.Println(result.Raw)
	case result.Type == gjson.Null:
		fmt.Println("null")
	default:
		fmt.Println(result.String())
	}
}
