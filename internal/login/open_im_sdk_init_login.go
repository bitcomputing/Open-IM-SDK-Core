package login

import (
	"open_im_sdk/open_im_sdk_callback"
	"open_im_sdk/pkg/log"
)

func (u *LoginMgr) Login(callback open_im_sdk_callback.Base, userID, token string, operationID string) {
	go func() {
		u.login(userID, token, callback, operationID)
	}()
}

func (u *LoginMgr) WakeUp(callback open_im_sdk_callback.Base, operationID string) {
	go func() {
		u.wakeUp(callback, operationID)
	}()
}

func (u *LoginMgr) Logout(callback open_im_sdk_callback.Base, operationID string) {
	if callback == nil {
		u.logout(callback, operationID)
		return
	}
	go func() {
		u.logout(callback, operationID)
	}()
}

func (u *LoginMgr) SetAppBackgroundStatus(callback open_im_sdk_callback.Base, isBackground bool, operationID string) {
	go func() {
		log.NewInfo(operationID, "SetAppBackgroundStatus", "isBackground", isBackground)
		u.setAppBackgroundStatus(callback, isBackground, operationID)
	}()
}
