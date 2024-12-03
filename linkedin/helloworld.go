// helloworld.go
// https://www.linkedin.com/learning/go-for-python-developers
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync"
	"time"
	"unicode"
	"unicode/utf8"
)

func fileHead(fileName string, size int) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	// same as using 'with' in python
	defer file.Close()

	buf := make([]byte, size)
	fileSize, err := file.Read(buf)

	if err != nil {
		return nil, err
	}
	if fileSize != size {
		return nil, fmt.Errorf("%q file size mismatch", fileName)
	}
	return buf, nil
}

// Role enum
type Role string

const (
	Viewer    Role = "viewer"
	Developer Role = "developer"
	Admin     Role = "admin"
)

// User "class"
type User struct {
	Login string
	Role  Role
}

// Our first use of passing by ref!
func Promote(u *User, r Role) {
	u.Role = r
}

// NOTE: Fixes issue with fatal error: all goroutines are asleep - deadlock!
var waitGroup sync.WaitGroup

func main() {
	msg := "Hello Gophers"
	fmt.Println(msg)
	showTime()
	averageDivisibleBy3or5()
	printBanner()
	getMedian()
	countOfWord()
	charFrequency()

	n := 4
	result := collatzStep(n)
	fmt.Printf("collatzStep(%d): %d\n", n, result)

	n = 5
	result = collatzStep(n)
	fmt.Printf("collatzStep(%d): %d\n", n, result)

	filename := "app.go"
	root, ext := SplitExt(filename)
	fmt.Printf("%s -> %s, %s\n", filename, root, ext)
	fmt.Printf("%s -> %#v, %#v\n", filename, root, ext)

	values := []float64{2, 4, 8}
	mean, err := Mean(values)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fmt.Print("values: ")
	fmt.Println(values)
	fmt.Printf("mean: %.2f\n", mean)

	user := User{"Horus", Viewer}
	fmt.Printf("User %s -> %s\n", user.Login, user.Role)
	fmt.Printf("Promoting user %s\n", user.Login)
	Promote(&user, Admin)
	fmt.Printf("User %s -> %s\n", user.Login, user.Role)

	data, err := fileHead("head.png", 8)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File bytes: ", data)
	}

	loc, err := NewLocation(32.5253837, 34.9427434)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(loc)

	latitude := 45.5678901
	longitude := -23.456780

	fmt.Printf("Moving to %f, %f\n", latitude, longitude)
	loc.Move(latitude, longitude)
	fmt.Println("New location:", loc)

	name := "2021 BMW X7"
	fmt.Println("Creating new car")
	car, err := NewCar(name, latitude, longitude)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Car status: %#v\n", car)
	// NOTE: Why does this work? Shouldn't it be something like car.location.Move()?
	car.Move(0.0, 0.0)
	fmt.Printf("Car status: %#v\n", car)

	items := []Moveable{
		&Location{-1.0, -2.0},
		&Car{
			ID: "Volvo XC60",
			Location: Location{
				Latitude:  -3.0,
				Longitude: -4.0,
			},
		},
	}
	fmt.Println("Before move of items")
	for _, item := range items {
		fmt.Printf("- %#v", item)
	}
	moveAll(items, 9.0, 10.0)
	fmt.Println("After move of items")
	for _, item := range items {
		fmt.Printf("- %#v", item)
	}

	usingGenerics()

	// NOTE: Looks like you can't do both of these at the same time. You get a channels block.
	// time.Sleep(time.Duration(5) * time.Second)

	fmt.Println("Using Go routines")
	for i := 0; i < 5; i++ {
		waitGroup.Add(1)
		go worker(i)
	}

	fmt.Println("Resuming main, waiting to allow threads to complete...")
	waitGroup.Wait()
	// time.Sleep(time.Duration(5) * time.Second)

	fmt.Println("Using channels")
	// NOTE: Channels are blocking by default
	channels := make(chan int)
	go func() {
		channels <- 99 // send using a go routine to avoid deadlock
	}()

	val := <-channels // receive
	fmt.Printf("Received %d from the channel\n", val)

	demoWaitGroup()

	url := "https://www.linkedin.com/learning/go-for-python-developers"
	timeout := 5 * time.Millisecond
	res := CheckURL(url, timeout)
	fmt.Printf("CheckURL response with timeout %v: %t\n", timeout, res)
	timeout = 5 * time.Second
	res = CheckURL(url, timeout)
	fmt.Printf("CheckURL response with timeout %v: %t\n", timeout, res)

}

