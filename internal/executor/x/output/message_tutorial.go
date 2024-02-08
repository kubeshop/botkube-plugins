package output

import (
	"fmt"

	"github.com/kubeshop/botkube-cloud-plugins/internal/executor/x"
	"github.com/kubeshop/botkube-cloud-plugins/internal/executor/x/state"
	"github.com/kubeshop/botkube-cloud-plugins/internal/executor/x/template"
	"github.com/kubeshop/botkube/pkg/api"
	"github.com/kubeshop/botkube/pkg/mathx"
)

// TutorialWrapper allows constructing interactive message with predefined steps.
type TutorialWrapper struct{}

// NewTutorialWrapper returns a new TutorialWrapper instance.
func NewTutorialWrapper() *TutorialWrapper {
	return &TutorialWrapper{}
}

// RenderMessage returns  interactive message with predefined steps.
func (p *TutorialWrapper) RenderMessage(cmd, _ string, _ *state.Container, msgCtx *template.Template) (api.Message, error) {
	var (
		msg   = msgCtx.TutorialMessage
		start = mathx.Min(msg.Paginate.CurrentPage*msg.Paginate.Page, len(msg.Buttons)-2)
		stop  = mathx.Min(start+msg.Paginate.Page, len(msg.Buttons))
	)

	return api.Message{
		OnlyVisibleForYou: true,
		ReplaceOriginal:   msg.Paginate.CurrentPage > 0,
		Sections: []api.Section{
			{
				Base: api.Base{
					Header: msg.Header,
				},
			},
			{
				Buttons: msg.Buttons[start:stop],
			},
			{
				Buttons: p.getPaginationButtons(msg, msg.Paginate.CurrentPage, cmd),
			},
		},
	}, nil
}

func (p *TutorialWrapper) getPaginationButtons(msg template.TutorialMessage, pageIndex int, cmd string) []api.Button {
	allItems := len(msg.Buttons)
	if allItems <= msg.Paginate.Page {
		return nil
	}

	btnsBuilder := api.NewMessageButtonBuilder()

	var out []api.Button
	if pageIndex > 0 {
		out = append(out, btnsBuilder.ForCommandWithoutDesc("Prev", fmt.Sprintf("%s %s @page:%d", x.BuiltinCmdPrefix, cmd, mathx.DecreaseWithMin(pageIndex, 0))))
	}

	if pageIndex*msg.Paginate.Page < allItems-1 {
		out = append(out, btnsBuilder.ForCommandWithoutDesc("Next", fmt.Sprintf("%s %s @page:%d", x.BuiltinCmdPrefix, cmd, mathx.IncreaseWithMax(pageIndex, allItems-1)), api.ButtonStylePrimary))
	}
	return out
}
