package main

import (
	"fmt"
	"os"
	"text-adventure/internal"

	figure "github.com/common-nighthawk/go-figure"
)

func main() {
	start()
	state := internal.NewState()
	state.Init()
	loop(*state)
}

func loop(state internal.State) {
	for {
		internal.Important(state.CurrentRoom.Description)

		if state.CurrentRoom.HasEnemies {
			internal.Error("This room has enemies!")
			fmt.Printf("\n%s's Health: %d\n", state.Player.Name, state.Player.Health)
			fmt.Println("What will you do? (move/attack/room)")

		} else {
			internal.Important("What will you do? (move/room)")
		}

		input := internal.ProcessInput()

		switch input {
		case "r":
		case "room":
			internal.ClearScreen()
			internal.Important(state.CurrentRoom.Description)
			internal.Info("This room has the following exits:")
			for exit, room := range state.CurrentRoom.Exits {
				fmt.Printf("- %s to %s\n", exit, room.Name)
			}
		case "a":
		case "atk":
		case "attack":
			handleAttack(state.Player, state.Enemy)
		case "m":
		case "move":
		case "mv":
			internal.Important("Where do you want to go? ")
			direction := internal.ProcessInput()

			if _, ok := state.CurrentRoom.Exits[direction]; ok {
				internal.Success("Moving rooms...")
				state.MoveRooms(state.Player)

			} else {

				internal.Error("That's not a valid direction.")
			}
		case "quit":
			internal.Success("Thanks for playing!")
			os.Exit(0)
		default:
			internal.Error("Invalid command. Try again.")
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

func start() {
	internal.ClearScreen()
	logo := figure.NewFigure(internal.GAME_NAME, "", true)
	fmt.Println(logo.String())
	//fmt.Println("Welcome to " + internal.GAME_NAME + ". A text-based RPG.")
}
