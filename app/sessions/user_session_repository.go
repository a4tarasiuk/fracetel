package sessions

import (
	"errors"
)

var SessionDoesNotExist = errors.New("user session does not exist")

type UserSessionRepository interface {
	Create(session UserSession) error

	Update(session *UserSession) error

	GetByExternalID(externalID uint64) (*UserSession, error)
}
