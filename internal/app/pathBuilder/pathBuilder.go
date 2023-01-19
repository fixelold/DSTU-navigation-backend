package pathBuilder

import (
	"context"
	"errors"
	"fmt"
	"navigation/internal/database/client/postgresql"
	"navigation/internal/logging"
	"navigation/internal/models"

	"github.com/jackc/pgconn"
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
	matrix := make(map[int][]int)
	sectorLink, err := p.getSectorLink()
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(sectorLink); i++ {
		sector := sectorLink[i].NumberSector
		link := sectorLink[i].NumberLink
		matrix[sector] = append(matrix[sector], link)
	}

	return matrix, nil
}

func (p *pathBuilder) getSectorLink() ([]models.SectorLink, error) {
	var sectorLink []models.SectorLink
	req := `SELECT number_sector, link FROM sector_link;`

	tx, err := p.client.Begin(context.Background())
	if err != nil {
		_ = tx.Rollback(context.Background())
		p.logger.Tracef("can't start transaction: %s", err.Error())
		return nil, err
	}

	rows, err := tx.Query(context.Background(), req)

	if err != nil {
		_ = tx.Rollback(context.Background())
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf("SQL Error: %s, Detail: %s, Where %s, Code: %s, SQLState: %s",
				pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			p.logger.Error(newErr)
			return nil, newErr
		}
		p.logger.Error(err)
		return nil, err
	}
	_ = tx.Commit(context.Background())

	for rows.Next() {
		var sl models.SectorLink
		err := rows.Scan(&sl.NumberSector, &sl.NumberLink)
		if err != nil {
			p.logger.Errorf("getSectorLink function. Scan error: %s", err.Error())
			return nil, err
		}
		sectorLink = append(sectorLink, sl)
	}

	return sectorLink, nil
}
