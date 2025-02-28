package main
import (
	"fmt"
	"strconv"
	"math"
	"strings"
)

var pi float64 = math.Pi

func runDataStructure() {
    var city string = "Kolkata"
    var city_2 string = "Calcutta"
    fmt.Println(city == city_2)
    title := "Hello the world!"
    fmt.Print(title)
    var a, b float64 = 24.4, 3.0
    fmt.Println(a / b)
    fmt.Println(int(a) % int(b))
    grades := [][]int{{10, 20}, {20, 40, 60}, {2, 1, 2, 2, 3, 4}}
    clone :=grades
    clone[0][0] = 10	
    fmt.Println(grades)
    //slice
    arr := []int{-1, -2}
    for _, value := range arr {
	fmt.Println(value)
    }
    arr1 :=[5]int{10, 20, 90, 70, 60}
    slice := arr1[: 3] 
    slice[2] = 900
    new_slice := append(arr1[:3],arr1[3:]...)
    fmt.Println(new_slice)
    // Map
    codes := map[string]string{"en": "English", "fr": "French", "hi": "Hindi"}
    value, found :=codes["en"]
    fmt.Println(found, value, codes)

}

func printDetails(s string) string{
	fmt.Println(s)
    rs := "Hi " + s
	return rs
}
//Complete the missing line (to calculate the factorial of a number)
func fact(n int) int {
	if n == 0 {
		return 1
	}
  	return n * fact(n-1)
}
// Anonymous Functions
var (
	cube = func(i int) string {
		c := i * i * i
		return strconv.Itoa(c)
	}
)

// High Order Functions
//Find the circumference and the area of a circle
func printResult(radius float64, calcFunction func(r float64) float64) {
	result := calcFunction(radius)
	fmt.Println("Result: ", result)
	fmt.Println("Thank you!")
}
func calcArea(r float64) float64 {
	rs :=  pi * r * r
	return rs
	return math.Round(rs)
}
func calcPerimeter(r float64) float64 {
	var rs float64 = pi * 2 * r
	return rs
	return math.Round(rs)
}
func calcDiameter(r float64) float64 {
	return r * 2
}
func getFunction(query int) func(r float64) float64 {
	query_to_func := map[int]func(r float64) float64 {
		1: calcArea,
		2: calcPerimeter,
		3: calcDiameter,
	}
	return query_to_func[query]
}

//Defer statement
func printString(str string) {
	fmt.Printf("%q", str)
}

func printInt(i int) {
	fmt.Printf("%d", i)
}

func printFloat(f float64) {
	fmt.Printf("%.2f", f)
}

func runDeferSta() {
	printString("Browser")
	defer printInt(32)
	defer printFloat(0.24)
	printString("chrome")
	printInt(90)
	defer printFloat(89)
	printInt(900)
}

func runPointer() {
	y :=[3]int{10, 20, 30}
	py := &y
	fmt.Printf("%T %v \n", py, *py)
	fmt.Printf("%v \n", y)
	(*&y)[0] = 100
	fmt.Printf("%v \n", y)
 	
	var y1 int
	var ptr *int = &y1
	*ptr = 0
	*ptr += 5
	fmt.Println(y1)

	s := "Hello"
	var ptr1 *string = &s
	fmt.Println(s)
	*ptr1 += strings.ToUpper(s)
	fmt.Println(s)

}

// ref
func modify(numbers ...int) {
	for i:=range numbers {
		numbers[i] -= 5
	}
}
// none
func modify1(numbers [3]int) {
	for i:=range numbers {
		numbers[i] -= 5
	}
}

//cannot use arr1 (variable of type [3]int) as *[3]int value in argument to modify2
func modify2(numbers *[3]int) {
	for i:=range numbers {
		numbers[i] -= 5
	}
}


func modify3(s *string) {
	*s = strings.ToUpper(*s)
}

func modify4(s map[string]int) {
	s["A"] = 100
}

func runMo4() {
	asii_codes := map[string]int{}
	asii_codes["A"] = 65
	fmt.Println(asii_codes)
	modify4(asii_codes)
	fmt.Println(asii_codes)
}

// passes value by reference
// by default: map, slice
func runDefaultPointer() {
	arr := []int{10, 20, 30}
	fmt.Println(arr)
	modify(arr...)
	fmt.Println(arr)
	//arr1 := [3]int{10, 20, 30}
	//fmt.Println(arr1)
	//modify2(arr1)
	//fmt.Println(arr1)
	s := "hello"
	fmt.Println(s)
	modify3(&s)
	fmt.Println(s)
	runMo4()
}

type Movie struct {
	name string
	rating float32
}
func getMovie(s string, r float32)(m Movie) {
	m = Movie {
		name: s,
		rating: r,
	}
	return m
}

func increaseRating(m *Movie) {
	m.rating += 1.0
}

func runStruct() {
	m := Movie{name: "ABCD"}
	var m2 Movie
	fmt.Printf("%+v", m)
	fmt.Printf("%+v", m2)
	fmt.Printf("%+v", getMovie("xyz", 3.5))
	mov := getMovie("zyz", 2.0)
	// increaseRating(mov) => error
	increaseRating(&mov)
	fmt.Printf("%+v", mov)
	mov1 := getMovie("abc", 3.3)
	movies := make([]Movie, 5)
	movies = append(movies, mov)
	movies = append(movies, mov1)
	for _, value := range movies {
		fmt.Println(value)
	} 
	// condition
	if mov.rating == mov1.rating || mov != mov1 {
		fmt.Println("condition met")
	} else if mov.rating == mov1.rating {
		fmt.Println("condition_2 met")
	}
}

type Student interface {
	getPercentage() int
	getName() string
}

type Postgrad struct {
	name string
	grades []int
}

type Undergrad struct {
	name string
	grades []int
}


func (p Postgrad) getPercentage() int {
	sum := 0
	for _, v := range p.grades {
		sum += v
	}
	return ((sum * 100) / (len(p.grades) * 200))
}

func (p Postgrad) getName() string {
	return p.name
}

func (u Undergrad) getName() string {
	return u.name
}

func (u Undergrad) getPercentage() int {
	sum := 0
	for _, v:= range u.grades {
		sum += v	
	}
	return sum / len(u.grades)
}

func printPercentage(s Student) {
	fmt.Println(s.getPercentage())
}

func printData(s Student) {
	fmt.Println(s.getName())
	fmt.Println(s.getPercentage())
}

func runInterface() {	
	//stu1 Student := Student{}
	//fmt.Printf(stu1)
	grades := []int{90, 75, 80}
	u := Undergrad{"Ross", grades}
	printData(u)
	p := Postgrad{"Joe", []int{150, 190, 185}}
	printData(p)

}

func main() {
	//rs := printDetails("Mr.Tan")
	//fmt.Println(rs)
	//fact10 := fact(10)
	//fmt.Println("the factorial of 10 is  %d", fact10)
 	//x1 := cube(8)
	//fmt.Printf("%T %v \n", x1, x1)
	//High Order Function
	//Find the circumference and the area of a circle whose radius is xx cm
    //var radius float64
	//var query int
	//fmt.Print("Enter the radius of the circle: ")
	//fmt.Scanf("%f", &radius)
	//fmt.Printf("Enter \n 1 - Area \n 2 - Perimeter \n 3 - Diameter: ")
	//fmt.Scanf("%d", &query)
	//printResult(radius,  getFunction(query))
	//Defer
	//runDeferSta()
	//runPointer()
	//runDefaultPointer()
	//runStruct()
	runInterface()
}
