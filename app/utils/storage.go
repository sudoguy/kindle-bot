package utils

import (
	"encoding/json"
	"errors"
	"os"
	"sync"

	"github.com/rs/zerolog/log"
)

type SenderInfo struct {
	Email      string `json:"email,omitempty"`
	UserName   string `json:"username"`
	TelegramId int64  `json:"telegram_id"`
	BooksSent  int    `json:"sent_books"`
}

type Storage struct {
	mu      sync.Mutex
	Senders []SenderInfo `json:"senders"`
}

func NewStorage() *Storage {
	storage := &Storage{}
	storage.openSenders()

	return storage
}

func (storage *Storage) GetSenderById(id int64) (*SenderInfo, error) {
	storage.openSenders()
	for _, sender := range storage.Senders {
		if sender.TelegramId == id {
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

func (storage *Storage) UpdateSender(sender *SenderInfo) {
	storage.mu.Lock()
	defer storage.mu.Unlock()
	senders := storage.Senders

	for index, item := range senders {
		if item.TelegramId == sender.TelegramId {
			senders[index] = *sender
		}
	}

	storage.WriteSendersToFile()
}

func (storage *Storage) WriteSendersToFile() {
	data, err := json.Marshal(storage.Senders)
	// print senders
	log.Info().Msg(string(data))

	if err != nil {
		log.Panic().Err(err)
	}

	err = os.WriteFile("./senders.json", data, 0644)
	if err != nil {
		log.Panic().Err(err)
	}
}
