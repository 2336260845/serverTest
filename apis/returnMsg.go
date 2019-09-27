package apis

type ResponeMsg struct {
	Ok   string      `json:"ok"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SendJson(err error, data interface{}) (responeMsg ResponeMsg) {
	if err == nil {
		responeMsg.Ok = "success"
	} else {
		responeMsg.Ok = "failed"
		responeMsg.Msg = err.Error()
	}

	responeMsg.Data = data

	return
}
