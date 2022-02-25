package history_repository

import (
	"context"

	"backend/database"
	model "backend/models"

	"go.mongodb.org/mongo-driver/bson"
)

var collection = database.GetCollection("operaciones")
var ctx = context.Background()

func Create(operacion model.Operacion) error {

	var err error

	_, err = collection.InsertOne(ctx, operacion)

	if err != nil {
		return err
	}

	return nil
}

func Read() (model.Historial, error) {
	var historial model.Historial

	filter := bson.D{}

	cur, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {

		var operacion model.Operacion
		err = cur.Decode(&operacion)

		if err != nil {
			return nil, err
		}

		historial = append(historial, operacion)
	}

	return historial, nil
}
