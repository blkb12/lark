// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// BatchGetMeetingRoomBuildingID 该接口用于删除建筑物（办公大楼）。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzMxYjLzMTM24yMzEjN
func (r *MeetingRoomService) BatchGetMeetingRoomBuildingID(ctx context.Context, request *BatchGetMeetingRoomBuildingIDReq, options ...MethodOptionFunc) (*BatchGetMeetingRoomBuildingIDResp, *Response, error) {
	if r.cli.mock.mockMeetingRoomBatchGetMeetingRoomBuildingID != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] MeetingRoom#BatchGetMeetingRoomBuildingID mock enable")
		return r.cli.mock.mockMeetingRoomBatchGetMeetingRoomBuildingID(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "MeetingRoom",
		API:                   "BatchGetMeetingRoomBuildingID",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/meeting_room/building/delete",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(batchGetMeetingRoomBuildingIDResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockMeetingRoomBatchGetMeetingRoomBuildingID(f func(ctx context.Context, request *BatchGetMeetingRoomBuildingIDReq, options ...MethodOptionFunc) (*BatchGetMeetingRoomBuildingIDResp, *Response, error)) {
	r.mockMeetingRoomBatchGetMeetingRoomBuildingID = f
}

func (r *Mock) UnMockMeetingRoomBatchGetMeetingRoomBuildingID() {
	r.mockMeetingRoomBatchGetMeetingRoomBuildingID = nil
}

type BatchGetMeetingRoomBuildingIDReq struct {
	BuildingID string `json:"building_id,omitempty"` // 要删除的建筑ID
}

type batchGetMeetingRoomBuildingIDResp struct {
	Code int64                              `json:"code,omitempty"` // 返回码，非 0 表示失败
	Msg  string                             `json:"msg,omitempty"`  // 返回码的描述，"success" 表示成功，其他为错误提示信息
	Data *BatchGetMeetingRoomBuildingIDResp `json:"data,omitempty"`
}

type BatchGetMeetingRoomBuildingIDResp struct{}
