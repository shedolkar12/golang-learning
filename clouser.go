package main
import "fmt"

func main(){
	nextGen := evenGenrator()
	fmt.Println(nextGen())
	fmt.Println(nextGen())
	fmt.Println(nextGen())
	fmt.Println(nextGen())

	// nextEven := makeEvenGenerator() 
	// fmt.Println(nextEven())
	// fmt.Println(nextEven())

}

func evenGenrator() (func() uint){
	i := uint(0)
	return func() (ret uint){
		ret = i
		i += 2
		return ret
	}
}
