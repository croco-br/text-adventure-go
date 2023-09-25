package internal

type State struct {
	Player      Player
	Enemy       Enemy
	CurrentRoom *Room
}

func NewState(player Player, enemy Enemy) *State {
	return &State{
		Player: player,
		Enemy:  enemy,
	}
}
