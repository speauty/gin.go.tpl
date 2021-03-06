package code

import "fmt"

type Code int

const (
	StdOk Code = 200

	StdErr               Code = 10000
	StdRouteNotFound     Code = 10001
	StdRequestRateExceed Code = 10002
	StdInput             Code = 10010
	StdParam             Code = 10020

	StdDb         Code = 10030
	StdDbQuery    Code = 10031
	StdDbUpdate   Code = 10032
	StdDbDelete   Code = 10033
	StdDbUnique   Code = 10034
	StdDbNotFound Code = 10035
)

var codeMsg = map[Code]string{
	StdOk:                "请求成功",
	StdErr:               "当前服务异常, 请稍后重试",
	StdRouteNotFound:     "当前路由未定义, 请检测请求地址是否正确",
	StdRequestRateExceed: "当前访问人数过多, 请稍后重试",
	StdInput:             "当前输入数据验证失败, 请及时检测数据是否合法",
	StdParam:             "当前参数非法",
	StdDb:                "当前数据库操作异常, 请稍后重试",
	StdDbQuery:           "数据查询异常",
	StdDbUpdate:          "数据更新异常",
	StdDbDelete:          "数据删除异常",
	StdDbUnique:          "当前数据已存在",
	StdDbNotFound:        "当前数据不存在",
}

func (code Code) Trans() string {
	if _, ok := codeMsg[code]; !ok {
		return fmt.Sprintf("当前编码[%d]未设置, 请及时修复", code)
	}
	return codeMsg[code]
}
