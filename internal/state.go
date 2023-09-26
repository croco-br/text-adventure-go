package internal

type State struct {
	Player      Player
	Enemy       Enemy
	CurrentRoom *Room
	Rooms       []*Room
}

func NewState() *State {
	return &State{}
}

func (s *State) Init() State {
	c := NewCourse(10)
	p := Player{
		Name:    "Player",
		Health:  100,
		Attack:  10,
		Defense: 5,
	}

	e := Enemy{
		Name:    "Goblin",
		Health:  80,
		Attack:  5,
		Defense: 5,
	}

	s.Player = p
	s.Enemy = e

	s.Rooms = c.GetRooms()
	s.CurrentRoom = s.Rooms[0]

	return *s
}

func (s *State) MoveRooms(player Player) *State {
	oldRoom := s.CurrentRoom
	s.CurrentRoom = s.Rooms[oldRoom.Id+1]
	return s
}
