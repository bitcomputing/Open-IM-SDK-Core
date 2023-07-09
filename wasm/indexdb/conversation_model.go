// Copyright © 2023 OpenIM SDK. All rights reserved.
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

//go:build js && wasm
// +build js,wasm

package indexdb

import (
	"context"
	"open_im_sdk/pkg/db/model_struct"
	"open_im_sdk/pkg/utils"
	"open_im_sdk/wasm/indexdb/temp_struct"
)

type LocalConversations struct {
}

func NewLocalConversations() *LocalConversations {
	return &LocalConversations{}
}

func (i *LocalConversations) GetAllConversationListDB(ctx context.Context) (result []*model_struct.LocalConversation, err error) {
	cList, err := Exec()
	if err != nil {
		return nil, err
	} else {
		if v, ok := cList.(string); ok {
			var temp []model_struct.LocalConversation
			err := utils.JsonStringToStruct(v, &temp)
			if err != nil {
				return nil, err
			}
			for _, v := range temp {
				v1 := v
				result = append(result, &v1)
			}
			return result, err
		} else {
			return nil, ErrType
		}
	}
}

func (i *LocalConversations) GetConversation(ctx context.Context, conversationID string) (*model_struct.LocalConversation, error) {
	c, err := Exec(conversationID)
	if err != nil {
		return nil, err
	} else {
		if v, ok := c.(string); ok {
			result := model_struct.LocalConversation{}
			err := utils.JsonStringToStruct(v, &result)
			if err != nil {
				return nil, err
			}
			return &result, err
		} else {
			return nil, ErrType
		}
	}
}

func (i *LocalConversations) GetHiddenConversationList(ctx context.Context) (result []*model_struct.LocalConversation, err error) {
	cList, err := Exec()
	if err != nil {
		return nil, err
	} else {
		if v, ok := cList.(string); ok {
			var temp []model_struct.LocalConversation
			err := utils.JsonStringToStruct(v, &temp)
			if err != nil {
				return nil, err
			}
			for _, v := range temp {
				v1 := v
				result = append(result, &v1)
			}
			return result, err
		} else {
			return nil, ErrType
		}
	}
}
func (i *LocalConversations) GetAllConversations(ctx context.Context) (result []*model_struct.LocalConversation, err error) {
	cList, err := Exec()
	if err != nil {
		return nil, err
	} else {
		if v, ok := cList.(string); ok {
			var temp []model_struct.LocalConversation
			err := utils.JsonStringToStruct(v, &temp)
			if err != nil {
				return nil, err
			}
			for _, v := range temp {
				v1 := v
				result = append(result, &v1)
			}
			return result, err
		} else {
			return nil, ErrType
		}
	}
}
func (i *LocalConversations) UpdateColumnsConversation(ctx context.Context, conversationID string, args map[string]interface{}) error {
	_, err := Exec(conversationID, utils.StructToJsonString(args))
	return err
}
func (i *LocalConversations) GetConversationByUserID(ctx context.Context, userID string) (*model_struct.LocalConversation, error) {
	c, err := Exec(userID)
	if err != nil {
		return nil, err
	} else {
		if v, ok := c.(string); ok {
			result := model_struct.LocalConversation{}
			err := utils.JsonStringToStruct(v, &result)
			if err != nil {
				return nil, err
			}
			return &result, err
		} else {
			return nil, ErrType
		}
	}
}

func (i *LocalConversations) GetConversationListSplitDB(ctx context.Context, offset, count int) (result []*model_struct.LocalConversation, err error) {
	cList, err := Exec(offset, count)
	if err != nil {
		return nil, err
	} else {
		if v, ok := cList.(string); ok {
			var temp []model_struct.LocalConversation
			err := utils.JsonStringToStruct(v, &temp)
			if err != nil {
				return nil, err
			}
			for _, v := range temp {
				v1 := v
				result = append(result, &v1)
			}
			return result, err
		} else {
			return nil, ErrType
		}
	}
}

