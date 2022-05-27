package pages

type messageType string

const (
	SuccessMessage messageType = "success"
	ErrorMessage   messageType = "failure"
)

type flashMessage struct {
	Content     string
	MessageType messageType
}

type FlashMessages struct {
	messages []flashMessage
}

func NewFlashMessages() *FlashMessages {
	return &FlashMessages{}
}

func (m *FlashMessages) Add(content string, t messageType) {
	m.messages = append(
		m.messages,
		flashMessage{Content: content, MessageType: t},
	)
}

func (m *FlashMessages) HasErrorMessages() bool {
	for _, message := range m.messages {
		if message.MessageType == ErrorMessage {
			return true
		}
	}

	return false
}
