// helloworld.go
// https://www.linkedin.com/learning/go-for-python-developers
package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

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
}

// a string enum
type Role string

const (
	Viewer    Role = "viewer"
	Developer Role = "developer"
	Admin     Role = "admin"
)

type User struct {
	Login string
	Role  Role
}

func Promote(user *User, role Role) {
	user.Role = role
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
