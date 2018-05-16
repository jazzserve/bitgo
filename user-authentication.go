package bitgo

// TODO

// Unlock

type UnlockParams struct {
	Otp      string `json:"otp" valid:"required"`
	Duration int    `json:"duration,omitempty"`
}

func (b *BitGo) Unlock(params UnlockParams) (err error) {
	err = b.post("user/unlock",
		params,
		nil)
	return
}
