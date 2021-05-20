// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetAttendanceUserApproval
//
// 获取员工在某段时间内的请假、加班、外出和出差四种审批的通过数据。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//RetrieveUserApprovals
func (r *AttendanceService) GetAttendanceUserApproval(ctx context.Context, request *GetAttendanceUserApprovalReq, options ...MethodOptionFunc) (*GetAttendanceUserApprovalResp, *Response, error) {
	if r.cli.mock.mockAttendanceGetAttendanceUserApproval != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Attendance#GetAttendanceUserApproval mock enable")
		return r.cli.mock.mockAttendanceGetAttendanceUserApproval(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Attendance",
		API:                   "GetAttendanceUserApproval",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/attendance/v1/user_approvals/query",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getAttendanceUserApprovalResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockAttendanceGetAttendanceUserApproval(f func(ctx context.Context, request *GetAttendanceUserApprovalReq, options ...MethodOptionFunc) (*GetAttendanceUserApprovalResp, *Response, error)) {
	r.mockAttendanceGetAttendanceUserApproval = f
}

func (r *Mock) UnMockAttendanceGetAttendanceUserApproval() {
	r.mockAttendanceGetAttendanceUserApproval = nil
}

type GetAttendanceUserApprovalReq struct {
	EmployeeType  EmployeeType `query:"employee_type" json:"-"`   // 请求体中的 user_ids 的员工工号类型，必选字段，可用值：【employee_id（员工的 employeeId），employee_no（员工工号）】，示例值："employee_id"
	UserIDs       []string     `json:"user_ids,omitempty"`        // employee_no 或 employee_id 列表
	CheckDateFrom int64        `json:"check_date_from,omitempty"` // 查询的起始工作日
	CheckDateTo   int64        `json:"check_date_to,omitempty"`   // 查询的结束工作日
}

type getAttendanceUserApprovalResp struct {
	Code int64                          `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                         `json:"msg,omitempty"`  // 错误描述
	Data *GetAttendanceUserApprovalResp `json:"data,omitempty"` // -
}

type GetAttendanceUserApprovalResp struct {
	UserApprovals []*GetAttendanceUserApprovalRespUserApproval `json:"user_approvals,omitempty"` // 审批结果列表
}

type GetAttendanceUserApprovalRespUserApproval struct {
	UserID        string                                                   `json:"user_id,omitempty"`        // 审批用户 ID
	Date          string                                                   `json:"date,omitempty"`           // 审批作用时间
	Outs          []*GetAttendanceUserApprovalRespUserApprovalOut          `json:"outs,omitempty"`           // 外出信息
	Leaves        []*GetAttendanceUserApprovalRespUserApprovalLeave        `json:"leaves,omitempty"`         // 请假信息
	OvertimeWorks []*GetAttendanceUserApprovalRespUserApprovalOvertimeWork `json:"overtime_works,omitempty"` // 加班信息
	Trips         []*GetAttendanceUserApprovalRespUserApprovalTrip         `json:"trips,omitempty"`          // 出差信息
}

type GetAttendanceUserApprovalRespUserApprovalOut struct {
	UniqID           string     `json:"uniq_id,omitempty"`            // 外出类型唯一 ID，代表一种外出类型，长度小于 14
	Unit             int64      `json:"unit,omitempty"`               // 外出时长单位，可用值：【1（天），2（小时），3（半天），4（半小时）】
	Interval         int64      `json:"interval,omitempty"`           // 假期时长（单位秒）
	StartTime        string     `json:"start_time,omitempty"`         // 开始时间，时间格式为 yyyy-MM-dd HH:mm:ss
	EndTime          string     `json:"end_time,omitempty"`           // 结束时间，时间格式为 yyyy-MM-dd HH:mm:ss
	I18nNames        *I18nNames `json:"i18n_names,omitempty"`         // 外出多语言展示，格式为 map，key 为["ch"、"en"、"ja"]，其中 ch 代表中文，en 代表英文、ja 代表日文
	DefaultLocale    string     `json:"default_locale,omitempty"`     // 默认语言类型，由于飞书客户端支持中、英、日三种语言，当用户切换语言时，如果外出名称没有所对应语言的名称，则会使用默认语言的名称
	Reason           string     `json:"reason,omitempty"`             // 外出理由
	ApprovePassTime  string     `json:"approve_pass_time,omitempty"`  // 审批通过时间
	ApproveApplyTime string     `json:"approve_apply_time,omitempty"` // 审批申请时间
}

type GetAttendanceUserApprovalRespUserApprovalLeave struct {
	UniqID           string     `json:"uniq_id,omitempty"`            // 假期类型唯一 ID，代表一种假期类型，长度小于 14
	Unit             int64      `json:"unit,omitempty"`               // 假期时长单位，可用值：【1（天），2（小时），3（半天），4（半小时）】
	Interval         int64      `json:"interval,omitempty"`           // 假期时长（单位秒）
	StartTime        string     `json:"start_time,omitempty"`         // 开始时间，时间格式为 yyyy-MM-dd HH:mm:ss
	EndTime          string     `json:"end_time,omitempty"`           // 结束时间，时间格式为 yyyy-MM-dd HH:mm:ss
	I18nNames        *I18nNames `json:"i18n_names,omitempty"`         // 假期多语言展示，格式为 map，key 为["ch"、"en"、"ja"]，其中 ch 代表中文，en 代表英文、ja 代表日文
	DefaultLocale    string     `json:"default_locale,omitempty"`     // 默认语言类型，由于飞书客户端支持中、英、日三种语言，当用户切换语言时，如果假期名称没有所对应语言的名称，则会使用默认语言的名称，可用值：【ch（中文），en（英文），ja（日文）】
	Reason           string     `json:"reason,omitempty"`             // 请假理由
	ApprovePassTime  string     `json:"approve_pass_time,omitempty"`  // 审批通过时间，时间格式为 yyyy-MM-dd HH:mm:ss
	ApproveApplyTime string     `json:"approve_apply_time,omitempty"` // 审批申请时间，时间格式为 yyyy-MM-dd HH:mm:ss
}

type GetAttendanceUserApprovalRespUserApprovalOvertimeWork struct {
	Duration  float64 `json:"duration,omitempty"`   // 加班时长
	Unit      int64   `json:"unit,omitempty"`       // 加班时长单位，可用值：【1（天），2（小时）】
	Category  int64   `json:"category,omitempty"`   // 加班日期类型，可用值：【1（工作日），2（休息日），3（节假日）】
	Type      int64   `json:"type,omitempty"`       // 加班规则类型，可用值：【0（不关联加班规则），1（调休），2（加班费）】
	StartTime string  `json:"start_time,omitempty"` // 开始时间，时间格式为 yyyy-MM-dd HH:mm:ss
	EndTime   string  `json:"end_time,omitempty"`   // 结束时间，时间格式为 yyyy-MM-dd HH:mm:ss
}

type GetAttendanceUserApprovalRespUserApprovalTrip struct {
	StartTime        string `json:"start_time,omitempty"`         // 开始时间，时间格式为 yyyy-MM-dd HH:mm:ss
	EndTime          string `json:"end_time,omitempty"`           // 结束时间，时间格式为 yyyy-MM-dd HH:mm:ss
	Reason           string `json:"reason,omitempty"`             // 出差理由
	ApprovePassTime  string `json:"approve_pass_time,omitempty"`  // 审批通过时间，时间格式为 yyyy-MM-dd HH:mm:ss
	ApproveApplyTime string `json:"approve_apply_time,omitempty"` // 审批申请时间，时间格式为 yyyy-MM-dd HH:mm:ss
}
