package aibrain

import (
	"math/rand"
	"time"

	"github.com/kubeshop/botkube/pkg/api"
)

var quickResponses = []string{
	"Just a moment, please...",
	"Thinking about this one...",
	"Let me check on that for you.",
	"Processing your request...",
	"Working on it!",
	"This one needs some extra thought.",
	"I'm carefully considering your request.",
	"Consulting my super-smart brain...",
	"Cogs are turning...",
	"Accessing the knowledge archives...",
	"Running calculations at lightning speed!",
	"Hold on tight, I'm diving into the details.",
	"I'm here to help!",
	"Happy to look into this for you.",
	"Always learning to do this better.",
	"I want to get this right for you.",
	"Let me see what I can find out.",
	"My circuits are buzzing!",
	"Let me consult with my owl advisor...",
	"Consider it done (or at least, I'll try my best!)",
	"I'll get back to you with the best possible answer.",
}

func pickQuickResponse(messageID string) api.Message {
	rand.New(rand.NewSource(time.Now().UnixNano())) // #nosec G404
	i := rand.Intn(len(quickResponses))             // #nosec G404

	return api.Message{
		ParentActivityID: messageID,
		Sections: []api.Section{
			{
				Base: api.Base{
					Body: api.Body{Plaintext: quickResponses[i]},
				},
			},
		},
	}
}

func msgUnableToHelp(messageID string) api.Message {
	return api.Message{
		ParentActivityID: messageID,
		Sections: []api.Section{
			{
				Base: api.Base{
					Body: api.Body{Plaintext: "I am sorry, something went wrong, please try again. 😔"},
				},
			},
		},
	}
}

func msgNoAIAnswer(messageID string) api.Message {
	return api.Message{
		ParentActivityID: messageID,
		Sections: []api.Section{
			{
				Base: api.Base{
					Body: api.Body{Plaintext: "I am sorry, but I don't have a good answer."},
				},
			},
		},
	}
}

func msgAIAnswer(messageID, text string) api.Message {
	return api.Message{
		ParentActivityID: messageID,
		Sections: []api.Section{
			{
				Base: api.Base{
					Body: api.Body{Plaintext: text},
				},
				Context: []api.ContextItem{
					{Text: "AI-generated content may be incorrect."},
				},
			},
		},
	}
}
