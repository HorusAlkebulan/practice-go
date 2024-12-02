// helloworld.go
// https://www.linkedin.com/learning/go-for-python-developers
package main

import (
	"fmt"
	"sort"
	"time"
	"unicode/utf8"
)

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

func Promote(u *User, r Role) {
	u.Role = r
}

func main() {
	msg := "Hello Gophers"
	fmt.Println(msg)
	showTime()
	averageDivisibleBy3or5()
	printBanner()
	getMedian()

	u := User{
		"elliot",
		Viewer,
	}
	fmt.Printf("Promoting user: name=%s, role=%s\n", u.Login, u.Role)
	Promote(&u, Admin)
	fmt.Printf("user: name=%s, role=%s\n", u.Login, u.Role)
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
