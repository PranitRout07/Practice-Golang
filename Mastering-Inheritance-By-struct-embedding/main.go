package main
import (
	"fmt"
)
type Position struct{
	x float64
	y float64
}
func (p *Position) Move(x,y float64){
	p.x += x
	p.y += y
}
func (p *Position) Teleport(x,y float64){
	p.x = x
	p.y = y 
}
type player struct{
	*Position
}
type enemy struct{
	*Position
}
func newPlayer() *player{
	return &player{
		Position: &Position{},
	}
} 
func newEnemy() *player{
	return &player{
		Position: &Position{},
	}
} 
func main(){
	p := newPlayer()
	p.Move(12,2.89)
	fmt.Println("Moved to position:",p.Position)
	x := newEnemy()
	x.Teleport(300.23,400.289)
	fmt.Println("Enemy teleported to",x.Position)
	p.Teleport(3,8.9)
	fmt.Println("Player now teleported to ",p.Position)
}