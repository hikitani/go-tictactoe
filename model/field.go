package model

const Height, Width = 3, 3

type Field struct {
	grid [][]*Cell
}

func (f *Field) WhoIsWin() (*Player, bool) {
	for _, row := range f.grid {
		var lastPlayer *Player
		rowFound := true
		for _, cell := range row {
			if cell.player == nil {
				rowFound = false
				break
			}

			if cell.player.playerType == PlayerTypeNone {
				rowFound = false
				break
			}

			if lastPlayer == nil {
				lastPlayer = cell.player
			} else if cell.player.playerType != lastPlayer.playerType {
				rowFound = false
				break
			}
		}

		if rowFound {
			return lastPlayer, true
		}
	}

	for i := 0; i < len(f.grid[0]); i++ {
		var lastPlayer *Player
		rowFound := true
		for j := 0; j < len(f.grid); j++ {
			cell := f.grid[j][i]

			if cell.player == nil {
				rowFound = false
				break
			}

			if cell.player.playerType == PlayerTypeNone {
				rowFound = false
				break
			}

			if lastPlayer == nil {
				lastPlayer = cell.player
			} else if cell.player.playerType != lastPlayer.playerType {
				rowFound = false
				break
			}
		}

		if rowFound {
			return lastPlayer, true
		}
	}

	var lastPlayer *Player
	rowFound := true
	for i := 0; i < len(f.grid); i++ {
		cell := f.grid[i][i]

		if cell.player == nil {
			rowFound = false
			break
		}

		if cell.player.playerType == PlayerTypeNone {
			rowFound = false
			break
		}

		if lastPlayer == nil {
			lastPlayer = cell.player
		} else if cell.player.playerType != lastPlayer.playerType {
			rowFound = false
			break
		}
	}

	if rowFound {
		return lastPlayer, true
	}

	lastPlayer = nil
	rowFound = true
	for i := 0; i < len(f.grid); i++ {
		cell := f.grid[i][len(f.grid)-1-i]

		if cell.player == nil {
			rowFound = false
			break
		}

		if cell.player.playerType == PlayerTypeNone {
			rowFound = false
			break
		}

		if lastPlayer == nil {
			lastPlayer = cell.player
		} else if cell.player.playerType != lastPlayer.playerType {
			rowFound = false
			break
		}
	}

	if rowFound {
		return lastPlayer, true
	}

	gridIsFilled := true

OUT:
	for _, row := range f.grid {
		for _, cell := range row {
			if cell.player == nil {
				gridIsFilled = false
				break OUT
			}
		}
	}

	if gridIsFilled {
		return nil, true
	}

	return nil, false
}

func (f *Field) State() [][]CellState {
	var stateGrid [][]CellState
	for _, row := range f.grid {
		var rowState []CellState
		for _, cell := range row {
			state := PlayerTypeNone
			if cell.player != nil {
				state = cell.player.playerType
			}

			rowState = append(rowState, CellState{
				X:     cell.x,
				Y:     cell.y,
				State: state,
			})
		}

		stateGrid = append(stateGrid, rowState)
	}

	return stateGrid
}

func (f *Field) SetPlayer(x, y int, player *Player) {
	if x >= Width {
		panic("x > width")
	}

	if y >= Height {
		panic("y > height")
	}

	if player == nil {
		panic("player is nil")
	}

	cell := f.grid[y][x]

	if cell.player != nil {
		panic("cell is busy")
	}

	cell.player = player
	player.cells = append(player.cells, cell)
}

func NewField() *Field {
	var grid [][]*Cell
	for i := 0; i < Height; i++ {

		var row []*Cell
		for j := 0; j < Width; j++ {
			row = append(row, &Cell{
				x: i,
				y: j,
			})
		}

		grid = append(grid, row)
	}

	return &Field{
		grid: grid,
	}
}
