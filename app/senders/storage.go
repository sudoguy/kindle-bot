package senders

import (
	"encoding/json"
	"errors"
	"os"
	"path"
	"strconv"
	"sync"

	"github.com/rs/zerolog/log"
)

const (
	// StorageFolder is a folder where stored books and senders.
	StorageFolder = "storage"
	// SenderFileName is a name of a file with sender information.
	SenderFileName = "info.json"
)

// SenderInfo represents information about a sender.
type SenderInfo struct {
	Email      string `json:"email,omitempty"`
	UserName   string `json:"username"`
	TelegramID int64  `json:"telegram_id"`
	BooksSent  int    `json:"sent_books"`
}

// Path returns a path to the sender folder.
func (sender *SenderInfo) Path() string {
	return path.Join(StorageFolder, strconv.FormatInt(sender.TelegramID, 10))
}

// Storage represents a storage for senders.
type Storage struct {
	mu sync.Mutex
}

// NewStorage creates a new storage.
func NewStorage() *Storage {
	storage := &Storage{}
	storage.prepareFolders()

	return storage
}

// Create storage folder if not exist
func (storage *Storage) prepareFolders() {
	_, err := os.Stat(StorageFolder)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(StorageFolder, 0o750)
		if errDir != nil {
			log.Panic().Err(errDir)
		}
	}
}

// GetSenderByID returns a sender by ID.
func (storage *Storage) GetSenderByID(id int64) (*SenderInfo, error) {
	storage.mu.Lock()
	defer storage.mu.Unlock()

	formattedID := strconv.FormatInt(id, 10)
	storage.createSenderFolderIfNeeded(formattedID)

	sender, err := storage.readSenderData(formattedID)

	return &sender, err
}

func (*Storage) readSenderData(formattedID string) (SenderInfo, error) {
	filePath := path.Join(StorageFolder, formattedID, SenderFileName)
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return SenderInfo{}, errors.New("sender not found")
	} else if err != nil {
		log.Error().Err(err)
		return SenderInfo{}, errors.New("error while checking sender")
	}

	data, err := os.ReadFile(filePath) // nolint:gosec  // path does not contain user's input
	if err != nil {
		log.Error().Err(err)
		return SenderInfo{}, errors.New("error while reading sender")
	}

	var sender SenderInfo
	err = json.Unmarshal(data, &sender)
	if err != nil {
		log.Error().Err(err)
		return SenderInfo{}, errors.New("error while unmarshalling sender")
	}

	return sender, nil
}

func (*Storage) createSenderFolderIfNeeded(formattedID string) {
	folderPath := path.Join(StorageFolder, formattedID)
	_, err := os.Stat(folderPath)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(folderPath, 0o750)
		if errDir != nil {
			log.Panic().Err(errDir)
		}
	}
}

// UpdateSender updates a sender.
func (storage *Storage) UpdateSender(sender *SenderInfo) error {
	storage.mu.Lock()
	defer storage.mu.Unlock()

	formattedID := strconv.FormatInt(sender.TelegramID, 10)
	storage.createSenderFolderIfNeeded(formattedID)

	data, err := json.Marshal(sender)
	if err != nil {
		log.Error().Err(err)
		return errors.New("error while marshaling sender")
	}

	filePath := path.Join(StorageFolder, formattedID, SenderFileName)
	err = os.WriteFile(filePath, data, 0o600)
	if err != nil {
		log.Error().Err(err)
		return errors.New("error while writing sender")
	}

	return nil
}
