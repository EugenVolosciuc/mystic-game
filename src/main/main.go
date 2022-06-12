package main

import (
	"bufio"
	"fmt"
	"mystic/src/utils"
	"os"
	"strconv"
)

type Name string
type GameStep byte

type Attack struct {
	name Name
	cost int
	damage int
}

type Class struct {
	name Name
	health int
	maxHealth int
	energy int
	maxEnergy int
	attacks []Attack
}

type Player struct {
	name Name
	class Class
}

func printSpacing() {
	fmt.Println("-----")
}

func getFreestyleInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return scanner.Text()
}

func getSpecificInput(inputChoices []string) string {
	var input string
	var choicesInfoLine string
	defaultInputChoices := []string{
		"Help",
		"Character info",
	}

	allInputChoices := append(inputChoices, defaultInputChoices...)

	for i, choice := range allInputChoices {
		choicesInfoLine += fmt.Sprintf("[%v] %v  ", i, choice)
	}

	fmt.Println(choicesInfoLine)

	validInput := utils.StringInSlice(input, allInputChoices)

	for !validInput {
		input = getFreestyleInput()

		if len(input) != 0 {
			break
		}

		fmt.Println("Please choose a valid input")
	}

	return input
}

func (player *Player) setupPlayer(classes []Class) {
	classChoices := []string{}

	for _, class := range classes {
		classChoices = append(classChoices, string(class.name))
	}

	fmt.Println("Player, what is your name? ")

	for len(player.name) == 0 {
		player.name = Name(getFreestyleInput())
		
		if len(player.name) != 0 {
			break
		}

		fmt.Println("Please write your character's name")
	}

	fmt.Println("Alright, now please choose your character class: ")
	classIndex, _ := strconv.ParseInt(getSpecificInput(classChoices), 0, 8)
	chosenClass := classes[classIndex]
	
	fmt.Printf("Welcome %v, the %v!\n", player.name, chosenClass.name)
	printSpacing()
}

type Game struct {
	players [2]Player
	gameStep GameStep
}

func (game *Game) advanceGame() {
	game.gameStep += 1
}

func (game *Game) displayIntro() {
	fmt.Println("Welcome to Mystic!")
	printSpacing()
	game.advanceGame()
}

func (game *Game) setupPlayers(classes []Class) {
	for i := range game.players {
		game.players[i].setupPlayer(classes)
	}
}

func initialeClasses() []Class {
	return []Class{
		{
			name: "Knight",
			health: 50,
			maxHealth: 50,
			energy: 50,
			maxEnergy: 50,
			attacks: []Attack{
				{
					name: "Kick",
					cost: 5,
					damage: 5,
				},
				{
					name: "Slash",
					cost: 5,
					damage: 5,
				},
				{
					name: "Dash",
					cost: 5,
					damage: 5,
				},
				{
					name: "Blitz",
					cost: 5,
					damage: 5,
				},
			},
		},
		{
			name: "Archer",
			health: 50,
			maxHealth: 50,
			energy: 50,
			maxEnergy: 50,
			attacks: []Attack{
				{
					name: "Snipe",
					cost: 5,
					damage: 5,
				},
				{
					name: "Electric arrow",
					cost: 5,
					damage: 5,
				},
				{
					name: "Triple shot",
					cost: 5,
					damage: 5,
				},
				{
					name: "Arrow shower",
					cost: 5,
					damage: 5,
				},
			},
		},
	}
}

func main() {
	const (
		Intro GameStep = iota
		PlayersSetup
		Playing
		GameStop
	)

	classes := initialeClasses()

	fmt.Println(classes)

	game := Game{}
	game.displayIntro()
	game.setupPlayers(classes)
}