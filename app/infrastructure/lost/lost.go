package megaLost

import (
	"ml-elizabeth/app/domain/repository"
)

type asdlost struct {
	nuevo string
}

func NewLost() repository.RandomRepository {
	return &asdlost{
		nuevo: "ddd",
	}
}

func (l *asdlost) RandomFunction(string) string {
	return "RandomFunction"
}
