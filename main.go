package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/djherbis/times"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing filename")
		os.Exit(2)
	}
	filename := os.Args[1]

	fmt.Println("Probing", filename, "for capture time in Metadata")
	data, err := Probe(filename)
	if err != nil {
		panic(err)
	}

	// fmt.Println(data.Format.Tags)
	creation_time := data.Format.Tags["creation_time"]

	t, err := time.Parse(time.RFC3339, creation_time)
	if err != nil {
		panic(err)
	}

	fmt.Println(t, "Metadata")
	fmt.Println(t.Local(), "Metadata in localtime")

	info, err := os.Stat(filename)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(info.ModTime(), "File modification time")

	tb, err := times.Stat(filename)
	if err != nil {
		log.Fatal(err.Error())
	}

	if tb.HasBirthTime() {
		fmt.Println(tb.BirthTime(), "File creation time")

		duration := t.Sub(tb.BirthTime())
		switch {
		case duration < 0:
			fmt.Println(filename, "capture time in the future by", -duration)
		case duration > 0:
			fmt.Println(filename, "capture time is in the past by", duration)
		case duration == 0:
			fmt.Println(filename, "Filetime matches Metadata capture time")
			os.Exit(0)
		}
	}
	// Not neccessary
	err = os.Chtimes(filename, t, t)
	if err == nil {
		fmt.Println("Set Filetime to capture time", t)
	}
	// -d date           # creation date (mm/dd/[yy]yy [hh:mm[:ss] [AM | PM]])*
	fmt.Printf("setfile -d \"%s\" %s\n", t.Local().Format("01/02/2006 15:04:05"), filename)

}
