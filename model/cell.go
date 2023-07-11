package model

type Cell struct {
	x, y   int
	player *Player
}

type CellState struct {
	X, Y  int
	State PlayerType
}
