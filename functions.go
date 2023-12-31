package main
import "fmt"
func main(){
    printMe("hello")
}


func printMe(name string) string {
	fmt.Println(`{name} World`)
}