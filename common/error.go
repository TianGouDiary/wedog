package common

import "fmt"

type ErrConst struct {
	Code uint32
	Msg  string
}

func (err ErrConst) String() string {
	return fmt.Sprintf("errCode:%d, errMsg:%s", err.Code, err.Msg)
}

func (err ErrConst) ErrorCode() uint32 {
	return err.Code
}

func (err ErrConst) ErrorMsg() string {
	return err.Msg
}

var (
	ERR_SUCCESS = ErrConst{0, "SUCCESS"}

	//通用错误
	ERR_PARAM_IN = ErrConst{1000, "入参非法"}
	ERR_COMMON   = ErrConst{1001, "通用错误"}

	//db相关
	ERR_DB_FAIL               = ErrConst{2000, "操作DB异常"}
	ERR_DB_NO_ROW_AFFECT      = ErrConst{2001, "数据库操作没有受影响的行数"}
	ERR_DB_NO_DATA            = ErrConst{2002, "数据库未查询到数据"}
	ERR_DB_NOT_ALL_ROW_AFFECT = ErrConst{2003, "数据库操作未影响到所有行"}
)
