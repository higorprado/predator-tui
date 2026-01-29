package ui

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/higorprado/predator-tui/internal/hardware"
)

const (
	defaultWidth  = 40
	defaultHeight = 12
)

type tickMsg time.Time
type updateMsg *hardware.State

// profileItem implements list.Item interface
type profileItem struct {
	name     string
	isActive bool
}

func (i profileItem) Title() string       { return i.name }
func (i profileItem) Description() string { return "" }
func (i profileItem) FilterValue() string { return i.name }

// itemDelegate handles rendering of list items
type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }

func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(profileItem)
	if !ok {
		return
	}

	name := strings.ToLower(i.name)
	var str string

	isSelected := index == m.Index()

	if isSelected {
		str = selectedItemStyle.Render("▸ " + name)
	} else {
		str = normalItemStyle.Render("  " + name)
	}

	// Add active badge with proper spacing
	if i.isActive {
		// Calculate padding to align the badge
		padding := 20 - len(name)
		if padding < 2 {
			padding = 2
		}
		str += strings.Repeat(" ", padding) + activeBadgeStyle.Render("● active")
	}

	fmt.Fprint(w, str)
}

type Model struct {
	hw     *hardware.Client
	state  *hardware.State
	list   list.Model
	status string
	isErr  bool
	width  int
	height int
}

func InitialModel() Model {
	// Create initial empty list
	delegate := itemDelegate{}
	l := list.New([]list.Item{}, delegate, defaultWidth, 5)

	// Configure list appearance
	l.SetShowTitle(false)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)
	l.SetShowPagination(false)
	l.InfiniteScrolling = true

	// Customize list styles
	l.Styles.NoItems = lipgloss.NewStyle().Foreground(subtle)

	return Model{
		hw:     hardware.NewClient(),
		state:  &hardware.State{},
		list:   l,
		width:  defaultWidth,
		height: defaultHeight,
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
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		// Adjust list dimensions
		containerWidth := min(msg.Width-4, 42)
		m.list.SetWidth(containerWidth - 4)
		return m, nil

	case updateMsg:
		if msg != nil {
			m.state = msg
			m.updateListItems()
		}
		return m, nil

	case tickMsg:
		return m, tea.Batch(fetchState(m.hw), tickCmd())

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		case "enter":
			if i, ok := m.list.SelectedItem().(profileItem); ok {
				err := m.hw.SetProfile(i.name)
				if err != nil {
					m.status = "Error: " + err.Error()
					m.isErr = true
				} else {
					m.status = "Success: " + i.name + " activated"
					m.isErr = false
					m.state.CurrentProfile = i.name
					m.updateListItems()
				}
				return m, nil
			}
		}
	}

	// Delegate to list for navigation
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *Model) updateListItems() {
	if m.state == nil {
		return
	}
	items := make([]list.Item, len(m.state.AvailableProfiles))
	for i, p := range m.state.AvailableProfiles {
		items[i] = profileItem{
			name:     p,
			isActive: p == m.state.CurrentProfile,
		}
	}
	m.list.SetItems(items)
}

func (m Model) View() string {
	var b strings.Builder

	// Title centered
	title := titleStyle.Render("PREDATOR CONTROL")
	b.WriteString(title)
	b.WriteString("\n\n")

	// Section header
	b.WriteString(sectionStyle.Render("Power Profiles"))
	b.WriteString("\n\n")

	// Profile list
	if m.state == nil || len(m.state.AvailableProfiles) == 0 {
		b.WriteString(errorStyle.Render("No profiles available"))
		b.WriteString("\n")
		b.WriteString(helpStyle.Render("Platform profile not found at /sys/firmware/acpi/platform_profile"))
	} else {
		b.WriteString(m.list.View())
	}
	b.WriteString("\n")

	// Status message
	if m.status != "" {
		b.WriteString("\n")
		if m.isErr {
			b.WriteString(errorStyle.Render("✗ " + m.status))
		} else {
			b.WriteString(successStyle.Render("✓ " + m.status))
		}
	}

	// Help text
	b.WriteString("\n")
	b.WriteString(helpStyle.Render("↑/↓ navigate • enter apply • q quit"))

	// Wrap in container
	content := containerStyle.Render(b.String())

	// Center the container
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		content,
	)
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
