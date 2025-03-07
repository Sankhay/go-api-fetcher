package models

type NetworkError struct {
	Code int
	Msg  string
}

func (n *NetworkError) Error() string {
	return n.Msg
}
