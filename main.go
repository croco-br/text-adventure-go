package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text-adventure/internal"
)

// ProcessInput processes player input.
func ProcessInput() string {
	var input string
	fmt.Print("\n> ")
	fmt.Scanln(&input)
	return strings.ToLower(input)
}

// ClearScreen clears the terminal screen.
func ClearScreen() {
	cmd := exec.Command("clear") // for Unix-like systems
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	course := internal.NewCourse(10)
	player := internal.Player{
		Name:    "Player",
		Health:  100,
		Attack:  10,
		Defense: 5,
	}

	rooms := course.GetRooms()
	player.CurrentRoom = rooms[0]

	fmt.Println(internal.GAME_NAME)
	fmt.Println("Welcome to the Text Adventure Game!")

	for {
		fmt.Printf("\n%s's Health: %d\n", player.Name, player.Health)
		fmt.Println(player.CurrentRoom.Description)
		fmt.Println("Exits:")
		for exit, room := range player.CurrentRoom.Exits {
			fmt.Printf("- %s to %s\n", exit, room.Name)
		}
		fmt.Println("What will you do? (move/quit)")

		input := ProcessInput()

		switch input {
		case "move":
			fmt.Print("Where do you want to go? ")
			dest := ProcessInput()
			if room, ok := player.CurrentRoom.Exits[dest]; ok {
				player.CurrentRoom = room
				ClearScreen()
			} else {
				fmt.Println("That's not a valid direction.")
			}
		case "quit":
			fmt.Println("Thanks for playing!")
			os.Exit(0)
		default:
			fmt.Println("Invalid command. Try again.")
		}
	}
}
