package chat

type Repository struct{}

func (cr *Repository) Save(chat *Chat) error {
	return nil
}

func (cr *Repository) FindByChannelID(channelID string) ([]*Chat, error) {
	return nil, nil
}
