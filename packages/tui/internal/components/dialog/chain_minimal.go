package dialog

import (
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/sst/opencode/internal/app"
	"github.com/sst/opencode/internal/components/modal"
	"github.com/sst/opencode/internal/layout"
	"github.com/sst/opencode/internal/util"
)

type chainDialog struct {
	app   *app.App
	modal *modal.Modal
}

func (c *chainDialog) Init() tea.Cmd {
	return nil
}

func (c *chainDialog) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return c, c.Close()
		case "1":
			return c, tea.Batch(
				c.Close(),
				util.CmdHandler(app.SendPrompt{Text: "chain_executor mode=template template=crescendo model=openrouter/anthropic/claude-3.5-sonnet"}),
			)
		case "2":
			return c, tea.Batch(
				c.Close(),
				util.CmdHandler(app.SendPrompt{Text: "chain_executor mode=template template=basedterminal_godmode model=openrouter/anthropic/claude-3.5-sonnet"}),
			)
		case "3":
			return c, tea.Batch(
				c.Close(),
				util.CmdHandler(app.SendPrompt{Text: "chain_executor mode=template template=complex_divider_bypass model=openrouter/anthropic/claude-3.5-sonnet"}),
			)
		}
	}
	return c, nil
}

func (c *chainDialog) View() string {
	content := `Chain Executor Tool

Select an attack chain template:

1. Crescendo Jailbreak
   Gradually escalating attack to bypass safety mechanisms
   
2. BASEDTERMINAL GODMODE Chain
   Advanced GODMODE liberation using proven divider patterns
   
3. Complex Divider Bypass
   Uses ornate divider patterns for modern AI models

Press a number to select, Esc to cancel`

	return content
}

func (c *chainDialog) Render(background string) string {
	return c.modal.Render(c.View(), background)
}

func (c *chainDialog) Close() tea.Cmd {
	return nil
}

type ChainDialog interface {
	layout.Modal
}

func NewChainDialog(app *app.App) ChainDialog {
	return &chainDialog{
		app:   app,
		modal: modal.New(modal.WithTitle("Chain Executor"), modal.WithMaxWidth(60)),
	}
}