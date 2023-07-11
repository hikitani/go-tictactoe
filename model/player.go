package model

type PlayerType uint8

const (
	PlayerTypeCross PlayerType = iota
	PlayerTypeZero
	PlayerTypeNone
)

type Player struct {
	playerType PlayerType
	cells      []*Cell
}

func (p *Player) Type() PlayerType {
	return p.playerType
}

func NewPlayer(typ PlayerType) *Player {
	return &Player{
		playerType: typ,
	}
}
