package main
// Basically dictionaries in python
import "fmt"

func main(){
	grades:= make(map[string]float32)//If you want to have multiple values for one key then use a struct basically defining your type
	grades["Timmy"]= 42
	grades["Jess"] = 54
	grades["Sam"] = 76

	fmt.Println(grades)

	TimsGrade := grades["Timmy"]
	fmt.Println(TimsGrade)

	for k,v := range grades{
		fmt.Println(k,":",v)
	}
	delete(grades, "Timmy")
}

//Note: For linting issues you can use "gofmt `script-name`"