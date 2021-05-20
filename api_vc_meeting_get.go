// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetVCMeeting 获取一个会议的详细数据
//
// 只能获取归属于自己（或参与）的会议，支持查询最近90天内的会议
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/get
func (r *VCService) GetVCMeeting(ctx context.Context, request *GetVCMeetingReq, options ...MethodOptionFunc) (*GetVCMeetingResp, *Response, error) {
	if r.cli.mock.mockVCGetVCMeeting != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] VC#GetVCMeeting mock enable")
		return r.cli.mock.mockVCGetVCMeeting(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "VC",
		API:                   "GetVCMeeting",
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/vc/v1/meetings/:meeting_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getVCMeetingResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockVCGetVCMeeting(f func(ctx context.Context, request *GetVCMeetingReq, options ...MethodOptionFunc) (*GetVCMeetingResp, *Response, error)) {
	r.mockVCGetVCMeeting = f
}

func (r *Mock) UnMockVCGetVCMeeting() {
	r.mockVCGetVCMeeting = nil
}

type GetVCMeetingReq struct {
	WithParticipants   *bool   `query:"with_participants" json:"-"`    // 是否需要参会人列表, 示例值：false
	WithMeetingAbility *bool   `query:"with_meeting_ability" json:"-"` // 是否需要会中使用能力统计（仅限tenant_access_token）, 示例值：false
	UserIDType         *IDType `query:"user_id_type" json:"-"`         // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 userid
	MeetingID          string  `path:"meeting_id" json:"-"`            // 会议ID, 示例值："6911188411932033028"
}

type getVCMeetingResp struct {
	Code int64             `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string            `json:"msg,omitempty"`  // 错误描述
	Data *GetVCMeetingResp `json:"data,omitempty"` //
}

type GetVCMeetingResp struct {
	Meeting *GetVCMeetingRespMeeting `json:"meeting,omitempty"` // 会议数据
}

type GetVCMeetingRespMeeting struct {
	ID               string                                `json:"id,omitempty"`                // 会议ID
	Topic            string                                `json:"topic,omitempty"`             // 会议主题
	URL              string                                `json:"url,omitempty"`               // 会议链接
	CreateTime       string                                `json:"create_time,omitempty"`       // 会议创建时间（unix时间，单位sec）
	StartTime        string                                `json:"start_time,omitempty"`        // 会议开始时间（unix时间，单位sec）
	EndTime          string                                `json:"end_time,omitempty"`          // 会议结束时间（unix时间，单位sec）
	HostUser         *GetVCMeetingRespMeetingHostUser      `json:"host_user,omitempty"`         // 主持人
	Status           int64                                 `json:"status,omitempty"`            // 会议状态, 可选值有: `1`：会议呼叫中, `2`：会议进行中, `3`：会议已结束
	ParticipantCount string                                `json:"participant_count,omitempty"` // 参会人数
	Participants     []*GetVCMeetingRespMeetingParticipant `json:"participants,omitempty"`      // 参会人列表
	Ability          *GetVCMeetingRespMeetingAbility       `json:"ability,omitempty"`           // 会中使用的能力
}

type GetVCMeetingRespMeetingHostUser struct {
	ID       string `json:"id,omitempty"`        // 用户ID
	UserType int64  `json:"user_type,omitempty"` // 用户类型, 可选值有: `1`：lark用户, `2`：rooms用户, `3`：文档用户, `4`：neo单品用户, `5`：neo单品游客用户, `6`：pstn用户, `7`：sip用户
}

type GetVCMeetingRespMeetingParticipant struct {
	ID         string `json:"id,omitempty"`          // 用户ID
	UserType   int64  `json:"user_type,omitempty"`   // 用户类型, 可选值有: `1`：lark用户, `2`：rooms用户, `3`：文档用户, `4`：neo单品用户, `5`：neo单品游客用户, `6`：pstn用户, `7`：sip用户
	IsHost     bool   `json:"is_host,omitempty"`     // 是否为主持人
	IsCohost   bool   `json:"is_cohost,omitempty"`   // 是否为联席主持人
	IsExternal bool   `json:"is_external,omitempty"` // 是否为外部参会人
	Status     int64  `json:"status,omitempty"`      // 参会人状态, 可选值有: `1`：呼叫中, `2`：在会中, `3`：正在响铃, `4`：不在会中或已经离开会议
}

type GetVCMeetingRespMeetingAbility struct {
	UseVideo        bool `json:"use_video,omitempty"`         // 是否使用视频
	UseAudio        bool `json:"use_audio,omitempty"`         // 是否使用音频
	UseShareScreen  bool `json:"use_share_screen,omitempty"`  // 是否使用共享屏幕
	UseFollowScreen bool `json:"use_follow_screen,omitempty"` // 是否使用妙享（magic share）
	UseRecording    bool `json:"use_recording,omitempty"`     // 是否使用录制
	UsePstn         bool `json:"use_pstn,omitempty"`          // 是否使用PSTN
}
