package piece

// type Color represents the chess piece colors
type Color int

// White and Black are enums representing chess piece colors
const (
	None Color = iota
	White
	Black
)

// ColorString returns the corresponding color as a string
func (color Color) String() string {
	switch color {
	case White:
		return "white"
	case Black:
		return "black"
	default:
		return ""
	}
}
