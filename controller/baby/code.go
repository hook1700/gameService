/**
 @author:way
 @date:2021/12/16
 @note
**/

package baby

type Recode int64

const (
	CodeSuccess      Recode = 200
	CodeInvalidParam Recode = 501
	CodeServerBusy   Recode = 500
)

var codeMsgMap = map[Recode]string{
	CodeSuccess:      "ok",
	CodeInvalidParam: "请求参数错误",
	CodeServerBusy:   "服务器繁忙",

}

func (c Recode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}

