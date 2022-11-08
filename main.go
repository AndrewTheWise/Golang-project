package main

import (
    "fmt"
    "os"
    "os/exec"
    "strings"
)

func clearScreen() error {
    c := exec.Command("clear")
    c.Stdout = os.Stdout
    return c.Run()
}

func main() {
    clearScreen()
    fmt.Println("(Password is 1234)\nPassword > ")

    var password string

    fmt.Scanln(&password)

    if password != "1234" {
        fmt.Println("Error: Incorrect Password")
        os.Exit(0)
    }

    for {
        fmt.Println("\nPages:\n◦Profile\n◦Schedule\n◦Agenda")
        fmt.Println("\nType the page you would like to view or type 'exit' to stop the program > ")

        var action string

        fmt.Scanln(&action)

        switch strings.ToLower(action) {
        case "profile":
            clearScreen()
            fmt.Println("Name: Andrew Wise\nAge: 19\nBirthday: September 12th, 2003\nEducation: High school graduate")
        case "schedule":
            clearScreen()
            fmt.Println("8:00 AM: Wake up\n8:30 AM: Eat breakfast\n9:00 AM: Whatever\n12:00 AM: Eat lunch\n1:00 PM: Whatever\n5:00 PM: Eat dinner\n10:00 PM: Go to sleep")
        case "agenda":
            clearScreen()
            fmt.Println("Clean room\nStudy Go\nWalk dog\nMow lawn")
        case "exit":
            clearScreen()
            fmt.Println("Program ended")
            return
        default:
            clearScreen()
            fmt.Println("Error: That is not a valid page")
        }
    }
}



