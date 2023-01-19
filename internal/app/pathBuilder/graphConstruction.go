package pathBuilder

import (
	"navigation/internal/database/client/postgresql"
	"navigation/internal/logging"
)

type pathBuilder struct {
	client postgresql.Client
	logger *logging.Logger
}

func NewPathBuilder(logger *logging.Logger, client postgresql.Client) *pathBuilder {
	return &pathBuilder{
		client: client,
		logger: logger,
	}
}

// you must provide a start sector and an end sector
func (p *pathBuilder) Builder(start, end uint) ([]int, error) {
	_, _ = p.adjacencyMatrix()
	return nil, nil
}

func (p *pathBuilder) adjacencyMatrix() (map[int][]int, error) {
	return nil, nil
}