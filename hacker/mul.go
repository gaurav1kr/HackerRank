package main
import (
"fmt"
"sync"
)

var wg sync.WaitGroup

func mul(c chan int , val int){
	c <- val*5
	wg.Done()
	
}
func main(){
	 c := make(chan int)
	 cnt := 0
	 for cnt < 10 {
		go mul(c , cnt)
		wg.Add(1)
		cnt++
	 }
	 
	 ct := 0
	 for ct < 10 {
			if ct > 2{
			fmt.Println(<-c)
			} else{
			fmt.Errorf("%v",<-c)
			}
		ct++
	 }
}