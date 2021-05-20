// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// DeleteAttendanceShift
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//shift_delete
func (r *AttendanceService) DeleteAttendanceShift(ctx context.Context, request *DeleteAttendanceShiftReq, options ...MethodOptionFunc) (*DeleteAttendanceShiftResp, *Response, error) {
	if r.cli.mock.mockAttendanceDeleteAttendanceShift != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Attendance#DeleteAttendanceShift mock enable")
		return r.cli.mock.mockAttendanceDeleteAttendanceShift(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Attendance",
		API:                   "DeleteAttendanceShift",
		Method:                "DELETE",
		URL:                   "https://open.feishu.cn/open-apis/attendance/v1/shifts/:shift_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(deleteAttendanceShiftResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockAttendanceDeleteAttendanceShift(f func(ctx context.Context, request *DeleteAttendanceShiftReq, options ...MethodOptionFunc) (*DeleteAttendanceShiftResp, *Response, error)) {
	r.mockAttendanceDeleteAttendanceShift = f
}

func (r *Mock) UnMockAttendanceDeleteAttendanceShift() {
	r.mockAttendanceDeleteAttendanceShift = nil
}

type DeleteAttendanceShiftReq struct {
	ShiftID string `path:"shift_id" json:"-"` // 班次 ID，示例值："6919358778597097404"
}

type deleteAttendanceShiftResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *DeleteAttendanceShiftResp `json:"data,omitempty"`
}

type DeleteAttendanceShiftResp struct{}
