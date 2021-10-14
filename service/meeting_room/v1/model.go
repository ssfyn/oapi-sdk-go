// Code generated by lark suite oapi sdk gen
package v1

import (
	"github.com/larksuite/oapi-sdk-go/api/core/tools"
	"github.com/larksuite/oapi-sdk-go/event/core/model"
)

type Building struct {
	BuildingId      string   `json:"building_id,omitempty"`
	Description     string   `json:"description,omitempty"`
	Floors          []string `json:"floors,omitempty"`
	Name            string   `json:"name,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *Building) MarshalJSON() ([]byte, error) {
	type cp Building
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type BuildingId struct {
	BuildingId       string   `json:"building_id,omitempty"`
	CustomBuildingId string   `json:"custom_building_id,omitempty"`
	ForceSendFields  []string `json:"-"`
}

func (s *BuildingId) MarshalJSON() ([]byte, error) {
	type cp BuildingId
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type Country struct {
	CountryId       string   `json:"country_id,omitempty"`
	Name            string   `json:"name,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *Country) MarshalJSON() ([]byte, error) {
	type cp Country
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type District struct {
	DistrictId      string   `json:"district_id,omitempty"`
	Name            string   `json:"name,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *District) MarshalJSON() ([]byte, error) {
	type cp District
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type ErrorEventUid struct {
	Uid             string   `json:"uid,omitempty"`
	OriginalTime    int      `json:"original_time,omitempty"`
	ErrorMsg        string   `json:"error_msg,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *ErrorEventUid) MarshalJSON() ([]byte, error) {
	type cp ErrorEventUid
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type EventInfo struct {
	Uid             string   `json:"uid,omitempty"`
	OriginalTime    int      `json:"original_time,omitempty"`
	Summary         string   `json:"summary,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *EventInfo) MarshalJSON() ([]byte, error) {
	type cp EventInfo
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type EventUid struct {
	Uid             string   `json:"uid,omitempty"`
	OriginalTime    int      `json:"original_time,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *EventUid) MarshalJSON() ([]byte, error) {
	type cp EventUid
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type OrganizerInfo struct {
	Name            string   `json:"name,omitempty"`
	OpenId          string   `json:"open_id,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *OrganizerInfo) MarshalJSON() ([]byte, error) {
	type cp OrganizerInfo
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type Room struct {
	RoomId          string   `json:"room_id,omitempty"`
	BuildingId      string   `json:"building_id,omitempty"`
	BuildingName    string   `json:"building_name,omitempty"`
	Capacity        int      `json:"capacity,omitempty"`
	Description     string   `json:"description,omitempty"`
	DisplayId       string   `json:"display_id,omitempty"`
	FloorName       string   `json:"floor_name,omitempty"`
	IsDisabled      bool     `json:"is_disabled,omitempty"`
	Name            string   `json:"name,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *Room) MarshalJSON() ([]byte, error) {
	type cp Room
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type RoomFreeBusy struct {
	StartTime       string         `json:"start_time,omitempty"`
	EndTime         string         `json:"end_time,omitempty"`
	Uid             string         `json:"uid,omitempty"`
	OriginalTime    int            `json:"original_time,omitempty"`
	OrganizerInfo   *OrganizerInfo `json:"organizer_info,omitempty"`
	ForceSendFields []string       `json:"-"`
}

func (s *RoomFreeBusy) MarshalJSON() ([]byte, error) {
	type cp RoomFreeBusy
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type RoomId struct {
	RoomId          string   `json:"room_id,omitempty"`
	CustomRoomId    string   `json:"custom_room_id,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *RoomId) MarshalJSON() ([]byte, error) {
	type cp RoomId
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type MeetingRoom struct {
	RoomId          int      `json:"room_id,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *MeetingRoom) MarshalJSON() ([]byte, error) {
	type cp MeetingRoom
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type BuildingBatchGetResult struct {
	Buildings []*Building `json:"buildings,omitempty"`
}

type BuildingBatchGetIdResult struct {
	Buildings []*BuildingId `json:"buildings,omitempty"`
}

type BuildingCreateReqBody struct {
	Name             string   `json:"name,omitempty"`
	Floors           []string `json:"floors,omitempty"`
	CountryId        string   `json:"country_id,omitempty"`
	DistrictId       string   `json:"district_id,omitempty"`
	CustomBuildingId string   `json:"custom_building_id,omitempty"`
	ForceSendFields  []string `json:"-"`
}

func (s *BuildingCreateReqBody) MarshalJSON() ([]byte, error) {
	type cp BuildingCreateReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type BuildingCreateResult struct {
	BuildingId string `json:"building_id,omitempty"`
}

type BuildingDeleteReqBody struct {
	BuildingId      string   `json:"building_id,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *BuildingDeleteReqBody) MarshalJSON() ([]byte, error) {
	type cp BuildingDeleteReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type BuildingListResult struct {
	PageToken string      `json:"page_token,omitempty"`
	HasMore   bool        `json:"has_more,omitempty"`
	Buildings []*Building `json:"buildings,omitempty"`
}

type BuildingUpdateReqBody struct {
	BuildingId       string   `json:"building_id,omitempty"`
	Name             string   `json:"name,omitempty"`
	Floors           []string `json:"floors,omitempty"`
	CountryId        string   `json:"country_id,omitempty"`
	DistrictId       string   `json:"district_id,omitempty"`
	CustomBuildingId string   `json:"custom_building_id,omitempty"`
	ForceSendFields  []string `json:"-"`
}

func (s *BuildingUpdateReqBody) MarshalJSON() ([]byte, error) {
	type cp BuildingUpdateReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type CountryListResult struct {
	Countries []*Country `json:"countries,omitempty"`
}

type DistrictListResult struct {
	Districts []*District `json:"districts,omitempty"`
}

type FreebusyBatchGetResult struct {
	TimeMin  string                     `json:"time_min,omitempty"`
	TimeMax  string                     `json:"time_max,omitempty"`
	FreeBusy map[string][]*RoomFreeBusy `json:"free_busy,omitempty"`
}

type InstanceReplyReqBody struct {
	RoomId          string   `json:"room_id,omitempty"`
	Uid             string   `json:"uid,omitempty"`
	OriginalTime    int      `json:"original_time,omitempty"`
	Status          string   `json:"status,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *InstanceReplyReqBody) MarshalJSON() ([]byte, error) {
	type cp InstanceReplyReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type RoomBatchGetResult struct {
	Rooms []*Room `json:"rooms,omitempty"`
}

type RoomBatchGetIdResult struct {
	Rooms []*RoomId `json:"rooms,omitempty"`
}

type RoomCreateReqBody struct {
	BuildingId      string   `json:"building_id,omitempty"`
	Floor           string   `json:"floor,omitempty"`
	Name            string   `json:"name,omitempty"`
	Capacity        int      `json:"capacity,omitempty"`
	IsDisabled      bool     `json:"is_disabled,omitempty"`
	CustomRoomId    string   `json:"custom_room_id,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *RoomCreateReqBody) MarshalJSON() ([]byte, error) {
	type cp RoomCreateReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type RoomCreateResult struct {
	RoomId string `json:"room_id,omitempty"`
}

type RoomDeleteReqBody struct {
	RoomId          string   `json:"room_id,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *RoomDeleteReqBody) MarshalJSON() ([]byte, error) {
	type cp RoomDeleteReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type RoomListResult struct {
	PageToken string  `json:"page_token,omitempty"`
	HasMore   bool    `json:"has_more,omitempty"`
	Rooms     []*Room `json:"rooms,omitempty"`
}

type RoomUpdateReqBody struct {
	RoomId          string   `json:"room_id,omitempty"`
	Name            string   `json:"name,omitempty"`
	Capacity        int      `json:"capacity,omitempty"`
	IsDisabled      bool     `json:"is_disabled,omitempty"`
	CustomRoomId    string   `json:"custom_room_id,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *RoomUpdateReqBody) MarshalJSON() ([]byte, error) {
	type cp RoomUpdateReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type SummaryBatchGetReqBody struct {
	EventUids       []*EventUid `json:"EventUids,omitempty"`
	ForceSendFields []string    `json:"-"`
}

func (s *SummaryBatchGetReqBody) MarshalJSON() ([]byte, error) {
	type cp SummaryBatchGetReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type SummaryBatchGetResult struct {
	EventInfos     []*EventInfo     `json:"EventInfos,omitempty"`
	ErrorEventUids []*ErrorEventUid `json:"ErrorEventUids,omitempty"`
}

type MeetingRoomStatusChangedEventData struct {
	RoomName string `json:"room_name,omitempty"`
	RoomId   string `json:"room_id,omitempty"`
}

type MeetingRoomStatusChangedEvent struct {
	*model.BaseEventV2
	Event *MeetingRoomStatusChangedEventData `json:"event"`
}

type MeetingRoomCreatedEventData struct {
	RoomName string `json:"room_name,omitempty"`
	RoomId   string `json:"room_id,omitempty"`
}

type MeetingRoomCreatedEvent struct {
	*model.BaseEventV2
	Event *MeetingRoomCreatedEventData `json:"event"`
}

type MeetingRoomDeletedEventData struct {
	RoomName string `json:"room_name,omitempty"`
	RoomId   string `json:"room_id,omitempty"`
}

type MeetingRoomDeletedEvent struct {
	*model.BaseEventV2
	Event *MeetingRoomDeletedEventData `json:"event"`
}

type MeetingRoomUpdatedEventData struct {
	RoomName string `json:"room_name,omitempty"`
	RoomId   string `json:"room_id,omitempty"`
}

type MeetingRoomUpdatedEvent struct {
	*model.BaseEventV2
	Event *MeetingRoomUpdatedEventData `json:"event"`
}

type RoomCreatedEventData struct {
	RoomId   string `json:"room_id,omitempty"`
	RoomName string `json:"room_name,omitempty"`
}

type RoomCreatedEvent struct {
	*model.BaseEventV2
	Event *RoomCreatedEventData `json:"event"`
}

type RoomDeletedEventData struct {
	RoomId   string `json:"room_id,omitempty"`
	RoomName string `json:"room_name,omitempty"`
}

type RoomDeletedEvent struct {
	*model.BaseEventV2
	Event *RoomDeletedEventData `json:"event"`
}

type RoomStatusChangedEventData struct {
	RoomId   string `json:"room_id,omitempty"`
	RoomName string `json:"room_name,omitempty"`
}

type RoomStatusChangedEvent struct {
	*model.BaseEventV2
	Event *RoomStatusChangedEventData `json:"event"`
}

type RoomUpdatedEventData struct {
	RoomId   string `json:"room_id,omitempty"`
	RoomName string `json:"room_name,omitempty"`
}

type RoomUpdatedEvent struct {
	*model.BaseEventV2
	Event *RoomUpdatedEventData `json:"event"`
}
