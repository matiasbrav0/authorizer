package memory

import "github.com/mbravovaisma/authorizer/internal/core/ports"

type memory struct {
}

func NewMemory() ports.AuthorizerRepository {
	return &memory{}
}
