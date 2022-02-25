package models

import "time"

type Operacion struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	ValorIzq  string    `json:"valor_izq"`
	ValorDer  string    `json:"valor_der"`
	Operador  string    `json:"operador"`
	Resultado string    `json:"resultado"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at,omitempty"`
}

type Historial []Operacion
