package handlers

import "github.com/sudoguy/kindle-bot/app/utils"

var sharedStorage = utils.NewStorage()

func getStorage() *utils.Storage {
	return sharedStorage
}

// SetStorage allows tests to override the default storage
func SetStorage(custom *utils.Storage) {
	if custom == nil {
		sharedStorage = utils.NewStorage()
		return
	}

	sharedStorage = custom
}
