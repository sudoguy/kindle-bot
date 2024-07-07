package utils

import (
	"encoding/json"
	"errors"
	"os"
	"sync"

	"github.com/rs/zerolog/log"
)

// SenderInfo represents information about a sender.
type SenderInfo struct {
	Email      string `json:"email,omitempty"`
	UserName   string `json:"username"`
	TelegramID int64  `json:"telegram_id"`
	BooksSent  int    `json:"sent_books"`
}

// Storage represents a storage for senders.
type Storage struct {
	mu      sync.Mutex
	Senders []SenderInfo `json:"senders"`
}

// NewStorage creates a new storage.
func NewStorage() *Storage {
	storage := &Storage{}
	storage.openSenders()

	return storage
}

// GetSenderByID returns a sender by ID.
func (storage *Storage) GetSenderByID(id int64) (*SenderInfo, error) {
	storage.openSenders()
	for _, sender := range storage.Senders {
		if sender.TelegramID == id {
			return &sender, nil
		}
	}

	return &SenderInfo{}, errors.New("sender not found")
}

func (storage *Storage) openSenders() {
	data, err := os.ReadFile("./senders.json")
	if err != nil {
		log.Panic().Err(err)
	}

	var senders []SenderInfo
	err = json.Unmarshal(data, &senders)
	if err != nil {
		log.Panic().Err(err)
	}

	storage.Senders = senders
}

// UpdateSender updates a sender.
func (storage *Storage) UpdateSender(sender *SenderInfo) {
	storage.mu.Lock()
	defer storage.mu.Unlock()
	senders := storage.Senders

	for index, item := range senders {
		if item.TelegramID == sender.TelegramID {
			senders[index] = *sender
		}
	}

	storage.WriteSendersToFile()
}

// WriteSendersToFile writes senders to a file.
func (storage *Storage) WriteSendersToFile() {
	data, err := json.Marshal(storage.Senders)
	// print senders
	log.Info().Msg(string(data))

	if err != nil {
		log.Panic().Err(err)
	}

	err = os.WriteFile("./senders.json", data, 0o600)
	if err != nil {
		log.Panic().Err(err)
	}
}
