package bitgo

import (
	"time"
)

type UserProfile struct {
	User struct {
		ID       string `json:"id"`
		IsActive bool   `json:"isActive"`
		Name     struct {
			First string `json:"first"`
			Full  string `json:"full"`
			Last  string `json:"last"`
		} `json:"name"`
		Username string `json:"username"`
	} `json:"user"`
}

type Token struct {
	TokenType string   `json:"token_type"`
	ExpiresIn int      `json:"expires_in"`
	ExpiresAt int      `json:"expires_at"`
	Scope     []string `json:"scope"`
	User      struct {
		ID       string `json:"id"`
		Username string `json:"username"`
		Name     struct {
			Full  string `json:"full"`
			First string `json:"first"`
			Last  string `json:"last"`
		} `json:"name"`
		Email struct {
			Email    string `json:"email"`
			Verified bool   `json:"verified"`
		} `json:"email"`
		Phone struct {
			Phone    string `json:"phone"`
			Verified bool   `json:"verified"`
		} `json:"phone"`
		Identity struct {
			Civic struct {
				State string `json:"state"`
			} `json:"civic"`
			Verified bool `json:"verified"`
		} `json:"identity"`
		OtpDevices []struct {
			ID       string `json:"id"`
			Type     string `json:"type"`
			Label    string `json:"label"`
			Verified bool   `json:"verified"`
		} `json:"otpDevices"`
		RateLimits struct {
		} `json:"rateLimits"`
		DisableReset2FA bool `json:"disableReset2FA"`
		Currency        struct {
			Currency    string `json:"currency"`
			BitcoinUnit string `json:"bitcoinUnit"`
		} `json:"currency"`
		Timezone     string `json:"timezone"`
		IsActive     bool   `json:"isActive"`
		EcdhKeychain string `json:"ecdhKeychain"`
		Referrer     struct {
		} `json:"referrer"`
		Apps struct {
			Coinbase struct {
			} `json:"coinbase"`
		} `json:"apps"`
		ForceResetPassword bool          `json:"forceResetPassword"`
		AllowedCoins       []interface{} `json:"allowedCoins"`
		Agreements         struct {
			TermsOfUse               int       `json:"termsOfUse"`
			TermsOfUseAcceptanceDate time.Time `json:"termsOfUseAcceptanceDate"`
		} `json:"agreements"`
		LastLogin   time.Time `json:"lastLogin"`
		Enterprises []struct {
			ID string `json:"id"`
		} `json:"enterprises"`
	} `json:"user"`
	EncryptedToken    string `json:"encryptedToken"`
	DerivationPath    string `json:"derivationPath"`
	EncryptedECDHXprv string `json:"encryptedECDHXprv"`
	AccessToken       string `json:"access_token"`
}

type Session struct {
	Client  string    `json:"client"`
	User    string    `json:"user"`
	Scope   []string  `json:"scope"`
	Expires time.Time `json:"expires"`
	Origin  string    `json:"origin"`
	Unlock  struct {
		Time    time.Time `json:"time"`
		Expires time.Time `json:"expires"`
		TxCount int       `json:"txCount"`
		TxValue int       `json:"txValue"`
	} `json:"unlock"`
}

// Current User Profile

func (b *BitGo) CurrentUserProfile() (userProfile UserProfile, err error) {
	err = b.get("user/me",
		nil,
		&userProfile)
	return
}

// Session Information

func (b *BitGo) SessionInformation() (session Session, err error) {
	err = b.get("user/session",
		nil,
		&session)
	return
}

// Lock

func (b *BitGo) Lock() (err error) {
	err = b.post("user/lock",
		nil,
		nil)
	return
}

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

// Login

type LoginParams struct {
	Email      string `json:"email" valid:"required"`
	Password   string `json:"password" valid:"required"`
	Otp        string `json:"otp" valid:"required"`
	Extensible bool   `json:"extensible,omitempty"`
}

func (b *BitGo) Login(params LoginParams) (token Token, err error) {
	b.token = ""

	err = b.post("user/login",
		params,
		&token)
	if err != nil {
		return
	}

	b.token = token.AccessToken

	return
}

//  Logout

func (b *BitGo) Logout() (err error) {
	err = b.get("user/logout",
		nil,
		nil)
	if err != nil {
		b.token = ""
	}
	return
}
