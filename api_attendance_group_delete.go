// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// DeleteAttendanceGroup
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//group_delete
func (r *AttendanceService) DeleteAttendanceGroup(ctx context.Context, request *DeleteAttendanceGroupReq, options ...MethodOptionFunc) (*DeleteAttendanceGroupResp, *Response, error) {
	if r.cli.mock.mockAttendanceDeleteAttendanceGroup != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Attendance#DeleteAttendanceGroup mock enable")
		return r.cli.mock.mockAttendanceDeleteAttendanceGroup(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Attendance",
		API:                   "DeleteAttendanceGroup",
		Method:                "DELETE",
		URL:                   "https://open.feishu.cn/open-apis/attendance/v1/groups/:group_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(deleteAttendanceGroupResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockAttendanceDeleteAttendanceGroup(f func(ctx context.Context, request *DeleteAttendanceGroupReq, options ...MethodOptionFunc) (*DeleteAttendanceGroupResp, *Response, error)) {
	r.mockAttendanceDeleteAttendanceGroup = f
}

func (r *Mock) UnMockAttendanceDeleteAttendanceGroup() {
	r.mockAttendanceDeleteAttendanceGroup = nil
}

type DeleteAttendanceGroupReq struct {
	GroupID string `path:"group_id" json:"-"` // 考勤组的 ID，需要从获取打卡结果的接口中获取 group_id，示例值："6919358128597097404"
}

type deleteAttendanceGroupResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *DeleteAttendanceGroupResp `json:"data,omitempty"`
}

type DeleteAttendanceGroupResp struct{}
