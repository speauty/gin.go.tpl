package code

import "fmt"

type Code int

const (
	StdOk  Code = 0
	StdErr Code = 10000
)

var codeMsg = map[Code]string{
	StdOk:  "请求成功",
	StdErr: "当前服务异常, 请稍后请求",
}

func (code Code) GetMsg() string {
	if _, ok := codeMsg[code]; !ok {
		return fmt.Sprintf("当前编码[%d]未设置, 请及时修复", code)
	}
	return codeMsg[code]
}
