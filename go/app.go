package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ssh/terminal"
)

var m = make(map[int]string)
var vnum = 1
var snum = false

func main() {
	var in string
	cls()
	color.Cyan("Conrad Terminal")
	color.Yellow("F1: Login")
	color.Yellow("F2: Settings")
	color.Yellow("F3: Exit")
	fmt.Scanln(&in)
	switch in {
	case "f1":
		loginAuth()
	case "f2":
		drawSettings()
	case "f3":
		os.Exit(0)
	}
}
func cls() { os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J") }

//Auth for pass
func loginAuth() {
	cls()
	var password string
	color.Cyan("Password:")
	passwor, err := terminal.ReadPassword(0)
	password = string(passwor)
	b, err := ioutil.ReadFile("pass.txt")
	if err != nil {
		fmt.Print(err)
	}

	hash := string(b)
	match := checkPasswordHash(password, hash)
	if match == true {
		postAuth()
	} else {
		color.Red("Access Denied!")
		os.Exit(0)
	}
}

//Settings
func drawSettings() {
	cls()
	color.Red("Sys Info")
	version := strconv.Itoa(vnum)
	stable := strconv.FormatBool(snum)
	fmt.Println("Version: " + version)
	fmt.Println("Stable?: " + stable)
	time.Sleep(2 * time.Second)
	main()
}

//After authentication
func postAuth() {
	cls()
	var input string
	color.Cyan("Welcome to Conrad Terminal")
	color.Yellow("c. Sellling Calculations")
	color.Yellow("d. Database Functions")
	color.Yellow("lo. Logout")
	fmt.Scanln(&input)
	switch input {
	case "c":
		//Selling algorigthms
		var calcOpt int
		cls()
		color.White("Available Calculations")
		fmt.Println("1. ")
		fmt.Println("2. ")
		fmt.Println("3. ")
		fmt.Println("4. ")
		fmt.Println("5. Return to Main Screen")
		fmt.Scanln(&calcOpt)
		switch calcOpt {
		case 1:
		case 2:
		case 3:
		case 4:
		case 5:
			postAuth()
		}
	case "d":
		var read string
		cls()
		color.Cyan("Input or Output? (i/o)")
		fmt.Scanln(&read)
		switch read {
		case "i":
			inputDatabase()
		case "o":
			readDatabase()
		}
	case "lo":
		os.Exit(0)
	}
}

//Database functionality
func inputDatabase() {
	var place int
	var info string
	fmt.Println("Placement?: ")
	fmt.Scanln(&place)
	fmt.Println("Info to add?: ")
	fmt.Scanln(&info)
	m[place] = info
	time.Sleep(2 * time.Second)
	postAuth()
}
func readDatabase() {
	if len(m) == 0 {
		color.White("Nothing! Restarting in 2 seconds...")
		time.Sleep(2 * time.Second)
		postAuth()
	}
	for i := 0; i <= len(m); i++ {
		fmt.Println(m[i])
	}
	time.Sleep(2 * time.Second)
	postAuth()
}

//Password auth

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
