package intermediate

import (
	"fmt"
	"time"
)

func TimeOperations() {
	fmt.Println("=== Time Operations Demo ===")
	creationdemo()
	// Epoch time functions
	epochDemo()
	// Component extraction
	componentDemo()
	// Time operations
	operationDemo()
	// Formatting/parsing
	formattingDemo()
	// Time zone operations
	timezoneDemo()

	// Truncation/rounding
	truncationDemo()
	// Duration operations
	durationDemo()

}

func creationdemo() {
	fmt.Println("\n--- Time Creation ---")

	// Current time
	now := time.Now()
	fmt.Println("Current Time is:", now)

	// Create specific time
	specificT := time.Date(1995, time.November, 30, 22, 38, 29, 0, time.UTC)
	fmt.Println("Specific time is:", specificT)

	// From Unix timestamp
	fromUnix := time.Unix(1699574400, 0)
	fmt.Println("From Unix:", fromUnix)

	// Parsed time
	parsed, err := time.Parse("2006-01-02", "2023-11-10")
	if err == nil {
		fmt.Println("Parsed time:", parsed)
	}
}

func epochDemo() {
	fmt.Println("\n--- Epoch Time ---")

	now := time.Now()
	fmt.Println("Now time is:", now)

	// To unix
	sec := now.Unix()
	nsec := now.UnixNano()
	fmt.Printf("Unix time: %d seconds, %d nano seconds\n", sec, nsec)

	// From unix back to normal time
	formUnix := time.Unix(sec, 0)
	fmt.Println("Now time from unix to normal:", formUnix)
}

func componentDemo() {
	fmt.Println("\n--- Time Components ---")

	t := time.Date(2023, time.November, 10, 15, 30, 45, 0, time.UTC)

	fmt.Println("Year:", t.Year())
	fmt.Println("Month:", t.Month())
	fmt.Println("Day:", t.Day())
	fmt.Println("Hour:", t.Hour())
	fmt.Println("Minute:", t.Minute())
	fmt.Println("Second:", t.Second())
	fmt.Println("Weekday:", t.Weekday())
}

func operationDemo() {
	fmt.Println("\n--- Time Operations ---")

	t1 := time.Date(2023, time.November, 10, 15, 0, 0, 0, time.UTC)
	t2 := time.Date(2023, time.November, 10, 16, 0, 0, 0, time.UTC)

	// Addition
	added := t1.Add(2 * time.Hour)
	fmt.Println("Added 2 hours:", added)

	// Subtraction
	subtracted := t1.Add(-30 * time.Minute)
	fmt.Println("Subtracted 30 minutes:", subtracted)

	// Date addition
	dateAdded := t1.AddDate(0, 1, 0)
	fmt.Println("Added 1 month:", dateAdded)

	// Comparison
	fmt.Println("t1 before t2?", t1.Before(t2))
	fmt.Println("t1 after t2?", t1.After(t2))
	fmt.Println("t1 equal t2?", t1.Equal(t2))

	// Difference
	diff := t2.Sub(t1)
	fmt.Println("Duration between t1 and t2:", diff)
}

func formattingDemo() {
	fmt.Println("\n--- Formatting/Parsing ---")

	t := time.Date(2023, time.November, 10, 15, 30, 0, 0, time.UTC)

	// Formatting
	fmt.Println("Formatted:", t.Format("2006-01-02 15:04:05"))

	// Parsing
	parsed, err := time.Parse("2006-01-02", "2023-11-10")
	if err == nil {
		fmt.Println("Parsed:", parsed)
	}

	// RFC3339 format
	fmt.Println("RFC3339:", t.Format(time.RFC3339))
}

func timezoneDemo() {
	fmt.Println("\n--- Time Zone Operations ---")

	t := time.Date(2023, time.November, 10, 15, 30, 0, 0, time.UTC)

	// Load location
	ny, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}

	// Convert to location
	nyTime := t.In(ny)
	fmt.Println("NY time:", nyTime)

	// Convert to UTC
	utcTime := t.UTC()
	fmt.Println("UTC time:", utcTime)
}

func truncationDemo() {
	fmt.Println("\n--- Truncation/Rounding ---")

	t := time.Date(2023, time.November, 10, 15, 30, 45, 0, time.UTC)

	// Truncate
	trunc := t.Truncate(time.Hour)
	fmt.Println("Truncated to hour:", trunc)

	// Round
	rounded := t.Round(time.Hour)
	fmt.Println("Rounded to hour:", rounded)
}

func durationDemo() {
	fmt.Println("\n--- Duration Operations ---")

	// Create duration
	d := 2*time.Hour + 30*time.Minute

	fmt.Println("Duration:", d)

	// Convert to units
	fmt.Println("Hours:", d.Hours())
	fmt.Println("Minutes:", d.Minutes())
	fmt.Println("Seconds:", d.Seconds())

	// String representation
	fmt.Println("String:", d.String())
}
