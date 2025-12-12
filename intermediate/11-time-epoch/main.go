package main

import (
	"fmt"
	"strconv"
	"time"
)

// --- Define Standard Layouts for Parsing ---
// The standard Go time format constants are often used, but custom ones rely on the reference time.
const CustomLogLayout = "2006-01-02 15:04:05"
const CustomDateLayout = "01/02/2006"

func main() {
	// --- 1. Getting the Current Time ---
	// time.Now() always returns a Time struct, usually in the local timezone of the machine.
	now := time.Now()
	fmt.Println("--- 1. Current Time ---")
	fmt.Printf("Current Time (Local): %v\n", now)
	fmt.Printf("Current Time (UTC): %v\n", now.UTC())

	// --- 2. Formatting Time (Time -> String) ---
	fmt.Println("\n--- 2. Formatting Time ---")

	// Use the custom defined layout based on the reference time:
	formattedLog := now.Format(CustomLogLayout)
	fmt.Printf("Formatted for Logging (%s): %s\n", CustomLogLayout, formattedLog)

	// Use the standard layout constants provided by the time package:
	formattedISO := now.Format(time.RFC3339)
	fmt.Printf("Formatted ISO 8601 (API Standard): %s\n", formattedISO)

	// --- 3. Parsing Time (String -> Time) ---
	fmt.Println("\n--- 3. Parsing Time ---")

	timeString := "2025-12-09 10:30:00"

	// Parse the string using the EXACT same layout string used to format it.
	parsedTime, err := time.Parse(CustomLogLayout, timeString)
	if err != nil {
		fmt.Printf("Error parsing time: %v\n", err)
		return
	}
	fmt.Printf("Parsed Time Object: %v\n", parsedTime)

	// --- 4. Working with Epoch/Unix Time ---
	fmt.Println("\n--- 4. Epoch Time (Seconds since Jan 1, 1970 UTC) ---")

	// Convert a Time object to Epoch seconds (Unix time)
	epochSeconds := now.Unix()
	fmt.Printf("Current Epoch Seconds: %d\n", epochSeconds)

	// Convert a Time object to Epoch milliseconds (useful for JavaScript/APIs)
	epochMilliseconds := now.UnixMilli()
	fmt.Printf("Current Epoch Milliseconds: %d\n", epochMilliseconds)

	// --- 5. Converting Epoch back to Time ---
	// Create a Time object from Epoch seconds
	targetEpochStr := "1672531200" // Corresponds to Jan 1, 2023 00:00:00 UTC

	// Convert the string to int64 first
	targetEpoch, err := strconv.ParseInt(targetEpochStr, 10, 64)
	if err != nil {
		fmt.Printf("Error parsing epoch string: %v\n", err)
		return
	}

	// time.Unix(seconds, nanoseconds) creates the Time object
	epochToTime := time.Unix(targetEpoch, 0)
	fmt.Printf("Epoch %s converts to: %v\n", targetEpochStr, epochToTime.UTC())

	// --- 6. Handling Time Zones (Crucial for APIs) ---
	fmt.Println("\n--- 6. Time Zone Handling ---")

	// Load a specific location (e.g., for user settings or database storage)
	parisLocation, err := time.LoadLocation("Europe/Paris")
	if err != nil {
		// Time zone loading might fail if the OS lacks zoneinfo data
		fmt.Println("Could not load Paris location:", err)
		return
	}

	// Example: Display the current time in the Paris timezone
	parisTime := now.In(parisLocation)
	fmt.Printf("Time in Paris: %v\n", parisTime.Format(time.RFC822))

	// Converting a Parsed Time to a Specific Time Zone
	// Assume the parsedTime from step 3 (2025-12-09 10:30:00) was intended to be in UTC
	parsedTimeUTC := parsedTime.In(time.UTC)
	fmt.Printf("Parsed Time (as UTC): %v\n", parsedTimeUTC)

	// Convert that UTC time to Paris local time
	parisParsedTime := parsedTimeUTC.In(parisLocation)
	fmt.Printf("Time in Paris (via UTC conversion): %v\n", parisParsedTime)
}
