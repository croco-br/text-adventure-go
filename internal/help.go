package internal

import (
	"fmt"
)

func display() {
	fmt.Println("-------------------------------------------------------------")
	fmt.Printf("%s is a text-based RPG.\n", GAME_NAME)
	fmt.Println("You play by inputting text commands, commanding your character.")
	fmt.Println("----------------------- HOW TO PLAY: Commands ----------------")
	fmt.Println("start: begin the game.")
	fmt.Println("quit: exit the game.")
	fmt.Println("status: display your character's current status.")
	fmt.Println("move <direction> <quantity>: move your character in a specific direction for a quantity of tiles.")
	fmt.Println("The available directions are: up, down, left, right. Example: move up 1, move left 2")
	fmt.Println("spell <spell name>: use a spell. Example: spell heal")
	fmt.Println("attack <enemy name>: attack an enemy. Example: attack zombie")
	fmt.Println("-------------------------------------------------------------")
}
