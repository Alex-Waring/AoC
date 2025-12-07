package utils

import "fmt"

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right

	UpLeft
	UpRight
	DownLeft
	DownRight

	DontMove
)

func (d Direction) Turn(turn Direction) Direction {
	if turn != Left && turn != Right {
		panic("should be left or right")
	}

	switch d {
	case Up:
		return turn
	case Down:
		switch turn {
		case Left:
			return Right
		case Right:
			return Left
		}
	case Left:
		switch turn {
		case Left:
			return Down
		case Right:
			return Up
		}
	case Right:
		switch turn {
		case Left:
			return Up
		case Right:
			return Down
		}
	}

	panic("not handled")
}

func (d Direction) Flip() Direction {
	switch d {
	case Up:
		return Down
	case Down:
		return Up
	case Left:
		return Right
	case Right:
		return Left
	}

	panic("not handled")
}

func (d Direction) String() string {
	switch d {
	case Up:
		return "up"
	case Down:
		return "down"
	case Left:
		return "left"
	case Right:
		return "right"
	case UpLeft:
		return "up left"
	case UpRight:
		return "up right"
	case DownLeft:
		return "down left"
	case DownRight:
		return "down right"
	case DontMove:
		return "don't move"
	default:
		panic(fmt.Sprintf("unknown direction %d", d))
	}
}

type Position struct {
	Row int
	Col int
}

func NewPosition(row int, col int) Position {
	return Position{Row: row, Col: col}
}

func PosInSlice(a Position, list []Position) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (p Position) Move(direction Direction, moves int) Position {
	switch direction {
	case Up:
		return p.Slide(-moves, 0)
	case Down:
		return p.Slide(moves, 0)
	case Left:
		return p.Slide(0, -moves)
	case Right:
		return p.Slide(0, moves)
	case UpLeft:
		return p.Slide(-moves, -moves)
	case UpRight:
		return p.Slide(-moves, moves)
	case DownLeft:
		return p.Slide(moves, -moves)
	case DownRight:
		return p.Slide(moves, moves)
	case DontMove:
		return p.Slide(0, 0)
	}

	panic("not handled")
}

// Slide returns a new position from a row diff and col diff.
func (p Position) Slide(row, col int) Position {
	return Position{
		Row: p.Row + row,
		Col: p.Col + col,
	}
}

func (p Position) String() string {
	return fmt.Sprintf("row=%d, col=%d", p.Row, p.Col)
}

type Location struct {
	Pos Position
	Dir Direction
}

func NewLocation(row int, col int, dir Direction) Location {
	return Location{
		Pos: NewPosition(row, col),
		Dir: dir,
	}
}

func (l Location) Straight(moves int) Location {
	dir := l.Dir
	pos := l.Pos.Move(dir, moves)
	return Location{Pos: pos, Dir: dir}
}

func (l Location) Turn(d Direction, moves int) Location {
	dir := l.Dir.Turn(d)
	pos := l.Pos.Move(dir, moves)
	return Location{Pos: pos, Dir: dir}
}

// Manhattan returns the manhattan distance.
func (p Position) Manhattan(p2 Position) int {
	return Abs(p.Row-p2.Row) + Abs(p.Col-p2.Col)
}

// ManhattanZero returns the manhattan distance from the zero position.
func (p Position) ManhattanZero() int {
	return p.Manhattan(Position{})
}

// A Board is a map of a position to a string.
type Board map[Position]string

func NewBoard() Board {
	return Board{}
}

// Wrap returns the position when stepping outside of the board.
// Pass in the position before wrapping, and the direction, it takes 1 step
func (b Board) Wrap(d Direction, p Position) Position {
	// If the position doesn't exist, return the position
	if _, exists := b[p]; !exists {
		return p
	}

	// Move the other way as far as possible
	for {
		new_pos := p.Move(d.Flip(), 1)
		if _, exists := b[new_pos]; !exists {
			return p
		}
		p = new_pos
	}
}

func (b Board) Bottom() int {
	max := 0
	for p := range b {
		if p.Row > max {
			max = p.Row
		}
	}
	return max
}

func (b Board) Right() int {
	max := 0
	for p := range b {
		if p.Col > max {
			max = p.Col
		}
	}
	return max
}
