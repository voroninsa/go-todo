package authentication

import (
	"github.com/voroninsa/go-todo/storage"
	"github.com/voroninsa/go-todo/utils/dto"
)

func GetCredentials(req dto.AuthRequest) *storage.Credentials {
	switch req.Auth_type {
	case "jwt":
		return &storage.Credentials{}
	default:
		return nil
	}
}
