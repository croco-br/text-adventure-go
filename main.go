package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text-adventure/internal"

	figure "github.com/common-nighthawk/go-figure"
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
	start()
	state := initState()
	loop(state)
}

func loop(state internal.State) {
	for {

		fmt.Printf("\n%s's Health: %d\n", state.Player.Name, state.Player.Health)
		fmt.Println(state.Player.CurrentRoom.Description)
		fmt.Println("Exits:")
		for exit, room := range state.Player.CurrentRoom.Exits {
			fmt.Printf("- %s to %s\n", exit, room.Name)
		}
		fmt.Println("What will you do? (move/quit)")

		input := ProcessInput()

		switch input {
		case "attack":
			handleAttack(state.Player, state.Enemy)
		case "move":
			fmt.Print("Where do you want to go? ")
			dest := ProcessInput()
			if room, ok := state.Player.CurrentRoom.Exits[dest]; ok {
				state.Player.CurrentRoom = room
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

func handleAttack(player internal.Player, enemy internal.Enemy) {
	fmt.Println("You attack the enemy!")
	enemy.Health -= player.Attack
	if enemy.Health <= 0 {
		fmt.Printf("You defeated %s!\n", enemy.Name)

		os.Exit(0)
	}
	player.Health -= enemy.Attack
	if player.Health <= 0 {
		fmt.Printf("%s defeated you!\n", enemy.Name)

		os.Exit(0)
	}
}

func initState() internal.State {
	c := internal.NewCourse(10)
	p := internal.Player{
		Name:    "Player",
		Health:  100,
		Attack:  10,
		Defense: 5,
	}

	e := internal.Enemy{
		Name:    "Goblin",
		Health:  80,
		Attack:  5,
		Defense: 5,
	}

	rooms := c.GetRooms()
	p.CurrentRoom = rooms[0]
	state := internal.NewState(p, e)

	return *state
}

func start() {
	ClearScreen()
	logo := figure.NewFigure(internal.GAME_NAME, "", true)
	fmt.Println(logo.String())
	fmt.Println("Welcome to " + internal.GAME_NAME + ". A text-based RPG.")
}
