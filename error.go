package bitgo

type Error struct {
	Message   string `json:"error"`
	RequestId string `json:"requestId"`
	Name      string `json:"name"`
}

func (e Error) Error() string {
	return e.Message
}
