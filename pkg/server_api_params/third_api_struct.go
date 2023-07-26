package server_api_params

type FcmUpdateTokenReq struct {
	OperationID string `json:"operationID" binding:"required"`
	Platform    int    `json:"platform" binding:"required,min=1,max=2"` //only for ios + android
	FcmToken    string `json:"fcmToken" binding:"required"`
}

type FcmUpdateTokenResp struct {
	CommResp
}

type SetAppBadgeReq struct {
	OperationID    string `json:"operationID" binding:"required"`
	FromUserID     string `json:"fromUserID" binding:"required"`
	AppUnreadCount int32  `json:"appUnreadCount" binding:"required"`
}

type SetAppBadgeResp struct {
	CommResp
}
