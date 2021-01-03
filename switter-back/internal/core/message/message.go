package message

type Storage interface {
}

type MessageDispatcher struct {
}

func NewMessageDispatcher(s *Storage) MessageDispatcher {
	return MessageDispatcher{}
}
