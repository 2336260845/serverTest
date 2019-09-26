package apis

type ResponeMsg struct {
	Ok   string      `json:"ok"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SendJson(err error, data interface{}) (responeMsg ResponeMsg) {
	if err != nil {
		responeMsg.Ok = "success"
		responeMsg.Msg = err.Error()
	} else {
		responeMsg.Ok = "failed"
	}

	responeMsg.Data = data

	return
}
