package basics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Structs(t *testing.T) {
	t.Run("should add a new person", func(t *testing.T) {
		persons = nil

		p := Person{Name: "Maria", LastName: "Silva", Age: 21, Car: Car{"BD1", "Fake Brand"}}

		Add(p)

		assert.Equal(t, 1, len(GetAll()))
		assert.Equal(t, "Maria", p.Name)
		assert.Equal(t, "Fake Brand", p.Car.Brand)
	})

	t.Run("should get all persons", func(t *testing.T) {
		persons = nil

		Add(Person{Name: "Maria", LastName: "Silva", Age: 21})
		Add(Person{Name: "Maria", LastName: "Silva", Age: 21})
		Add(Person{Name: "Maria", LastName: "Silva", Age: 21})

		assert.Equal(t, 3, len(GetAll()))
	})

	t.Run("should update person name", func(t *testing.T) {
		persons = nil

		p := Person{Name: "Maria", LastName: "Silva", Age: 21, Car: Car{"BD1", "Fake Brand"}}

		Add(p)

		p.Update("Joana")

		assert.Equal(t, "Joana", p.Name)
	})
}
