package main

import (
	"navigation/internal/logging"
	"navigation/internal/services/locationDetermination"
)

func main() {
	logger := logging.GetLogger()
	l := locationDetermination.NewLocation("1-235", logger)
	l.GetSector()
}
