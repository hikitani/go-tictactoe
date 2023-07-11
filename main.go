package main

import (
	"fmt"
	"go-tictactoe/model"
	"go-tictactoe/screen"
	"os"
	"strconv"
)

func main() {
	field := model.NewField()
	fieldScreen := screen.NewField(field, os.Stdout)

	redraw := func() {
		fmt.Println("Текущее поле:")
		fieldScreen.Draw()
	}

	playerZero := model.NewPlayer(model.PlayerTypeZero)
	playerCross := model.NewPlayer(model.PlayerTypeCross)

	for {
		redraw()

		fmt.Println("Ход делает крестик")
		crossPosX := readPos("X", model.Width)
		crossPosY := readPos("Y", model.Height)
		field.SetPlayer(crossPosX, crossPosY, playerCross)
		redraw()
		if p, ok := field.WhoIsWin(); ok {
			sayWhoIsWin(p)
			os.Exit(0)
		}

		fmt.Println("Ход делает нолик")
		zeroPosX := readPos("X", model.Width)
		zeroPosY := readPos("Y", model.Height)
		field.SetPlayer(zeroPosX, zeroPosY, playerZero)
		redraw()
		if p, ok := field.WhoIsWin(); ok {
			sayWhoIsWin(p)
			os.Exit(0)
		}
	}
}

func readPos(posName string, maxVal int) int {
	var playerPos int
	for {
		fmt.Print("Укажите координату " + posName + ": ")
		var s string
		if _, err := fmt.Scan(&s); err != nil {
			panic("Ошибка при чтении команды: " + err.Error())
		}

		var err error
		playerPos, err = strconv.Atoi(s)
		if err != nil {
			fmt.Println("Введите корректную позицию")
			continue
		}

		if playerPos > maxVal {
			fmt.Printf("Позиция не больше %d\n", maxVal)
			continue
		}

		if playerPos < 1 {
			fmt.Println("Позиция не меньше 1")
			continue
		}

		break
	}

	return playerPos - 1
}

func sayWhoIsWin(p *model.Player) {
	if p == nil {
		fmt.Println("Вышла ничья")
		return
	}

	switch p.Type() {
	case model.PlayerTypeCross:
		fmt.Println("Выиграл крестик")
	case model.PlayerTypeZero:
		fmt.Println("Выиграл нолик")
	}
}
