/*
Passing a Struct as Function Arguments:
*/

package main

import "fmt"

type BasketballPlayer struct {
	name      string
	age       int
	team      string
	pointsPerGame float64
}

func main() {
	var player1 BasketballPlayer
	var player2 BasketballPlayer

	// Player1 specification
	player1.name = "LeBron James"
	player1.age = 39
	player1.team = "Los Angeles Lakers"
	player1.pointsPerGame = 27.2

	// Player2 specification
	player2.name = "Stephen Curry"
	player2.age = 35
	player2.team = "Golden State Warriors"
	player2.pointsPerGame = 24.6

	// Print player info by calling a function
	printPlayer(player1)
	printPlayer(player2)
}

func printPlayer(player BasketballPlayer) {
	fmt.Println("Name: ", player.name)
	fmt.Println("Age: ", player.age)
	fmt.Println("Team: ", player.team)
	fmt.Println("Points Per Game: ", player.pointsPerGame)
	fmt.Println()
}

/*
Name:  LeBron James
Age:  39
Team:  Los Angeles Lakers
Points Per Game:  27.2

Name:  Stephen Curry
Age:  35
Team:  Golden State Warriors
Points Per Game:  24.6

*/