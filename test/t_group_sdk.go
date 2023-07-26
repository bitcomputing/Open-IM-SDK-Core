package test

import (
	//	"encoding/json"
	"fmt"
	"open_im_sdk/pkg/log"
	"open_im_sdk/pkg/utils"
	//"open_im_sdk/internal/open_im_sdk"
	//"open_im_sdk/pkg/utils"
	//	"open_im_sdk/internal/common"
)

type XBase struct {
}

func (XBase) OnError(errCode int32, errMsg string) {
	fmt.Println("get groupmenberinfo OnError", errCode, errMsg)
}
func (XBase) OnSuccess(data string) {
	fmt.Println("get groupmenberinfo OnSuccess, ", data)
}

func (XBase) OnProgress(progress int) {
	fmt.Println("OnProgress, ", progress)
}

type testGroupListener struct {
}

func (testGroupListener) OnJoinedGroupAdded(callbackInfo string) {
	log.Info(utils.OperationIDGenerator(), utils.GetSelfFuncName(), callbackInfo)

}
func (testGroupListener) OnJoinedGroupDeleted(callbackInfo string) {
	log.Info(utils.OperationIDGenerator(), utils.GetSelfFuncName(), callbackInfo)

}

func (testGroupListener) OnGroupMemberAdded(callbackInfo string) {
	log.Info(utils.OperationIDGenerator(), utils.GetSelfFuncName(), callbackInfo)

}
func (testGroupListener) OnGroupMemberDeleted(callbackInfo string) {
	log.Info(utils.OperationIDGenerator(), utils.GetSelfFuncName(), callbackInfo)

}

func (testGroupListener) OnGroupApplicationAdded(callbackInfo string) {
	log.Info(utils.OperationIDGenerator(), utils.GetSelfFuncName(), callbackInfo)

}
func (testGroupListener) OnGroupApplicationDeleted(callbackInfo string) {
	log.Info(utils.OperationIDGenerator(), utils.GetSelfFuncName(), callbackInfo)

}

func (testGroupListener) OnGroupInfoChanged(callbackInfo string) {
	log.Info(utils.OperationIDGenerator(), utils.GetSelfFuncName(), callbackInfo)

}
func (testGroupListener) OnGroupMemberInfoChanged(callbackInfo string) {
	log.Info(utils.OperationIDGenerator(), utils.GetSelfFuncName(), callbackInfo)

}

func (testGroupListener) OnGroupApplicationAccepted(callbackInfo string) {
	log.Info(utils.OperationIDGenerator(), utils.GetSelfFuncName(), callbackInfo)

}
func (testGroupListener) OnGroupApplicationRejected(callbackInfo string) {
	log.Info(utils.OperationIDGenerator(), utils.GetSelfFuncName(), callbackInfo)

}

type testCreateGroup struct {
	OperationID string
}

func (t testCreateGroup) OnSuccess(data string) {
	log.Info(t.OperationID, utils.GetSelfFuncName(), data)

}

func (t testCreateGroup) OnError(errCode int32, errMsg string) {
	log.Info(t.OperationID, utils.GetSelfFuncName(), errCode, errMsg)
}

type testSetGroupInfo struct {
	OperationID string
}

func (t testSetGroupInfo) OnSuccess(data string) {
	log.Info(t.OperationID, utils.GetSelfFuncName(), data)

}

func (t testSetGroupInfo) OnError(errCode int32, errMsg string) {
	log.Info(t.OperationID, utils.GetSelfFuncName(), errCode, errMsg)
}

type testGetGroupsInfo struct {
	OperationID string
}

func (t testGetGroupsInfo) OnSuccess(data string) {
	log.Info(t.OperationID, "testGetGroupsInfo,onSuccess", data)
}

func (t testGetGroupsInfo) OnError(errCode int32, errMsg string) {
	log.Info(t.OperationID, "testGetGroupsInfo,onError", errCode, errMsg)
}

type testSearchGroups struct {
	OperationID string
}

func (t testSearchGroups) OnSuccess(data string) {
	log.Info(t.OperationID, "testSearchGroups,onSuccess", data)
}

func (t testSearchGroups) OnError(errCode int32, errMsg string) {
	log.Info(t.OperationID, "testSearchGroups,onError", errCode, errMsg)
}

