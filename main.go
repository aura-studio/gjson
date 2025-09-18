package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/tidwall/gjson"
)

func main() {
	jsonStr := flag.String("json", "", "要解析的JSON字符串")
	path := flag.String("path", "", "gjson路径表达式")
	indent := flag.Bool("indent", false, "如果为JSON则美化输出")
	flag.Parse()

	if *jsonStr == "" || *path == "" {
		fmt.Fprintln(os.Stderr, "用法: man -json '<json字符串>' -path '<gjson路径>' [-indent]")
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
		if *indent {
			prettyPrint(result.Raw)
		} else {
			fmt.Println(result.Raw)
		}
	case result.Type == gjson.Null:
		fmt.Println("null")
	default:
		fmt.Println(result.String())
	}
}

func prettyPrint(raw string) {
	var out []byte
	var err error
	out, err = indentJSON(raw)
	if err != nil {
		fmt.Println(raw)
		return
	}
	fmt.Println(string(out))
}

func indentJSON(raw string) ([]byte, error) {
	// 只处理对象或数组
	if len(raw) == 0 || (raw[0] != '{' && raw[0] != '[') {
		return []byte(raw), nil
	}
	var obj interface{}
	err := json.Unmarshal([]byte(raw), &obj)
	if err != nil {
		return nil, err
	}
	return json.MarshalIndent(obj, "", "  ")
}
