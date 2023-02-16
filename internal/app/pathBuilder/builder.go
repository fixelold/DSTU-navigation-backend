package pathBuilder

// you must provide a start sector and an end sector
func (h *handler) Builder(start, end int) ([]int, error) {
	matrix, err := h.adjacencyMatrix()
	if err != nil {
		return nil, err
	}

	res := h.bfs(start, end, matrix)

	return res, nil
}

func (h *handler) bfs(start, end int, matrix map[int][]int) []int {
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
			result := h.getPath(end, top)
			return reverse(result)
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
	return result
}

func (h *handler) adjacencyMatrix() (map[int][]int, error) {
	matrix := make(map[int][]int)
	sectorLink, err := h.repository.GetSectorLink()
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