func CheckURL(url string, timeout time.Duration) bool {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return false
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false
	}

	if resp.StatusCode != http.StatusOK {
		return false
	}

	return true
}
func usingGenerics() {
	fmt.Println("Using generics")
	fmt.Printf("%d + %d = %d\n", 1, 2, Add(1, 2))
	fmt.Printf("%f + %f = %f\n", 1.01, 2.02, Add(1.01, 2.02))
	fmt.Printf("%s + %s = %s\n", "Horus", "Alkebu-Lan", Add("Horus ", "Alkebu-Lan"))
}

func demoWaitGroup() {
	var wg sync.WaitGroup
	ch := make(chan string)

	for i := 0; i < 3; i++ {
		go func(id int) {
			for msg := range ch {
				time.Sleep(time.Duration(100) * time.Millisecond)
				fmt.Printf("demoWaitGroup(): %d finished %s\n", id, msg)
				wg.Done()
			}
		}(i)
	}

	for _, msg := range []string{"A", "B", "C", "D", "E", "F"} {
		wg.Add(1)
		ch <- msg
	}
	wg.Wait()
	fmt.Println("All jobs complete")
}

func worker(n int) {
	defer waitGroup.Done()

	ms := 100.0
	msd := time.Duration(ms)
	time.Sleep(msd * time.Millisecond)
	fmt.Printf("Using duration %f, running go routine worker %d\n", ms, n)
	// fmt.Printf("Using wait group, running go routine worker %d\n", n)
}

type Addable interface {
	int | float64 | string
}

// Generics
// without interface: func Add[T int | float64 | string](a, b T) T {
func Add[T Addable](a, b T) T {
	return a + b
}

type Moveable interface {
	Move(float64, float64)
}

func moveAll(items []Moveable, lat float64, lng float64) {
	for _, item := range items {
		item.Move(lat, lng)
	}
}

type Car struct {
	ID string
	Location
}

func NewCar(id string, latitude float64, longitude float64) (Car, error) {
	loc, err := NewLocation(latitude, longitude)
	if err != nil {
		return Car{}, err
	}
	car := Car{
		ID:       id,
		Location: loc,
	}
	return car, nil
}

// effectively Location class in python
type Location struct {
	Latitude  float64
	Longitude float64
}

func NewLocation(latitude, longitude float64) (Location, error) {
	if latitude < -90 || latitude > 90 {
		return Location{}, fmt.Errorf("Invalid latitude: %#v", latitude)
	}
	if longitude < -180 || longitude > 180 {
		return Location{}, fmt.Errorf("Invalid longitude: %#v", longitude)
	}
	loc := Location{
		Latitude:  latitude,
		Longitude: longitude,
	}
	return loc, nil
}

func (l *Location) Move(latitude, longitude float64) {
	l.Latitude = latitude
	l.Longitude = longitude
}

func Sum(values []float64) float64 {
	sum := 0.0
	for _, value := range values {
		sum = sum + value
	}
	return sum
}

func Mean(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0.0, fmt.Errorf("Mean of an empty slice")
	}

	sum := Sum(values)
	count := float64(len(values))
	mean := sum / count
	return mean, nil
}

func SplitExt(path string) (string, string) {
	i := strings.LastIndex(path, ".")
	if i == -1 {
		return path, ""
	}
	return path[:i], path[i:]
}

func collatzStep(n int) int {
	if n%2 == 0 {
		return n / 2
	}
	return n*3 + 1
}

