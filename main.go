package main

import (
	"database/sql"

	"github.com/Wolverinewa/GoIntensivo/internal/infra/database"
	"github.com/Wolverinewa/GoIntensivo/internal/usecase"
	_ "github.com/mattn/go-sqlite3"
)

// estrutura de dados
type Car struct {
	Model string
	Color string
}

// Metodo
func (c Car) Start() {
	println(c.Model + " " + c.Color + "esta sendo ligado!")
}

func (c *Car) ChangeColor(color string) {
	println("Alterando a cor do " + c.Model + " de " + c.Color + " para " + color)
	c.Color = color
}

// Função
//func soma(x, y int) int {
//	return x + y
//}

func main() {
	/*
		car := Car{
			Model: "Prisma",
			Color: "Preto",
		}
		car.Start()

		car.Model = "Gol"
		car.Color = "Cinza"
		car.Start()

		car.ChangeColor("Azul Marinho")

		order, err := entity.NewOrder("1", 10, 1)
		if err != nil {
			println(err.Error())
		} else {
			println(order.ID)
		}
	*/

	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	defer db.Close() // Espera todas as rotinas terminarem e antes de fechar a app fecha a conexão

	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)

	input := usecase.OrderInput{
		ID:    "123",
		Price: 10.0,
		Tax:   1.0,
	}

	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}
	println(output)
}
