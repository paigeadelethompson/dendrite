package jetstream

import (
	"fmt"
	"regexp"
	"time"

	"github.com/nats-io/nats.go"
)

const (
	UserID        = "user_id"
	RoomID        = "room_id"
	EventID       = "event_id"
	RoomEventType = "output_room_event_type"
)

var (
	InputRoomEvent          = "InputRoomEvent"
	InputDeviceListUpdate   = "InputDeviceListUpdate"
	InputSigningKeyUpdate   = "InputSigningKeyUpdate"
	OutputRoomEvent         = "OutputRoomEvent"
	OutputAppserviceEvent   = "OutputAppserviceEvent"
	OutputSendToDeviceEvent = "OutputSendToDeviceEvent"
	OutputKeyChangeEvent    = "OutputKeyChangeEvent"
	OutputTypingEvent       = "OutputTypingEvent"
	OutputClientData        = "OutputClientData"
	OutputNotificationData  = "OutputNotificationData"
	OutputReceiptEvent      = "OutputReceiptEvent"
	OutputStreamEvent       = "OutputStreamEvent"
	OutputReadUpdate        = "OutputReadUpdate"
	RequestPresence         = "GetPresence"
	OutputPresenceEvent     = "OutputPresenceEvent"
	InputFulltextReindex    = "InputFulltextReindex"
)

var safeCharacters = regexp.MustCompile("[^A-Za-z0-9$]+")

func Tokenise(str string) string {
	return safeCharacters.ReplaceAllString(str, "_")
}

func InputRoomEventSubj(roomID string) string {
	return fmt.Sprintf("%s.%s", InputRoomEvent, Tokenise(roomID))
}

var streams = []*nats.StreamConfig{
	{
		Name:      InputRoomEvent,
		Retention: nats.InterestPolicy,
		Storage:   nats.FileStorage,
		MaxAge:    time.Hour * 24,
		MaxBytes:  1048576,
	},
	{
		Name:      InputDeviceListUpdate,
		Retention: nats.InterestPolicy,
		Storage:   nats.FileStorage,
		MaxBytes:  1048576,
	},
	{
		Name:      InputSigningKeyUpdate,
		Retention: nats.InterestPolicy,
		Storage:   nats.FileStorage,
		MaxBytes:  1048576,
	},
	{
		Name:      OutputRoomEvent,
		Retention: nats.InterestPolicy,
		Storage:   nats.FileStorage,
		MaxBytes:  1048576,
	},
	{
		Name:      OutputAppserviceEvent,
		Retention: nats.InterestPolicy,
		Storage:   nats.FileStorage,
		MaxBytes:  1048576,
	},
	{
		Name:      OutputSendToDeviceEvent,
		Retention: nats.InterestPolicy,
		Storage:   nats.FileStorage,
		MaxBytes:  1048576,
	},
	{
		Name:      OutputKeyChangeEvent,
		Retention: nats.InterestPolicy,
		Storage:   nats.FileStorage,
		MaxBytes:  1048576,
	},
	{
		Name:      OutputTypingEvent,
		Retention: nats.InterestPolicy,
		Storage:   nats.FileStorage,
		MaxAge:    time.Second * 60,
		MaxBytes:  1048576,
	},
	{
		Name:      OutputClientData,
		Retention: nats.InterestPolicy,
		Storage:   nats.FileStorage,
		MaxBytes:  1048576,
	},
	{
		Name:      OutputReceiptEvent,
		Retention: nats.InterestPolicy,
		Storage:   nats.FileStorage,
		MaxBytes:  1048576,
	},
	{
		Name:      OutputNotificationData,
		Retention: nats.InterestPolicy,
		Storage:   nats.FileStorage,
		MaxBytes:  1048576,
	},
	{
		Name:      OutputPresenceEvent,
		Retention: nats.InterestPolicy,
		Storage:   nats.MemoryStorage,
		MaxAge:    time.Minute * 5,
		MaxBytes:  1048576,
	},
}
