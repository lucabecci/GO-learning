package main

import "fmt"

type movement interface {
	up() int
	down() int
	punch() string
	say() string
}

type player struct {
	position int
}

func (p *player) up() int {
	return p.position + 1
}

func (p *player) down() int {
	return p.position - 1
}

func (p *player) punch() string {
	return "punching!!!"
}

func (p *player) say() string {
	return "Hi, welcome to my game!!!"
}

func main() {
	p := player{
		position: 0,
	}

	fmt.Println(p.up())
	fmt.Println(p.down())
	fmt.Println(p.punch())

	fmt.Println(p.say())

}