type testJoinGroup struct {
	OperationID string
}

func (t testJoinGroup) OnSuccess(data string) {
	log.Info(t.OperationID, "testJoinGroup,onSuccess", data)
}

func (t testJoinGroup) OnError(errCode int32, errMsg string) {
	log.Info(t.OperationID, "testJoinGroup,onError", errCode, errMsg)
}

type testQuitGroup struct {
	OperationID string
}

func (t testQuitGroup) OnSuccess(data string) {
	log.Info(t.OperationID, "testQuitGroup,onSuccess", data)
}

func (t testQuitGroup) OnError(errCode int32, errMsg string) {
	log.Info(t.OperationID, "testQuitGroup,onError", errCode, errMsg)
}

type testGetJoinedGroupList struct {
	OperationID string
}

/*
OnError(errCode int, errMsg string)
OnSuccess(data string)
*/
func (t testGetJoinedGroupList) OnError(errCode int32, errMsg string) {
	log.Info(t.OperationID, "testGetJoinedGroupList OnError", errCode, errMsg)
}

func (t testGetJoinedGroupList) OnSuccess(data string) {
	log.Info(t.OperationID, "testGetJoinedGroupList OnSuccess, output", data)
}

type testGetGroupMemberList struct {
	OperationID string
}

func (t testGetGroupMemberList) OnSuccess(data string) {
	log.Info(t.OperationID, utils.GetSelfFuncName(), "testGetGroupMemberList: ", data)

}

func (t testGetGroupMemberList) OnError(errCode int32, errMsg string) {
	log.Info(t.OperationID, utils.GetSelfFuncName(), "testGetGroupMemberList", errCode, errMsg)
}

func DotestCos() {
	//var callback baseCallback
	//p := ws.NewPostApi(token, userForSDK.ImConfig().ApiAddr)
	//var storage common.ObjectStorage = common.NewCOS(p)
	//test(storage, callback)
}

//func DotestMinio() {
//	var callback baseCallback
//	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVSUQiOiIxMzkwMDAwMDAwMCIsIlBsYXRmb3JtIjoiSU9TIiwiZXhwIjoxNjQ1NzgyNDY0LCJuYmYiOjE2NDUxNzc2NjQsImlhdCI6MTY0NTE3NzY2NH0.T-SDoLxdlwRGOMZPIKriPtAlOGWCLodsGi1dWxN8kto"
//	p := ws.NewPostApi(token, "https://storage.rentsoft.cn")
//	minio := common.NewMinio(p)
//	var storage common.ObjectStorage = minio
//	log.NewInfo("", *minio)
//	test(storage, callback)
//}
//

type testGetGroupMembersInfo struct {
}

func (testGetGroupMembersInfo) OnError(errCode int32, errMsg string) {
	fmt.Println("testGetGroupMembersInfo OnError", errCode, errMsg)
}

func (testGetGroupMembersInfo) OnSuccess(data string) {
	fmt.Println("testGetGroupMembersInfo OnSuccess, output", data)
}

//
//func DotestGetGroupMembersInfo() {
//	var test testGetGroupMembersInfo
//	var memlist []string
//	memlist = append(memlist, "307edc814bb0d04a")
//	//memlist = append(memlist, "ded01dfe543700402608e19d4e2f839e")
//	jlist, _ := json.Marshal(memlist)
//	fmt.Println("GetGroupMembersInfo input : ", string(jlist))
//	sdk_interface.GetGroupMembersInfo("7ff61d8f9d4a8a0d6d70a14e2683aad5", string(jlist), test)
//	//GetGroupMemberList("05dc84b52829e82242a710ecf999c72c", 0, 0, test)
//}
//

type baseCallback struct {
	OperationID string
	callName    string
}

func (t baseCallback) OnSuccess(data string) {
	log.Info(t.OperationID, t.callName, utils.GetSelfFuncName(), data)

}

func (t baseCallback) OnError(errCode int32, errMsg string) {
	log.Info(t.OperationID, t.callName, utils.GetSelfFuncName(), errCode, errMsg)
}

type testProcessGroupApplication struct {
	baseCallback
}
