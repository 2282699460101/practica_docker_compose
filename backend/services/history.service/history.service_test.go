package history_service_test

import (
	"testing"
	historialService "backend/history.service"
	model "backend/models"
	"time"
)

func TestCreate(t *testing.T) {
	operacion := model.Operacion{
		ValorIzq : "10",
		ValorDer: "20",
		Operador: "SUMA",
		Resultado: "30",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now()
	}
	err := historialService.Create(operacion)
	if err!= nul{
		t.Error("La prueba de persistencia de datos a fallado")
		t.Fail()
	}else{

	}
}

func TestRead(t *testing.T) {

}
