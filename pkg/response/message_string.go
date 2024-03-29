// Code generated by "stringer -linecomment -output message_string.go -type=RespCode"; DO NOT EDIT.

package response

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SUCCESS-0]
	_ = x[ErrorServer - -1]
	_ = x[ErrorParams-4000]
	_ = x[ErrorAuthCheckTokenFail-4001]
	_ = x[ErrorAuthCheckTokenTimeout-4002]
	_ = x[ErrorAuthToken-4003]
	_ = x[ErrorAuth-4004]
	_ = x[ErrorRequestTooFrequent-4005]
}

const (
	_RespCode_name_0 = "操作失败操作成功"
	_RespCode_name_1 = "请求参数错误Token鉴权失败Token已超时Token生成失败Token错误请求太频繁"
)

var (
	_RespCode_index_0 = [...]uint8{0, 12, 24}
	_RespCode_index_1 = [...]uint8{0, 18, 35, 49, 66, 77, 92}
)

func (i RespCode) String() string {
	switch {
	case -1 <= i && i <= 0:
		i -= -1
		return _RespCode_name_0[_RespCode_index_0[i]:_RespCode_index_0[i+1]]
	case 4000 <= i && i <= 4005:
		i -= 4000
		return _RespCode_name_1[_RespCode_index_1[i]:_RespCode_index_1[i+1]]
	default:
		return "RespCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
