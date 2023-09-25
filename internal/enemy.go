package internal

// Player represents the player character.
type Enemy struct {
	Name    string
	Health  int
	Attack  int
	Defense int
}

func NewEnemy(name string, health int, attack int, defense int) *Enemy {
	return &Enemy{
		Name:    name,
		Health:  health,
		Attack:  attack,
		Defense: defense,
	}
}
