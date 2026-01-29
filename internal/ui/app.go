package ui

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/higorprado/predator-tui/internal/hardware"
)

type tickMsg time.Time
type updateMsg *hardware.State

type Model struct {
	hw            *hardware.Client
	state         *hardware.State
	status        string
	cursorProfile int
}

func InitialModel() Model {
	return Model{
		hw:    hardware.NewClient(),
		state: &hardware.State{},
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		fetchState(m.hw),
		tickCmd(),
	)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case updateMsg:
		m.state = msg

	case tickMsg:
		return m, tea.Batch(fetchState(m.hw), tickCmd())

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		case "up", "k":
			if m.cursorProfile > 0 {
				m.cursorProfile--
			}
		case "down", "j":
			if m.cursorProfile < len(m.state.AvailableProfiles)-1 {
				m.cursorProfile++
			}
		case "enter":
			if len(m.state.AvailableProfiles) > 0 {
				profile := m.state.AvailableProfiles[m.cursorProfile]
				err := m.hw.SetProfile(profile)
				if err != nil {
					m.status = "Erro: " + err.Error()
				} else {
					m.status = "Sucesso: " + profile + " ativado"
					m.state.CurrentProfile = profile // Atualização imediata visual
				}
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	s := strings.Builder{}

	s.WriteString(titleStyle.Render("PREDATOR CONTROL"))
	s.WriteString("\n\n")

	s.WriteString("SELECIONE O PERFIL:\n\n")
	for i, p := range m.state.AvailableProfiles {
		label := strings.ToUpper(p)
		if p == m.state.CurrentProfile {
			label += " [ATIVO]"
		}

		if i == m.cursorProfile {
			s.WriteString(selectedItem.Render("> "+label) + "\n")
		} else {
			s.WriteString(unselectedItem.Render("  "+label) + "\n")
		}
	}

	if m.status != "" {
		s.WriteString("\n" + statusStyle.Render("● "+m.status))
	}

	s.WriteString("\n\n" + infoStyle.Render("Setas: Mover • Enter: Aplicar • Q: Sair"))
	return lipgloss.NewStyle().Padding(1, 2).Render(s.String())
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func fetchState(hw *hardware.Client) tea.Cmd {
	return func() tea.Msg {
		s, _ := hw.GetState()
		return updateMsg(s)
	}
}
