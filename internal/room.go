package internal

type Room struct {
	Id          int
	Name        string
	Description string
	HasEnemies  bool
	Exits       map[string]*Room
}
