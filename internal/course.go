// course.go

package internal

import (
	"math/rand"
)

var directions = []string{
	"north", "south", "east", "west",
	"northeast", "northwest", "southeast", "southwest",
}

var places = []string{
	"a Grand Hall", "inside the Courtyard", "in a Bedroom", "inside a Conference Room",
	"inside a Dinning Room", "near the Kitchen", "inside a Cellar", "Outside",
}

type Room struct {
	Id          int
	Name        string
	Description string
	Exits       map[string]*Room
}

type Course struct {
	Size  int
	Rooms map[string]*Room
}

func NewCourse(size int) *Course {
	return &Course{
		Size: size,
	}
}

func getRandomPlace() string {
	randomIndex := rand.Intn(len(places))
	return places[randomIndex]
}

func getRandomExit() string {
	randomIndex := rand.Intn(len(directions))
	return directions[randomIndex]
}

func (c *Course) GetRooms() []*Room {
	rooms := []*Room{}

	for i := 0; i < c.Size; i++ {
		// Create a new room and add it to the rooms slice
		room := &Room{
			Id:          i,
			Name:        getRandomPlace(),
			Description: "You are " + getRandomPlace() + ".",
			Exits:       make(map[string]*Room),
		}

		if i == 0 {
			room.Exits[getRandomExit()] = room
		} else {
			room.Exits[getRandomExit()] = rooms[i-1]
		}

		rooms = append(rooms, room)
	}

	return rooms
}
