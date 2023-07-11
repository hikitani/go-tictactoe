package screen

import (
	"bytes"
	"fmt"
	"go-tictactoe/model"
	"io"
)

type Field struct {
	f *model.Field
	w io.Writer
}

func (f *Field) Draw() error {
	grid := f.f.State()
	var b bytes.Buffer

	for _, row := range grid {
		b.WriteByte('|')

		for _, state := range row {
			switch state.State {
			case model.PlayerTypeNone:
				b.WriteByte(' ')
			case model.PlayerTypeZero:
				b.WriteByte('O')
			case model.PlayerTypeCross:
				b.WriteByte('X')
			default:
				panic("player type doesnt exist")
			}

			b.WriteByte('|')
		}

		b.WriteByte('\n')
		b.WriteByte('-')
		for range row {
			b.WriteString("--")
		}

		b.WriteByte('\n')
	}

	_, err := f.w.Write(b.Bytes())
	if err != nil {
		return fmt.Errorf("write: %w", err)
	}

	return nil
}

func NewField(field *model.Field, out io.Writer) *Field {
	return &Field{
		f: field,
		w: out,
	}
}
