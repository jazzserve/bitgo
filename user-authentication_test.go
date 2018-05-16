package bitgo

import (
	"testing"
)

func testLogin(t *testing.T, b *BitGo, params *TestParams) {
	_, err := b.Login(LoginParams{
		Email:    params.Email,
		Otp:      "000000",
		Password: params.Passphrase,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

// Current User Profile

func TestCurrentUserProfile(t *testing.T) {
	b, params := getTestBitGo(t)

	testLogin(t, b, params)

	_, err := b.CurrentUserProfile()
	if err != nil {
		t.Fatal(err.Error())
	}
}

// Session Information

func TestSessionInformation(t *testing.T) {
	b, params := getTestBitGo(t)

	testLogin(t, b, params)

	_, err := b.SessionInformation()
	if err != nil {
		t.Fatal(err.Error())
	}
}

// Lock

func TestLock(t *testing.T) {
	b, _ := getTestBitGo(t)

	err := b.Lock()
	if err != nil {
		t.Fatal(err.Error())
	}
}

// Unlock

func TestUnlock(t *testing.T) {
	b, _ := getTestBitGo(t)

	err := b.Unlock(UnlockParams{
		Otp: "000000",
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

// Login

func TestLogin(t *testing.T) {
	b, params := getTestBitGo(t)

	testLogin(t, b, params)
}

//  Logout

func TestLogout(t *testing.T) {
	b, params := getTestBitGo(t)

	testLogin(t, b, params)

	err := b.Logout()
	if err != nil {
		t.Fatal(err.Error())
	}
}
