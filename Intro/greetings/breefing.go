package greetings

import (
	"fmt"
	"os"
	"time"
)

const logsFile = "logs.txt"

var lastSessionTime time.Time

func Greetings() {
	if !isNewSession() {
		// If it's not a new session, print the session information again
		printLastSession()
		return
	}

	// It's a new session, proceed with the normal greeting
	var amORpm, greet string
	now := time.Now()
	formattedTime := now.Format("3:04:05 PM")

	if now.Hour() < 12 {
		amORpm = "am"
	} else {
		amORpm = "pm"
	}

	if amORpm == "am" {
		greet = "Good Morning"
	} else {
		greet = "Good Evening"
	}

	// Print the greeting message along with the session information
	fmt.Println(greet, "BOSS!\nWelcome to your Golang bootcamp session", now.Day(), now.Month(), now.Year(), ",", formattedTime)

	// Update the last session time
	writeLastSessionTime(now)
}

//==================== Sessions Control functions ==========================

// -------------------- New Session ---------------------------
func isNewSession() bool {
	lastSessionTime, err := readLastSessionTime()
	if err != nil {
		return true // If there is an error read last session time as New Session
	}
	// calculate the difference between timestamps
	diff := time.Since(lastSessionTime)

	//Check if the difference is greater than or equal to 24 hours
	return diff >= 24*time.Hour
}

// ------------------ Read Last Session ----------------------------
func readLastSessionTime() (time.Time, error) {
	file, err := os.Open(logsFile)
	if err != nil {
		return time.Time{}, err
	}
	defer file.Close()

	var lastSessionTimeStr string
	_, err = fmt.Fscanf(file, "%s", &lastSessionTimeStr)
	if err != nil {
		return time.Time{}, err
	}

	// Parse the time string into a time.Time object
	lastSessionTime, err := time.Parse(time.RFC3339, lastSessionTimeStr)
	if err != nil {
		return time.Time{}, err
	}

	// Return the time object
	return lastSessionTime, nil
}

// ------------------ Write Last Session ----------------------------
func writeLastSessionTime(lastSessionTime time.Time) error {
	file, err := os.Create(logsFile)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "%s\n", lastSessionTime.Format(time.RFC3339))
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}

// ----------------- Print Last Session ----------------------
func printLastSession() {
	lastSessionTime, err := readLastSessionTime()
	if err != nil {
		fmt.Println("Error reading last session time:", err)
		return
	}

	formattedTime := lastSessionTime.Format("3:04:05 PM")
	fmt.Println("Last Session Log:", lastSessionTime.Day(), lastSessionTime.Month(), lastSessionTime.Year(), ":", formattedTime)
	fmt.Println("Welcome back Boss!")
}

// package greetings

// import (
// 	"fmt"
// 	"time"
// )

// // greet the user depending about time frames 				=> customGreet() []
// // display the Date and time of his session   				=> nowDateTime() []
// // sessions logs to print only the first session per day	=> writeLogs() 	 []
// //															=> readLogs()	 []
// //															=> compareLogs() []
// //															=> printlogs()   []

// func Greetings() {
// 	fmt.Println("========================== OlympiadTech ===========================")
// 	customGreet()
// }

// func customGreet(dayTime bool) {

// 	var welcome, morning, evening = "Welcome to your Golang bootcamp.", "Good Morning Sir!", "Good Evening Sir!"

// 	if dayTime == true {
// 		fmt.Print(morning,"\n", welcome)
// 	} else {
// 		fmt.Println(evening,"\n", welcome)
// 	}
// }

// func nowDateTime() {
// 	fmt.Println(":", time.Now().Format("02-01-06"), "Oujda City")
// }

// func writeLogs() {
// 	const logFile = "logs.txt"
// }

// func readLogs() {

// }

// func isNewLog() bool {

// }

// func printLogs() {

// }
