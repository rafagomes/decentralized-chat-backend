package chat

type Service struct {
	//chatRepo ChatRepository
}

func (cs *Service) SendMessage(chat *Chat) error {
	return nil
}

func (cs *Service) GetMessages(channelID string) ([]*Chat, error) {
	return nil, nil
}
