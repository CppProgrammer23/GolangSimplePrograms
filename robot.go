package main

import (
	"fmt"
)

var r Robot

func main() {
	temp := func(x int) int {
		return x - 2
	}
	test(temp)
	r = Robot{Pos{4, 7}, Dir{false, true}}
	display(r)
	setCurrentPos(Pos{6, 1})
	display(r)
	setCurrentDir(Dir{true, false})
	display(r)
	fmt.Print(getCurrentDir(r))
}

type Dir struct {
	toNorth bool
	toWest  bool
}
type Pos struct {
	x int16
	y int16
}

type Robot struct {
	position  Pos
	direction Dir
}

func display(r Robot) {
	fmt.Println("the position of the robot is: ", r.position.x, r.position.y)
	fmt.Println("Is he facing the north pole?")
	if r.direction.toNorth {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
	fmt.Println("Is he facing the west pole?")
	if r.direction.toWest {
		fmt.Print("Yes\n")
	} else {
		fmt.Print("No\n")
	}
}
func getCurrentPos(r Robot) Pos {
	return Pos{r.position.x, r.position.y}
}
func setCurrentPos(p Pos) {
	r.position.x = p.x
	r.position.y = p.y
}
func getCurrentDir(r Robot) Dir {
	return Dir{r.direction.toNorth, r.direction.toWest}
}
func setCurrentDir(d Dir) {
	r.direction.toNorth = d.toNorth
	r.direction.toWest = d.toWest
}
