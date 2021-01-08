package message

import (
	"errors"
	"log"
	"switter-back/internal/types"
)

var (
	ErrAccessDenied = errors.New("Access denied")
	ErrFail         = errors.New("Failed to get data")
	ErrNotCreated   = errors.New("Failed to create message")
	ErrNotFound     = errors.New("Message not found")
	ErrNotDeleted   = errors.New("Failed to delete message")
	ErrNotModified  = errors.New("Failed to modify message")
)

type Storage interface {
	CreateMessage(userID types.UserID, text string) error
	GetMessage(messageID types.MessageID) (types.Message, bool, error)
	GetMessageList(page int) ([]types.Message, error)
	UpdateMessage(ID types.MessageID, newText string) error
	DeleteMessage(userID types.UserID, ID types.MessageID) error
}

type MessageDispatcher struct {
	storage Storage
}

func New(s Storage) MessageDispatcher {
	return MessageDispatcher{storage: s}
}

func (md *MessageDispatcher) GetListPage(page int) ([]types.Message, error) {
	msgs, err := md.storage.GetMessageList(page)
	if err != nil {
		log.Println("Failed to create message")
		return []types.Message{}, ErrNotFound
	}
	return msgs, nil
}

func (md *MessageDispatcher) GetMessage(msgID types.MessageID) (types.Message, error) {
	msg, _, err := md.storage.GetMessage(msgID)
	if err != nil {
		log.Println("Failed to get message")
		return types.Message{}, ErrNotFound
	}
	return msg, nil
}

func (md *MessageDispatcher) CreateMessage(userID types.UserID, message string) error {
	err := md.storage.CreateMessage(userID, message)
	if err != nil {
		log.Println("Failed to create message")
		return ErrNotCreated
	}
	return nil
}

func (md *MessageDispatcher) UpdateMessage(userID types.UserID, msgID types.MessageID, message string) error {
	msg, _, err := md.storage.GetMessage(msgID)
	if err != nil {
		log.Println("Failed to get message")
		return ErrNotModified
	}
	if msg.UserID != userID {
		log.Println("Access denied to this user")
		return ErrNotModified
	}
	err = md.storage.UpdateMessage(msgID, message)
	if err != nil {
		log.Println("Failed to create message")
		return ErrNotModified
	}
	return nil
}

func (md *MessageDispatcher) DeleteMessage(userID types.UserID, msgID types.MessageID) error {
	msg, _, err := md.storage.GetMessage(msgID)
	if err != nil {
		log.Println("Failed to get message")
		return ErrNotDeleted
	}
	if msg.UserID != userID {
		log.Println("Access denied to this user")
		return ErrNotDeleted
	}
	err = md.storage.DeleteMessage(userID, msgID)
	if err != nil {
		log.Println("Failed to create message")
		return ErrNotDeleted
	}
	return nil
}