func (i *LocalConversations) BatchInsertConversationList(ctx context.Context, conversationList []*model_struct.LocalConversation) error {
	_, err := Exec(utils.StructToJsonString(conversationList))
	return err
}

func (i *LocalConversations) InsertConversation(ctx context.Context, conversationList *model_struct.LocalConversation) error {
	_, err := Exec(utils.StructToJsonString(conversationList))
	return err
}

func (i *LocalConversations) DeleteConversation(ctx context.Context, conversationID string) error {
	_, err := Exec(conversationID)
	return err
}

func (i *LocalConversations) UpdateConversation(ctx context.Context, c *model_struct.LocalConversation) error {
	if c.ConversationID == "" {
		return PrimaryKeyNull
	}
	tempLocalConversation := temp_struct.LocalConversation{
		ConversationType:      c.ConversationType,
		UserID:                c.UserID,
		GroupID:               c.GroupID,
		ShowName:              c.ShowName,
		FaceURL:               c.FaceURL,
		RecvMsgOpt:            c.RecvMsgOpt,
		UnreadCount:           c.UnreadCount,
		GroupAtType:           c.GroupAtType,
		LatestMsg:             c.LatestMsg,
		LatestMsgSendTime:     c.LatestMsgSendTime,
		DraftText:             c.DraftText,
		DraftTextTime:         c.DraftTextTime,
		IsPinned:              c.IsPinned,
		IsPrivateChat:         c.IsPrivateChat,
		BurnDuration:          c.BurnDuration,
		IsNotInGroup:          c.IsNotInGroup,
		UpdateUnreadCountTime: c.UpdateUnreadCountTime,
		AttachedInfo:          c.AttachedInfo,
		Ex:                    c.Ex,
	}
	_, err := Exec(c.ConversationID, utils.StructToJsonString(tempLocalConversation))
	return err
}

func (i *LocalConversations) UpdateConversationForSync(ctx context.Context, c *model_struct.LocalConversation) error {
	if c.ConversationID == "" {
		return PrimaryKeyNull
	}
	tempLocalConversation := temp_struct.LocalPartConversation{
		RecvMsgOpt:            c.RecvMsgOpt,
		GroupAtType:           c.GroupAtType,
		IsPinned:              c.IsPinned,
		IsPrivateChat:         c.IsPrivateChat,
		IsNotInGroup:          c.IsNotInGroup,
		UpdateUnreadCountTime: c.UpdateUnreadCountTime,
		BurnDuration:          c.BurnDuration,
		AttachedInfo:          c.AttachedInfo,
		Ex:                    c.Ex,
	}
	_, err := Exec(c.ConversationID, utils.StructToJsonString(tempLocalConversation))
	return err
}

func (i *LocalConversations) BatchUpdateConversationList(ctx context.Context, conversationList []*model_struct.LocalConversation) error {
	for _, v := range conversationList {
		err := i.UpdateConversation(ctx, v)
		if err != nil {
			return utils.Wrap(err, "BatchUpdateConversationList failed")
		}

	}
	return nil
}

func (i *LocalConversations) ConversationIfExists(ctx context.Context, conversationID string) (bool, error) {
	seq, err := Exec(conversationID)
	if err != nil {
		return false, err
	} else {
		if v, ok := seq.(bool); ok {
			return v, err
		} else {
			return false, ErrType
		}
	}
}

func (i *LocalConversations) ResetConversation(ctx context.Context, conversationID string) error {
	_, err := Exec(conversationID)
	return err
}

func (i *LocalConversations) ResetAllConversation(ctx context.Context) error {
	_, err := Exec()
	return err
}

func (i *LocalConversations) ClearConversation(ctx context.Context, conversationID string) error {
	_, err := Exec(conversationID)
	return err
}

func (i *LocalConversations) ClearAllConversation(ctx context.Context) error {
	_, err := Exec()
	return err
}

func (i *LocalConversations) SetConversationDraftDB(ctx context.Context, conversationID, draftText string) error {
	_, err := Exec(conversationID, draftText)
	return err
}

func (i *LocalConversations) RemoveConversationDraft(ctx context.Context, conversationID, draftText string) error {
	_, err := Exec(conversationID, draftText)
	return err
}

