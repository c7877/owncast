package chatrepository

import "github.com/owncast/owncast/storage/data"

type ChatRepository struct {
	datastore *data.Store
}

func New(datastore *data.Store) *ChatRepository {
	r := &ChatRepository{
		datastore: datastore,
	}

	return r
}

// NOTE: This is temporary during the transition period.
var temporaryGlobalInstance *ChatRepository

// GetUserRepository will return the user repository.
func GetChatRepository() *ChatRepository {
	if temporaryGlobalInstance == nil {
		i := New(data.GetDatastore())
		temporaryGlobalInstance = i
	}
	return temporaryGlobalInstance
}
