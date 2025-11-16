package handlers

import "github.com/sudoguy/kindle-bot/app/senders"

var sharedStorage = senders.NewStorage()

func getStorage() *senders.Storage {
	return sharedStorage
}

// SetStorage allows tests to override the default storage
func SetStorage(custom *senders.Storage) {
	if custom == nil {
		sharedStorage = senders.NewStorage()
		return
	}

	sharedStorage = custom
}
