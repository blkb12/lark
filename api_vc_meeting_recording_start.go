// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// StartVCMeetingRecording 在会议中开始录制。
//
// 会议正在进行中，且操作者具有相应权限（如果操作者为用户，必须是会中当前主持人）
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting-recording/start
func (r *VCService) StartVCMeetingRecording(ctx context.Context, request *StartVCMeetingRecordingReq, options ...MethodOptionFunc) (*StartVCMeetingRecordingResp, *Response, error) {
	if r.cli.mock.mockVCStartVCMeetingRecording != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] VC#StartVCMeetingRecording mock enable")
		return r.cli.mock.mockVCStartVCMeetingRecording(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "VC",
		API:                 "StartVCMeetingRecording",
		Method:              "PATCH",
		URL:                 "https://open.feishu.cn/open-apis/vc/v1/meetings/:meeting_id/recording/start",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
	}
	resp := new(startVCMeetingRecordingResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockVCStartVCMeetingRecording(f func(ctx context.Context, request *StartVCMeetingRecordingReq, options ...MethodOptionFunc) (*StartVCMeetingRecordingResp, *Response, error)) {
	r.mockVCStartVCMeetingRecording = f
}

func (r *Mock) UnMockVCStartVCMeetingRecording() {
	r.mockVCStartVCMeetingRecording = nil
}

type StartVCMeetingRecordingReq struct {
	MeetingID string `path:"meeting_id" json:"-"` // 会议ID, 示例值: "6911188411932033028"
	Timezone  *int64 `json:"timezone,omitempty"`  // 录制文件时间显示使用的时区[-12,12], 示例值: 8
}

type startVCMeetingRecordingResp struct {
	Code int64                        `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                       `json:"msg,omitempty"`  // 错误描述
	Data *StartVCMeetingRecordingResp `json:"data,omitempty"`
}

type StartVCMeetingRecordingResp struct{}
