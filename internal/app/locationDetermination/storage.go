package locationDetermination

type Repository interface {
	GetSector(number string, building uint) (uint, error)
}
