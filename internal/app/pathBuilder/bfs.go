package pathBuilder

import (
	"navigation/internal/appError"
)

// you must provide a start sector and an end sector
func (h *handler) Builder(start, end, transitionSector int, TypeTranstionSector int) ([]int, appError.AppError) {
	var err appError.AppError
	matrix, err := h.adjacencyMatrix()
	if err.Err != nil {
		err.Wrap("Builder")
		return nil, err
	}

	res, q := h.bfs(start, end, transitionSector, matrix, TypeTranstionSector)
	q = append(q, res...)
	if TypeTranstionSector == elevator {
		q = append(q, transitionSector)
	}
	return q, err
}

func (h *handler) bfs(start, end, transitionSector int, matrix map[int][]int, TypeTranstionSector int) ([]int, []int) {
	var queue []int
	var q []int
	visited := make(map[int]bool)
	distance := make(map[int]int)
	top := make(map[int]int) // top of the graph
	d := 0

	// if transitionSector != 0 {
	// 	q = append(queue, start)
	// 	q = append(queue, transitionSector)
	// 	start = (start % 10) + (end / 10 * 10)
	// }

	if TypeTranstionSector == stairs {
		q = append(q, start)
		q = append(q, transitionSector)
		return q, nil 
		// start = (start % 10) + (end / 10 * 10)
		// fmt.Println("posle: ", q, start, queue)
	} else if TypeTranstionSector == elevator {
		end = 100 + transitionSector % 100
	}
	queue = append(queue, start)
	visited[start] = true
	distance[start] = 0
	top[start] = 0

	for i := 0; i < len(matrix); i++ {
		d += 1
		current := queue[0]

		if current == end {
			result := h.getPath(end, top)
			return reverse(result), q
		}

		for _, neighbor := range matrix[current] {
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

	result := h.getPath(end, top)
	return result, q
}

func (h *handler) adjacencyMatrix() (map[int][]int, appError.AppError) {
	matrix := make(map[int][]int)
	sectorLink, err := h.repository.GetSectorLink()
	if err.Err != nil {
		err.Wrap("adjacencyMatrix")
		return nil, err
	}

	for i := 0; i < len(sectorLink); i++ {
		sector := sectorLink[i].NumberSector
		link := sectorLink[i].NumberLink
		matrix[sector] = append(matrix[sector], link)
	}

	return matrix, err
}

func (h *handler) getPath(end int, sectors map[int]int) []int {
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
	for i := len(array) - 1; i >= 0; i-- {
		res = append(res, array[i])
	}

	return res
}
