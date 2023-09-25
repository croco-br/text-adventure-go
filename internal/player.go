package internal

// Player represents the player character.
type Player struct {
	Name        string
	Health      int
	Attack      int
	Defense     int
	CurrentRoom *Room
}
