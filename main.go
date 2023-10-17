package main

import (
	"fmt"
	"os"
	"strconv"
	"text-adventure/internal"
	"time"

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
		input := handleDecision(state)
		switch input {
		case "s":
		case "state":
			handleState(state)
		case "r":
		case "room":
			handleRoom(state)
		case "a":
		case "atk":
		case "attack":
			handleAttack(state.Player, state.Enemy)
		case "m":
		case "move":
		case "mv":
			handleMove(state)
		case "quit":
			handleQuit()
		default:
			handleError()
		}
	}
}

func handleError() {
	internal.Error("Invalid command. Try again.")
}

func handleQuit() {
	internal.Success("Thanks for playing!")
	os.Exit(0)
}

func handleDecision(state internal.State) string {
	internal.Important(state.CurrentRoom.Description)

	if state.CurrentRoom.HasEnemies {
		internal.Error("This room has enemies!")
		fmt.Printf("\n%s's Health: %d\n", state.Player.Name, state.Player.Health)
		fmt.Println("What will you do? (move/attack/room/state)")

	} else {
		internal.Important("What will you do? (move/room/state)")
	}
	return internal.ProcessInput()
}

func handleMove(state internal.State) {
	internal.Important("In which direction do you want to go?")
	showExits(state)
	direction := internal.ProcessInput()

	if _, ok := state.CurrentRoom.Exits[direction]; ok {
		moveRoom(state)

	} else {

		internal.Error("That's not a valid direction.")
	}
}

func handleRoom(state internal.State) {
	internal.ClearScreen()
	internal.Important(state.CurrentRoom.Description)
	internal.Info("This room has the following exits:")
	showExits(state)
}

func handleState(state internal.State) {
	internal.ClearScreen()
	internal.Info(state.Player.Name)
	internal.Info("Cleared Rooms: " + strconv.Itoa(state.ClearedRooms))
}

func showExits(state internal.State) {
	internal.Info(("exits:"))
	for exit, room := range state.CurrentRoom.Exits {
		fmt.Printf("- %s to %s\n", exit, room.Name)
	}
}

func moveRoom(state internal.State) {
	internal.ClearScreen()
	internal.Success("Moving")
	time.Sleep(100 * time.Millisecond)
	internal.Success(".")
	time.Sleep(100 * time.Millisecond)
	internal.Success(".")
	time.Sleep(100 * time.Millisecond)
	internal.Success(".")
	state.MoveRooms(state.Player)
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
