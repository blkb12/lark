// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// DeleteHelpdeskTicketCustomizedField 该接口用于删除工单自定义字段。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket_customized_field/delete
func (r *HelpdeskService) DeleteHelpdeskTicketCustomizedField(ctx context.Context, request *DeleteHelpdeskTicketCustomizedFieldReq, options ...MethodOptionFunc) (*DeleteHelpdeskTicketCustomizedFieldResp, *Response, error) {
	if r.cli.mock.mockHelpdeskDeleteHelpdeskTicketCustomizedField != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#DeleteHelpdeskTicketCustomizedField mock enable")
		return r.cli.mock.mockHelpdeskDeleteHelpdeskTicketCustomizedField(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Helpdesk",
		API:                 "DeleteHelpdeskTicketCustomizedField",
		Method:              "DELETE",
		URL:                 "https://open.feishu.cn/open-apis/helpdesk/v1/ticket_customized_fields/:ticket_customized_field_id",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
		NeedHelpdeskAuth:    true,
	}
	resp := new(deleteHelpdeskTicketCustomizedFieldResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockHelpdeskDeleteHelpdeskTicketCustomizedField(f func(ctx context.Context, request *DeleteHelpdeskTicketCustomizedFieldReq, options ...MethodOptionFunc) (*DeleteHelpdeskTicketCustomizedFieldResp, *Response, error)) {
	r.mockHelpdeskDeleteHelpdeskTicketCustomizedField = f
}

func (r *Mock) UnMockHelpdeskDeleteHelpdeskTicketCustomizedField() {
	r.mockHelpdeskDeleteHelpdeskTicketCustomizedField = nil
}

type DeleteHelpdeskTicketCustomizedFieldReq struct {
	TicketCustomizedFieldID string `path:"ticket_customized_field_id" json:"-"` // 工单自定义字段ID, 示例值："6948728206392295444"
}

type deleteHelpdeskTicketCustomizedFieldResp struct {
	Code int64                                    `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                                   `json:"msg,omitempty"`  // 错误描述
	Data *DeleteHelpdeskTicketCustomizedFieldResp `json:"data,omitempty"`
}

type DeleteHelpdeskTicketCustomizedFieldResp struct{}
