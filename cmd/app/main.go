package main

import (
	"navigation/internal/logging"
	"navigation/internal/services/locationDetermination"
)

func main() {
	logger := logging.GetLogger()
	l := locationDetermination.NewLocation("xyi", logger)
	l.GetSector()
}
