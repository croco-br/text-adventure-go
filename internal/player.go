package internal

// Player represents the player character.
type Player struct {
	Name    string
	Health  int
	Attack  int
	Defense int
}

func NewPlayer(name string, health int, attack int, defense int) *Player {
	return &Player{
		Name:    name,
		Health:  health,
		Attack:  attack,
		Defense: defense,
	}
}
