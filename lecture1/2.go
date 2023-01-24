package main
import "fmt"
func main(){
	var num = 20
	while num > 0{
		if num % 2 == 0{
			fmt.Println(num, "- it is even variable"+"\n")
			num = num - 1;
		}
	}

}