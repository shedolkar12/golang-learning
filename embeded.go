package main
import "fmt"

type Person struct{
	Name string
} 

type Android struct{
	Person
	Model string
}

func main(){
	a := new(Android)
	a.Person.Name = "Rajesh"
	a.Person.talk()
	a.talk()
}

func (p *Person) talk(){
	fmt.Println("I am Talking. and My name is", p.Name)
}