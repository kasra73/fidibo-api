package e

var MsgFlags = map[int]string{
	SUCCESS:                     "ok",
	ERROR:                       "fail",
	INVALID_PARAMS:              "پارامترهای نادرست",
	ERROR_KEYWORD_EMPTY:         "کلمه عبور نباید خالی باشد",
	ERROR_KEYWORD_SEARCH_FAILED: "در جست و جو خطایی رخ داده است",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
