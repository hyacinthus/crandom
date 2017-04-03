package main

type errMsg struct {
	code    string
	message string
}

func newErrMsg(code string, msg string) errMsg {
	return errMsg{
		code:    code,
		message: msg,
	}
}
