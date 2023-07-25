// Copyright 2021 OpenIM Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package friend

import (
	"errors"
	ws "open_im_sdk/internal/interaction"
	"open_im_sdk/internal/user"
	"open_im_sdk/open_im_sdk_callback"
	"open_im_sdk/pkg/common"
	"open_im_sdk/pkg/db/db_interface"
	"open_im_sdk/pkg/db/model_struct"
	"open_im_sdk/pkg/log"
	"open_im_sdk/pkg/utils"
)

type Friend struct {
	loginUserID    string
	db             db_interface.DataBase
	user           *user.User
	p              *ws.PostApi
	loginTime      int64
	conversationCh chan common.Cmd2Value

	listenerForService open_im_sdk_callback.OnListenerForService
}

func (f *Friend) Db() db_interface.DataBase {
	return f.db
}

func NewFriend(loginUserID string, db db_interface.DataBase, user *user.User, p *ws.PostApi, conversationCh chan common.Cmd2Value) *Friend {
	return &Friend{loginUserID: loginUserID, db: db, user: user, p: p, conversationCh: conversationCh}
}

func (f *Friend) SetListenerForService(listener open_im_sdk_callback.OnListenerForService) {
	f.listenerForService = listener
}

func (f *Friend) GetUserNameAndFaceUrlByUid(friendUserID, operationID string) (faceUrl, name string, err error, isFromSvr bool) {
	isFromSvr = false
	friendInfo, err := f.db.GetFriendInfoByFriendUserID(friendUserID)
	if err == nil {
		if friendInfo.Remark != "" {
			return friendInfo.FaceURL, friendInfo.Remark, nil, isFromSvr
		} else {
			return friendInfo.FaceURL, friendInfo.Nickname, nil, isFromSvr
		}
	} else {
		if operationID == "" {
			operationID = utils.OperationIDGenerator()
		}
		userInfos, err := f.user.GetUsersInfoFromSvrNoCallback([]string{friendUserID}, operationID)
		if err != nil {
			return "", "", err, isFromSvr
		}
		for _, v := range userInfos {
			isFromSvr = true
			return v.FaceURL, v.Nickname, nil, isFromSvr
		}
		log.Info(operationID, "GetUsersInfoFromSvr ", friendUserID)
	}
	return "", "", errors.New("getUserNameAndFaceUrlByUid err"), isFromSvr
}

func (f *Friend) GetDesignatedFriendListInfo(callback open_im_sdk_callback.Base, friendUserIDList []string, operationID string) []*model_struct.LocalFriend {
	friendList, err := f.db.GetFriendInfoList(friendUserIDList)
	common.CheckDBErrCallback(callback, err, operationID)
	return friendList
}

func (f *Friend) GetDesignatedBlackListInfo(callback open_im_sdk_callback.Base, blackIDList []string, operationID string) []*model_struct.LocalBlack {
	blackList, err := f.db.GetBlackInfoList(blackIDList)
	common.CheckDBErrCallback(callback, err, operationID)
	return blackList
}
