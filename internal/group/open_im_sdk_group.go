package group

import (
	"open_im_sdk/open_im_sdk_callback"
	"open_im_sdk/pkg/common"
	"open_im_sdk/pkg/log"
	"open_im_sdk/pkg/sdk_params_callback"
	"open_im_sdk/pkg/utils"
)

func (g *Group) SetGroupListener(callback open_im_sdk_callback.OnGroupListener) {
	if callback == nil {
		return
	}
	g.listener = callback
}

func (g *Group) CreateGroup(callback open_im_sdk_callback.Base, groupBaseInfo string, memberList string, operationID string) {
	if callback == nil {
		return
	}
	fName := utils.GetSelfFuncName()
	go func() {
		log.NewInfo(operationID, fName, groupBaseInfo, memberList)
		var unmarshalCreateGroupBaseInfoParam sdk_params_callback.CreateGroupBaseInfoParam
		common.JsonUnmarshalAndArgsValidate(groupBaseInfo, &unmarshalCreateGroupBaseInfoParam, callback, operationID)
		var unmarshalCreateGroupMemberRoleParam sdk_params_callback.CreateGroupMemberRoleParam
		common.JsonUnmarshalAndArgsValidate(memberList, &unmarshalCreateGroupMemberRoleParam, callback, operationID)
		result := g.createGroup(callback, unmarshalCreateGroupBaseInfoParam, unmarshalCreateGroupMemberRoleParam, operationID)
		callback.OnSuccess(utils.StructToJsonString(result))
		log.NewInfo(operationID, fName, "callback: ", utils.StructToJsonString(result))
	}()
}

// 已经使用在opencord啦，别删
func (g *Group) GetGroupsInfo(callback open_im_sdk_callback.Base, groupIDList string, operationID string) {
	if callback == nil {
		return
	}
	fName := utils.GetSelfFuncName()
	go func() {
		log.NewInfo(operationID, fName, "args: ", groupIDList)
		var unmarshalGetGroupsInfoParam sdk_params_callback.GetGroupsInfoParam
		common.JsonUnmarshalAndArgsValidate(groupIDList, &unmarshalGetGroupsInfoParam, callback, operationID)
		groupsInfoList := g.getGroupsInfo(unmarshalGetGroupsInfoParam, callback, operationID)
		callback.OnSuccess(utils.StructToJsonStringDefault(groupsInfoList))
		log.NewInfo(operationID, fName, " callback: ", utils.StructToJsonStringDefault(groupsInfoList))

	}()
}

// getGroupMemberListByJoinTimeFilter
func (g *Group) GetGroupMemberListByJoinTimeFilter(callback open_im_sdk_callback.Base, groupID string, offset, count int32, joinTimeBegin, joinTimeEnd int64, filterUserID, operationID string) {
	if callback == nil {
		return
	}
	fName := utils.GetSelfFuncName()
	go func() {
		log.NewInfo(operationID, fName, "args: ", groupID, offset, count, filterUserID)
		var userIDList []string
		common.JsonUnmarshalAndArgsValidate(filterUserID, &userIDList, callback, operationID)
		groupMemberList := g.getGroupMemberListByJoinTimeFilter(callback, groupID, offset, count, joinTimeBegin, joinTimeEnd, userIDList, operationID)
		callback.OnSuccess(utils.StructToJsonStringDefault(groupMemberList))
		log.NewInfo(operationID, fName, " callback: ", utils.StructToJsonStringDefault(groupMemberList))
	}()
}

func (g *Group) GetGroupMembersInfo(callback open_im_sdk_callback.Base, groupID string, userIDList string, operationID string) {
	if callback == nil {
		return
	}
	fName := utils.GetSelfFuncName()
	go func() {
		log.NewInfo(operationID, fName, "args: ", groupID, userIDList)
		var unmarshalParam sdk_params_callback.GetGroupMembersInfoParam
		common.JsonUnmarshalCallback(userIDList, &unmarshalParam, callback, operationID)
		groupMemberList := g.getGroupMembersInfo(callback, groupID, unmarshalParam, operationID)
		callback.OnSuccess(utils.StructToJsonStringDefault(groupMemberList))
		log.NewInfo(operationID, fName, " callback: ", utils.StructToJsonStringDefault(groupMemberList))
	}()
}

func (g *Group) TransferGroupOwner(callback open_im_sdk_callback.Base, groupID, newOwnerUserID string, operationID string) {
	if callback == nil {
		return
	}
	fName := utils.GetSelfFuncName()
	go func() {
		log.NewInfo(operationID, fName, "args: ", groupID, newOwnerUserID)
		g.transferGroupOwner(callback, groupID, newOwnerUserID, operationID)
		callback.OnSuccess(utils.StructToJsonString(sdk_params_callback.TransferGroupOwnerCallback))
		log.NewInfo(operationID, fName, " callback: ", utils.StructToJsonStringDefault(sdk_params_callback.TransferGroupOwnerCallback))
	}()
}

