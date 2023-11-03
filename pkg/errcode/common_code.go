package errcode

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(100, "服务内部错误")
	InvalidParams             = NewError(101, "入参错误")
	NotFound                  = NewError(102, "找不到")
	UnauthorizedAuthNotExist  = NewError(103, "鉴权失败，找不到对应的 AppKey 和 AppSecret")
	UnauthorizedTokenError    = NewError(104, "鉴权失败，Token 错误")
	UnauthorizedTokenTimeout  = NewError(105, "鉴权失败，Token 超时")
	UnauthorizedTokenGenerate = NewError(106, "鉴权失败，Token 生成失败")
	TooManyRequests           = NewError(107, "请求过多")
)