func (i *LocalConversations) UnPinConversation(ctx context.Context, conversationID string, isPinned int) error {
	_, err := Exec(conversationID, isPinned)
	return err
}

func (i *LocalConversations) UpdateAllConversation(ctx context.Context, conversation *model_struct.LocalConversation) error {
	_, err := Exec()
	return err
}

func (i *LocalConversations) IncrConversationUnreadCount(ctx context.Context, conversationID string) error {
	_, err := Exec(conversationID)
	return err
}
func (i *LocalConversations) DecrConversationUnreadCount(ctx context.Context, conversationID string, count int64) error {
	_, err := Exec(conversationID, count)
	return err
}
func (i *LocalConversations) GetTotalUnreadMsgCountDB(ctx context.Context) (totalUnreadCount int32, err error) {
	count, err := Exec()
	if err != nil {
		return 0, err
	} else {
		if v, ok := count.(float64); ok {
			var result int32
			result = int32(v)
			return result, err
		} else {
			return 0, ErrType
		}
	}
}

func (i *LocalConversations) SetMultipleConversationRecvMsgOpt(ctx context.Context, conversationIDList []string, opt int) (err error) {
	_, err = Exec(utils.StructToJsonString(conversationIDList), opt)
	return err
}

func (i *LocalConversations) GetMultipleConversationDB(ctx context.Context, conversationIDList []string) (result []*model_struct.LocalConversation, err error) {
	cList, err := Exec(utils.StructToJsonString(conversationIDList))
	if err != nil {
		return nil, err
	} else {
		if v, ok := cList.(string); ok {
			var temp []model_struct.LocalConversation
			err := utils.JsonStringToStruct(v, &temp)
			if err != nil {
				return nil, err
			}
			for _, v := range temp {
				v1 := v
				result = append(result, &v1)
			}
			return result, err
		} else {
			return nil, ErrType
		}
	}
}

func (i *LocalConversations) GetAllSingleConversationIDList(ctx context.Context) (result []string, err error) {
	conversationIDs, err := Exec()
	if err != nil {
		return nil, err
	} else {
		if v, ok := conversationIDs.(string); ok {
			err := utils.JsonStringToStruct(v, &result)
			if err != nil {
				return nil, err
			}
			return result, nil
		} else {
			return nil, ErrType
		}
	}
}

func (i *LocalConversations) GetAllConversationIDList(ctx context.Context) ([]string, error) {
	conversationIDList, err := Exec()
	if err != nil {
		return nil, err
	} else {
		if v, ok := conversationIDList.(string); ok {
			var result []string
			err := utils.JsonStringToStruct(v, &result)
			if err != nil {
				return nil, err
			}
			return result, err
		} else {
			return nil, ErrType
		}
	}
}

func (i *LocalConversations) UpdateOrCreateConversations(ctx context.Context, conversationList []*model_struct.LocalConversation) error {
	//conversationIDs, err := Exec(ctx)
	return nil
	//if err != nil {
	//	return err
	//} else {
	//	if v, ok := conversationIDs.(string); ok {
	//		var conversationIDs []string
	//		err := utils.JsonStringToStruct(v, &conversationIDs)
	//		if err != nil {
	//			return err
	//		}
	//		var notExistConversations []*model_struct.LocalConversation
	//		var existConversations []*model_struct.LocalConversation
	//		for i, v := range conversationList {
	//			if utils.IsContain(v.ConversationID, conversationIDs) {
	//				existConversations = append(existConversations, v)
	//				continue
	//			} else {
	//				notExistConversations = append(notExistConversations, conversationList[i])
	//			}
	//		}
	//		if len(notExistConversations) > 0 {
	//			err := Exec(ctx, notExistConversations)
	//			if err != nil {
	//				return err
	//			}
	//		}
	//		for _, v := range existConversations {
	//			err := Exec(ctx, v)
	//			if err != nil {
	//				return err
	//			}
	//		}
	//		return nil
	//	} else {
	//		return ErrType
	//	}
	//}
}
