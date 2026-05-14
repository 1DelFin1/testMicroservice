package infra

import (
	"github.com/google/uuid"
)

type UUIDGenerator struct{}

func (u UUIDGenerator) NewID() string {
	return uuid.NewString()
}
