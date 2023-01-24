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
func (p *pathBuilder) Builder(start, end int) ([]int, error) {
	matrix, err := p.adjacencyMatrix()
	if err != nil {
		return nil, err
	}

	res := p.bfs(start, end, matrix) 

	return res, nil
}

func (p *pathBuilder) bfs(start, end int, matrix map[int][]int) []int {
	var queue []int
	visited := make(map[int]bool)
	distance := make(map[int]int)
	top := make(map[int]int) // top of the graph
	d := 0

	queue = append(queue, start)
	visited[start] = true
	distance[start] = 0
	top[start] = 0

	for i := 0; i < len(matrix); i++ {
		d += 1
		current := queue[0]

		if current == end {
			result := p.getPath(end, top)
			return reverse(result)
		}

		for _, neighbor := range(matrix[current]) {
			if visited[neighbor] {
				continue
			}

			visited[neighbor] = true
			queue = append(queue, neighbor)
			distance[neighbor] = d
			top[neighbor] = current
		}

		queue = queue[1:]
	}

	result := p.getPath(end, top)
	return result
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

func (p *pathBuilder) getPath(end int, sectors map[int]int) []int {
	b := true
	var res []int
	for b {
		res = append(res, end)
		if sectors[end] == 0 {
			b = false
			break
		}

		end = sectors[end]
	}

	return res
}

func reverse(array []int) []int {
	var res []int
	for i := len(array) -1; i >= 0; i-- {
		res = append(res, array[i])
	}

	return res
}