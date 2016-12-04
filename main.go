package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing filename")
		os.Exit(2)
	}
	ds := "2016-11-29T05:44:56.000000Z"
	t, _ := time.Parse(time.RFC3339, ds)
	fmt.Println(t, "Metadata")
	fmt.Println(t.Local(), "Metadata in localtime")

	filename := os.Args[1]
	info, err := os.Stat(filename)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(info.ModTime(), "ModTime")

	duration := t.Sub(info.ModTime())
	switch {
	case duration < 0:
		fmt.Println("actual capture time in the future by", -duration)
		err = os.Chtimes(filename, t, t)
		if err != nil {
			fmt.Println("Set filetime to capture time", t)
		}
	case duration > 0:
		fmt.Println("actual capture time is in the past by", duration)
		err = os.Chtimes(filename, t, t)
		if err != nil {
			fmt.Println("Set filetime to capture time", t)
		}
	case duration == 0:
		fmt.Println("ModTime time matches Metadata")
		os.Exit(0)
	}

	fmt.Println("Changing time to", t)
	err = os.Chtimes(filename, t, t)

	//info, err = os.Stat(filename)
	//if err != nil {
	//	log.Panic(err)
	//}
	//fmt.Println(info.ModTime())

	//	data, err := Probe(filename)
	//	if err != nil {
	//		panic(err)
	//	}
}
