package main

import (
	"os/exec"
	"testing"
)

func TestGjsonCLI(t *testing.T) {
	tests := []struct {
		jsonStr string
		path    string
		want    string
		name    string
		indent  bool
	}{
		{
			jsonStr: `{"name":"张三","age":18,"address":{"city":"北京"},"married":true,"children":null,"scores":[100,99],"info":{"height":180}}`,
			path:    "address.city",
			want:    "北京\n",
			name:    "字符串类型",
		},
		{
			jsonStr: `{"age":18}`,
			path:    "age",
			want:    "18\n",
			name:    "数字类型",
		},
		{
			jsonStr: `{"married":true}`,
			path:    "married",
			want:    "true\n",
			name:    "布尔类型true",
		},
		{
			jsonStr: `{"married":false}`,
			path:    "married",
			want:    "false\n",
			name:    "布尔类型false",
		},
		{
			jsonStr: `{"children":null}`,
			path:    "children",
			want:    "null\n",
			name:    "null类型",
		},
		{
			jsonStr: `{"scores":[100,99]}`,
			path:    "scores",
			want:    "[100,99]\n",
			name:    "数组类型",
		},
		{
			jsonStr: `{"info":{"height":180}}`,
			path:    "info",
			want:    "{\"height\":180}\n",
			name:    "对象类型",
		},
		{
			jsonStr: `{"info":{"height":180,"weight":70}}`,
			path:    "info",
			want:    "{\n  \"height\": 180,\n  \"weight\": 70\n}\n",
			name:    "对象类型indent美化",
			indent:  true,
		},
		{
			jsonStr: `{"arr":[1,2,3]}`,
			path:    "arr",
			want:    "[\n  1,\n  2,\n  3\n]\n",
			name:    "数组类型indent美化",
			indent:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := []string{"run", "main.go", "-json", tt.jsonStr, "-path", tt.path}
			if tt.indent {
				args = append(args, "-indent")
			}
			cmd := exec.Command("go", args...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("命令执行失败: %v, 输出: %s", err, string(output))
			}
			if string(output) != tt.want {
				t.Errorf("期望输出: %q, 实际输出: %q", tt.want, string(output))
			}
		})
	}
}
