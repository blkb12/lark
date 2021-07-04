// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetCalendarEvent 该接口用于以当前身份（应用 / 用户）获取日历上的一个日程。
//
// 当前身份必须对日历有访问权限。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/get
func (r *CalendarService) GetCalendarEvent(ctx context.Context, request *GetCalendarEventReq, options ...MethodOptionFunc) (*GetCalendarEventResp, *Response, error) {
	if r.cli.mock.mockCalendarGetCalendarEvent != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Calendar#GetCalendarEvent mock enable")
		return r.cli.mock.mockCalendarGetCalendarEvent(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Calendar",
		API:                   "GetCalendarEvent",
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getCalendarEventResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockCalendarGetCalendarEvent(f func(ctx context.Context, request *GetCalendarEventReq, options ...MethodOptionFunc) (*GetCalendarEventResp, *Response, error)) {
	r.mockCalendarGetCalendarEvent = f
}

func (r *Mock) UnMockCalendarGetCalendarEvent() {
	r.mockCalendarGetCalendarEvent = nil
}

type GetCalendarEventReq struct {
	CalendarID string `path:"calendar_id" json:"-"` // 日历ID, 示例值："feishu.cn_xxxxxxxxxx@group.calendar.feishu.cn"
	EventID    string `path:"event_id" json:"-"`    // 日程ID, 示例值："xxxxxxxxx_0"
}

type getCalendarEventResp struct {
	Code int64                 `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                `json:"msg,omitempty"`  // 错误描述
	Data *GetCalendarEventResp `json:"data,omitempty"`
}

type GetCalendarEventResp struct {
	Event *GetCalendarEventRespEvent `json:"event,omitempty"` // 日程实体
}

type GetCalendarEventRespEvent struct {
	EventID          string                               `json:"event_id,omitempty"`           // 日程ID
	Summary          string                               `json:"summary,omitempty"`            // 日程标题, 最大长度：`1000` 字符
	Description      string                               `json:"description,omitempty"`        // 日程描述, 最大长度：`8192` 字符
	StartTime        *GetCalendarEventRespEventStartTime  `json:"start_time,omitempty"`         // 日程开始时间
	EndTime          *GetCalendarEventRespEventEndTime    `json:"end_time,omitempty"`           // 日程结束时间
	Vchat            *GetCalendarEventRespEventVchat      `json:"vchat,omitempty"`              // 视频会议信息，仅当日程至少有一位attendee时生效
	Visibility       string                               `json:"visibility,omitempty"`         // 日程公开范围，新建日程默认为Default；仅新建日程时对所有参与人生效，之后修改该属性仅对当前身份生效, 可选值有: `default`：默认权限，跟随日历权限，默认仅向他人显示是否“忙碌”, `public`：公开，显示日程详情, `private`：私密，仅自己可见详情
	AttendeeAbility  string                               `json:"attendee_ability,omitempty"`   // 参与人权限, 可选值有: `none`：无法编辑日程、无法邀请其它参与人、无法查看参与人列表, `can_see_others`：无法编辑日程、无法邀请其它参与人、可以查看参与人列表, `can_invite_others`：无法编辑日程、可以邀请其它参与人、可以查看参与人列表, `can_modify_event`：可以编辑日程、可以邀请其它参与人、可以查看参与人列表
	FreeBusyStatus   string                               `json:"free_busy_status,omitempty"`   // 日程占用的忙闲状态，新建日程默认为Busy；仅新建日程时对所有参与人生效，之后修改该属性仅对当前身份生效, 可选值有: `busy`：忙碌, `free`：空闲
	Location         *GetCalendarEventRespEventLocation   `json:"location,omitempty"`           // 日程地点
	Color            int64                                `json:"color,omitempty"`              // 日程颜色，颜色RGB值的int32表示。仅对当前身份生效；客户端展示时会映射到色板上最接近的一种颜色；值为0或-1时默认跟随日历颜色。
	Reminders        []*GetCalendarEventRespEventReminder `json:"reminders,omitempty"`          // 日程提醒列表
	Recurrence       string                               `json:"recurrence,omitempty"`         // 重复日程的重复性规则；参考[rfc5545](https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10)；, 不支持COUNT和UNTIL同时出现；, 预定会议室重复日程长度不得超过两年。, 最大长度：`2000` 字符
	Status           string                               `json:"status,omitempty"`             // 日程状态, 可选值有: `tentative`：未回应, `confirmed`：已确认, `cancelled`：日程已取消
	IsException      bool                                 `json:"is_exception,omitempty"`       // 日程是否是一个重复日程的例外日程
	RecurringEventID string                               `json:"recurring_event_id,omitempty"` // 例外日程的原重复日程的event_id
	Schemas          []*GetCalendarEventRespEventSchema   `json:"schemas,omitempty"`            // 日程自定义信息
}

type GetCalendarEventRespEventStartTime struct {
	Date      string `json:"date,omitempty"`      // 仅全天日程使用该字段，如2018-09-01。需满足 RFC3339 格式。不能与 timestamp 同时指定
	Timestamp string `json:"timestamp,omitempty"` // 秒级时间戳，如1602504000(表示2020/10/12 20:0:00 +8时区)
	Timezone  string `json:"timezone,omitempty"`  // 时区名称，使用IANA Time Zone Database标准，如Asia/Shanghai；全天日程时区固定为UTC，非全天日程时区默认为Asia/Shanghai
}

type GetCalendarEventRespEventEndTime struct {
	Date      string `json:"date,omitempty"`      // 仅全天日程使用该字段，如2018-09-01。需满足 RFC3339 格式。不能与 timestamp 同时指定
	Timestamp string `json:"timestamp,omitempty"` // 秒级时间戳，如1602504000(表示2020/10/12 20:0:00 +8时区)
	Timezone  string `json:"timezone,omitempty"`  // 时区名称，使用IANA Time Zone Database标准，如Asia/Shanghai；全天日程时区固定为UTC，非全天日程时区默认为Asia/Shanghai
}

type GetCalendarEventRespEventVchat struct {
	VCType      string `json:"vc_type,omitempty"`     // 视频会议类型, 可选值有: `vc`：飞书视频会议，取该类型时，其他字段无效。, `third_party`：第三方链接视频会议，取该类型时，icon_type、description、meeting_url字段生效。, `no_meeting`：无视频会议，取该类型时，其他字段无效。, `lark_live`：Lark直播，内部类型，只读。, `unknown`：未知类型，做兼容使用，只读。
	IconType    string `json:"icon_type,omitempty"`   // 第三方视频会议icon类型；可以为空，为空展示默认icon。, 可选值有: `vc`：飞书视频会议icon, `live`：直播视频会议icon, `default`：默认icon
	Description string `json:"description,omitempty"` // 第三方视频会议文案，可以为空，为空展示默认文案, 长度范围：`0` ～ `500` 字符
	MeetingURL  string `json:"meeting_url,omitempty"` // 视频会议URL, 长度范围：`1` ～ `2000` 字符
}

type GetCalendarEventRespEventLocation struct {
	Name      string  `json:"name,omitempty"`      // 地点名称, 长度范围：`1` ～ `512` 字符
	Address   string  `json:"address,omitempty"`   // 地点地址, 长度范围：`1` ～ `255` 字符
	Latitude  float64 `json:"latitude,omitempty"`  // 地点坐标纬度信息，对于国内的地点，采用GCJ-02标准，海外地点采用WGS84标准
	Longitude float64 `json:"longitude,omitempty"` // 地点坐标经度信息，对于国内的地点，采用GCJ-02标准，海外地点采用WGS84标准
}

type GetCalendarEventRespEventReminder struct {
	Minutes int64 `json:"minutes,omitempty"` // 日程提醒时间的偏移量，正数时表示在日程开始前X分钟提醒，负数时表示在日程开始后X分钟提醒,新建或更新日程时传入该字段，仅对当前身份生效, 取值范围：`-20160` ～ `20160`
}

type GetCalendarEventRespEventSchema struct {
	UiName   string `json:"ui_name,omitempty"`   // UI名称。取值范围如下： \,ForwardIcon: 日程转发按钮 \,MeetingChatIcon: 会议群聊按钮 \,MeetingMinutesIcon: 会议纪要按钮 \,MeetingVideo: 视频会议区域 \,RSVP: 接受/拒绝/待定区域 \,Attendee: 参与者区域 \,OrganizerOrCreator: 组织者/创建者区域
	UiStatus string `json:"ui_status,omitempty"` // UI项自定义状态。目前只支持hide, 可选值有: `hide`：隐藏显示, `readonly`：只读, `editable`：可编辑, `unknown`：未知UI项自定义状态，仅用于读取时兼容
	AppLink  string `json:"app_link,omitempty"`  // 按钮点击后跳转的链接, 最大长度：`2000` 字符
}
