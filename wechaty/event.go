package wechaty

import (
	"github.com/wechaty/go-wechaty/wechaty-puppet/schemas"
	_interface "github.com/wechaty/go-wechaty/wechaty/interface"
	"github.com/wechaty/go-wechaty/wechaty/user"
	"time"
)

type (
	// EventDong ...
	EventDong func(context *PluginContext, data string)
	// EventError ...
	EventError func(context *PluginContext, err error)
	// EventFriendship ...
	EventFriendship func(context *PluginContext, friendship *user.Friendship)
	// EventHeartbeat ...
	EventHeartbeat func(context *PluginContext, data string)
	// EventLogin ...
	EventLogin func(context *PluginContext, user *user.ContactSelf)
	// EventLogout ...
	EventLogout func(context *PluginContext, user *user.ContactSelf, reason string)
	// EventMessage ...
	EventMessage func(context *PluginContext, message *user.Message)
	// EventReady ...
	EventReady func(context *PluginContext)
	// EventRoomInvite ...
	EventRoomInvite func(context *PluginContext, roomInvitation *user.RoomInvitation)
	// EventRoomJoin ...
	EventRoomJoin func(context *PluginContext, room *user.Room, inviteeList []_interface.IContact, inviter _interface.IContact, date time.Time)
	// EventRoomLeave ...
	EventRoomLeave func(context *PluginContext, room *user.Room, leaverList []_interface.IContact, remover _interface.IContact, date time.Time)
	// EventRoomTopic ...
	EventRoomTopic func(context *PluginContext, room *user.Room, newTopic string, oldTopic string, changer _interface.IContact, date time.Time)
	// EventScan ...
	EventScan func(context *PluginContext, qrCode string, status schemas.ScanStatus, data string)
	// EventStart ...
	EventStart func(context *PluginContext)
	// EventStop ...
	EventStop func(context *PluginContext)
)
