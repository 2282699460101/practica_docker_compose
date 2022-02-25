package history_service

import (
	"time"

	model "backend/models"
	historyRepository "backend/repositories/history.repository"
)

func Create(operacion model.Operacion) error {

	operacion.UpdatedAt = time.Now()
	operacion.CreatedAt = time.Now()

	err := historyRepository.Create(operacion)
	if err != nil {
		return err
	}
	return nil
}

func Read() (model.Historial, error) {
	historial, err := historyRepository.Read()
	if err != nil {
		return nil, err
	}
	return historial, nil
}
