package models

type HttpError struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
}

func (n *HttpError) Error() string {
	return n.Msg
}