func charFrequency() {
	var poem = `
	those who do not feel this love
	pulling them like a river
	those who do not drink dawn
	like a cup of spring water
	or take in sunset like supper
	those who do not want to change
	let them sleep
	`
	fmt.Print("Poem: ")
	fmt.Println(poem)

	counts := make(map[rune]int)
	globalCount := 0.0 // we will use in float calculation

	for _, c := range poem {
		if unicode.IsSpace(c) {
			continue
		}
		counts[c]++
		globalCount++
	}

	var chars []rune // rune = python char
	for c := range counts {
		chars = append(chars, c)
	}
	sort.Slice(chars, func(i, j int) bool {
		c1, c2 := chars[i], chars[j]
		return counts[c1] > counts[c2] // gives us reverse order
	})

	// report final results
	for _, c := range chars {
		n := counts[c]
		f := float64(n) / globalCount * 100
		fmt.Printf("%c: %.2f\n", c, f)
	}
}

func showTime() {
	nowtime := time.Now()
	fmt.Println("The time is now ")
	fmt.Println(nowtime)
}

func averageDivisibleBy3or5() {
	count, total := 0, 0
	for n := 1; n <= 100; n++ {
		if n%3 == 0 || n%5 == 0 {
			count++
			total += n
		}
	}

	fmt.Println("Average:")
	fmt.Println(float64(total) / float64(count))
}

func printBanner() {
	message := "Let's Goooooo! ❤️" // 16 chars
	width := 32
	messageLen := len(message)
	paddingLen := (width - messageLen) / 2 // note: this is integer division
	fmt.Println("Padding:")
	fmt.Println(paddingLen)
	fmt.Println("Length of Message:")
	fmt.Println(messageLen)
	for n := 0; n < width; n++ {
		fmt.Print("-")
	}
	fmt.Print("\n")
	for n := 0; n < paddingLen; n++ {
		fmt.Print(" ")
	}
	fmt.Print(message)
	fmt.Print("\n")
	for n := 0; n < width; n++ {
		fmt.Print("-")
	}
	fmt.Print("\n")

	fmt.Println("NOTE: The len() func gives difference result. It returns length in bytes, so when using unicode use RuneCount")

	messageRuneCount := utf8.RuneCountInString(message)
	paddingRuneCount := (width - messageRuneCount) / 2
	fmt.Println("Padding:")
	fmt.Println(paddingRuneCount)
	fmt.Println("Length of Message:")
	fmt.Println(messageRuneCount)
	for n := 0; n < width; n++ {
		fmt.Print("-")
	}
	fmt.Print("\n")
	for n := 0; n < paddingRuneCount; n++ {
		fmt.Print(" ")
	}
	fmt.Print(message)
	fmt.Print("\n")
	for n := 0; n < width; n++ {
		fmt.Print("-")
	}
	fmt.Print("\n")
}

func getMedian() {
	nums := []float64{2, 1, 3, 0, 6, 8, 9}
	fmt.Print("nums: ")
	fmt.Println(nums)

	fmt.Print("Sorted nums: ")
	sort.Float64s(nums)
	fmt.Println(nums)

	var median float64
	i := len(nums) / 2
	if len(nums)%2 == 1 {
		median = nums[i]
	} else {
		median = (nums[i-1] + nums[i]) / 2
	}
	fmt.Println("Median:")
	fmt.Println(median)
}

func countOfWord() {

	var poem = `
	those who do not feel this love
	pulling them like a river
	those who do not drink dawn
	like a cup of spring water
	or take in sunset like supper
	those who do not want to change
	let them sleep
	`

	frequency := make(map[string]int)
	for _, word := range strings.Fields(poem) {
		frequency[word]++
	}

	maxW, maxC := "", 0
	for w, c := range frequency { // range is similar to enumerate
		if c > maxC {
			maxC = c
			maxW = w
		}
	}
	fmt.Print("Poem: ")
	fmt.Println(poem)
	fmt.Print("maxW: ")
	fmt.Println(maxW)
	fmt.Print("maxC: ")
	fmt.Println(maxC)
	fmt.Print("Frequency: ")
	fmt.Println(frequency)
}