func (g *Group) InviteUserToGroup(callback open_im_sdk_callback.Base, groupID, reason string, userIDList string, operationID string) {
	if callback == nil {
		return
	}
	fName := utils.GetSelfFuncName()
	go func() {
		log.NewInfo(operationID, fName, "args: ", groupID, reason, userIDList)
		var unmarshalParam sdk_params_callback.InviteUserToGroupParam
		common.JsonUnmarshalAndArgsValidate(userIDList, &unmarshalParam, callback, operationID)
		result := g.inviteUserToGroup(callback, groupID, reason, unmarshalParam, operationID)
		callback.OnSuccess(utils.StructToJsonStringDefault(result))
		log.NewInfo(operationID, fName, "callback: ", utils.StructToJsonStringDefault(result))
	}()
}

func (g *Group) GetRecvGroupApplicationList(callback open_im_sdk_callback.Base, operationID string) {
	if callback == nil {
		return
	}
	fName := utils.GetSelfFuncName()
	go func() {
		log.NewInfo(operationID, fName, "args: ")
		result := g.getRecvGroupApplicationList(callback, operationID)
		callback.OnSuccess(utils.StructToJsonStringDefault(result))
		log.NewInfo(operationID, fName, "callback: ", utils.StructToJsonString(result))
	}()
}

func (g *Group) GetSendGroupApplicationList(callback open_im_sdk_callback.Base, operationID string) {
	if callback == nil {
		return
	}
	fName := utils.GetSelfFuncName()
	go func() {
		log.NewInfo(operationID, "output results")
		result := g.getSendGroupApplicationList(callback, operationID)
		callback.OnSuccess(utils.StructToJsonStringDefault(result))
		log.NewInfo(operationID, fName, "callback: ", utils.StructToJsonString(result))
	}()
}

func (g *Group) AcceptGroupApplication(callback open_im_sdk_callback.Base, groupID, fromUserID, handleMsg string, operationID string) {
	if callback == nil {
		return
	}
	fName := utils.GetSelfFuncName()
	go func() {
		log.NewInfo(operationID, fName, "args: ", groupID, fromUserID, handleMsg)
		g.processGroupApplication(callback, groupID, fromUserID, handleMsg, 1, operationID)
		callback.OnSuccess(utils.StructToJsonString(sdk_params_callback.AcceptGroupApplicationCallback))
		log.NewInfo(operationID, fName, "callback: ", utils.StructToJsonString(sdk_params_callback.AcceptGroupApplicationCallback))
	}()
}

func (g *Group) RefuseGroupApplication(callback open_im_sdk_callback.Base, groupID, fromUserID, handleMsg string, operationID string) {
	if callback == nil {
		return
	}
	fName := utils.GetSelfFuncName()
	go func() {
		log.NewInfo(operationID, fName, "args: ", groupID, fromUserID, handleMsg)
		g.processGroupApplication(callback, groupID, fromUserID, handleMsg, -1, operationID)
		callback.OnSuccess(utils.StructToJsonString(sdk_params_callback.RefuseGroupApplicationCallback))
		log.NewInfo(operationID, fName, "callback: ", utils.StructToJsonString(sdk_params_callback.RefuseGroupApplicationCallback))
	}()
}

func (g *Group) SetGroupMemberNickname(callback open_im_sdk_callback.Base, groupID, userID string, GroupMemberNickname string, operationID string) {
	if callback == nil {
		return
	}
	fName := utils.GetSelfFuncName()
	go func() {
		log.NewInfo(operationID, fName, "args: ", groupID, userID, GroupMemberNickname)
		g.setGroupMemberNickname(callback, groupID, userID, GroupMemberNickname, operationID)
		callback.OnSuccess(utils.StructToJsonString(sdk_params_callback.SetGroupMemberNicknameCallback))
		log.NewInfo(operationID, fName, "callback: ", utils.StructToJsonString(sdk_params_callback.SetGroupMemberNicknameCallback))
	}()
}

func (g *Group) SearchGroupMembers(callback open_im_sdk_callback.Base, searchParam string, operationID string) {
	if callback == nil {
		return
	}
	fName := utils.GetSelfFuncName()
	go func() {
		log.NewInfo(operationID, fName, "args: ", searchParam)
		var unmarshalSearchGroupMembersParam sdk_params_callback.SearchGroupMembersParam
		common.JsonUnmarshalAndArgsValidate(searchParam, &unmarshalSearchGroupMembersParam, callback, operationID)
		members := g.searchGroupMembers(callback, unmarshalSearchGroupMembersParam, operationID)
		callback.OnSuccess(utils.StructToJsonStringDefault(members))
		log.NewInfo(operationID, fName, "callback: ", utils.StructToJsonStringDefault(members))
	}()
}
