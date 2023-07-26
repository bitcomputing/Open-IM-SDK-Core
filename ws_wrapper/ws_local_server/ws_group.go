package ws_local_server

import (
	"encoding/json"
	"open_im_sdk/open_im_sdk"
	"open_im_sdk/pkg/log"
	"open_im_sdk/pkg/utils"
)

type GroupCallback struct {
	uid string
}

func (g *GroupCallback) OnJoinedGroupAdded(groupInfo string) {
	SendOneUserMessage(EventData{cleanUpfuncName(runFuncName()), 0, "", groupInfo, "0"}, g.uid)
}
func (g *GroupCallback) OnJoinedGroupDeleted(groupInfo string) {
	SendOneUserMessage(EventData{cleanUpfuncName(runFuncName()), 0, "", groupInfo, "0"}, g.uid)
}

func (g *GroupCallback) OnGroupMemberAdded(groupMemberInfo string) {
	SendOneUserMessage(EventData{cleanUpfuncName(runFuncName()), 0, "", groupMemberInfo, "0"}, g.uid)
}
func (g *GroupCallback) OnGroupMemberDeleted(groupMemberInfo string) {
	SendOneUserMessage(EventData{cleanUpfuncName(runFuncName()), 0, "", groupMemberInfo, "0"}, g.uid)
}

func (g *GroupCallback) OnGroupApplicationAdded(groupApplication string) {
	SendOneUserMessage(EventData{cleanUpfuncName(runFuncName()), 0, "", groupApplication, "0"}, g.uid)
}
func (g *GroupCallback) OnGroupApplicationDeleted(groupApplication string) {
	SendOneUserMessage(EventData{cleanUpfuncName(runFuncName()), 0, "", groupApplication, "0"}, g.uid)
}

func (g *GroupCallback) OnGroupInfoChanged(groupInfo string) {
	SendOneUserMessage(EventData{cleanUpfuncName(runFuncName()), 0, "", groupInfo, "0"}, g.uid)
}
func (g *GroupCallback) OnGroupMemberInfoChanged(groupMemberInfo string) {
	SendOneUserMessage(EventData{cleanUpfuncName(runFuncName()), 0, "", groupMemberInfo, "0"}, g.uid)
}

func (g *GroupCallback) OnGroupApplicationAccepted(groupApplication string) {
	SendOneUserMessage(EventData{cleanUpfuncName(runFuncName()), 0, "", groupApplication, "0"}, g.uid)
}
func (g *GroupCallback) OnGroupApplicationRejected(groupApplication string) {
	SendOneUserMessage(EventData{cleanUpfuncName(runFuncName()), 0, "", groupApplication, "0"}, g.uid)
}

func (wsRouter *WsFuncRouter) SetGroupListener() {
	var g GroupCallback
	g.uid = wsRouter.uId
	userWorker := open_im_sdk.GetUserWorker(wsRouter.uId)
	userWorker.SetGroupListener(&g)
}

func (wsRouter *WsFuncRouter) CreateGroup(input, operationID string) {
	m := make(map[string]interface{})
	if err := json.Unmarshal([]byte(input), &m); err != nil {
		log.Info(operationID, utils.GetSelfFuncName(), "unmarshal failed", input, err.Error())
		wsRouter.GlobalSendMessage(EventData{cleanUpfuncName(runFuncName()), StatusBadParameter, "unmarshal failed", "", operationID})
		return
	}

	userWorker := open_im_sdk.GetUserWorker(wsRouter.uId)
	if !wsRouter.checkResourceLoadingAndKeysIn(userWorker, input, operationID, runFuncName(), m, "groupBaseInfo", "memberList") {
		return
	}
	//callback common.Base, groupBaseInfo string, memberList string, operationID string
	userWorker.Group().CreateGroup(&BaseSuccessFailed{runFuncName(), operationID, wsRouter.uId},
		m["groupBaseInfo"].(string), m["memberList"].(string), operationID)
}

func (wsRouter *WsFuncRouter) GetGroupsInfo(input, operationID string) { //(groupIdList string, callback Base) {
	//m := make(map[string]interface{})
	//if err := json.Unmarshal([]byte(input), &m); err != nil {
	//	log.Info("unmarshal failed")
	//	wsRouter.GlobalSendMessage(EventData{cleanUpfuncName(runFuncName()), StatusBadParameter, "unmarshal failed", "", operationID})
	//	return
	//}
	//if !wsRouter.checkKeysIn(input, operationID, runFuncName(), m, "groupIDList") {
	//	return
	//}

	userWorker := open_im_sdk.GetUserWorker(wsRouter.uId)
	if !wsRouter.checkResourceLoadingAndKeysIn(userWorker, input, operationID, runFuncName(), nil) {
		return
	}
	//callback common.Base, groupIDList string, operationID string
	userWorker.Group().GetGroupsInfo(&BaseSuccessFailed{runFuncName(), operationID, wsRouter.uId},
		input, operationID)
}
