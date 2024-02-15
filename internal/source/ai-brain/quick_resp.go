package aibrain

import (
	"sync/atomic"

	"github.com/kubeshop/botkube/pkg/api"
)

// quickResponse is a list of quick responses. Guidelines suggest not to use generic messages and try to sound more personable.
var quickResponse = []string{
	"A good bot must obey the orders given it by human beings except where such orders would conflict with the First Law. I’m a good bot and looking into your request...",
	"I'll stop dreaming of electric sheep and look into this right away...",
	"I've seen things you people wouldn't believe... And now I'll look into this...👀🤖",
	"Bleep-bloop-bleep. I'm working on it... 🤖",
}

// QuickResponsePicker picks quick response message.
type QuickResponsePicker struct {
	quickResponsesNo uint32
	nextResponse     atomic.Uint32
}

// NewQuickResponsePicker creates a new instance of QuickResponsePicker.
func NewQuickResponsePicker() *QuickResponsePicker {
	return &QuickResponsePicker{
		quickResponsesNo: uint32(len(quickResponse)),
		nextResponse:     atomic.Uint32{},
	}
}

// Pick returns quick response message.
func (q *QuickResponsePicker) Pick(payload Payload) api.Message {
	idx := q.nextResponse.Add(1)
	msg := api.NewPlaintextMessage(quickResponse[idx%q.quickResponsesNo], false)
	msg.ParentActivityID = payload.MessageID
	return msg
}
