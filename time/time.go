package main

import (
	"fmt"
	"time"
)

func main() {
	testNow()
	testDuration()
	testSleep()
	testSince()
}

func testNow() {
	now := time.Now()
	fmt.Println(now)
}

func testDuration() {
	nanosecond := time.Nanosecond
	fmt.Println(nanosecond)

	microsecond := time.Microsecond
	fmt.Println(microsecond)

	millisecond := time.Millisecond
	fmt.Println(millisecond)

	second := time.Second
	fmt.Println(second)

	minute := time.Minute
	fmt.Println(minute)

	hour := time.Hour
	fmt.Println(hour)

	dayHours := 24
	day := time.Duration(dayHours) * time.Hour
	fmt.Println(day)

	year := 365 * day
	fmt.Println(year)
}

func testSleep() {
	time.Sleep(time.Second)
}

func testSince() {
	start := time.Now()

	time.Sleep(time.Second)

	since := time.Since(start)
	fmt.Println(since)
}
