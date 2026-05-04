package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/JHONYKK22/local-finder/finder"
)

func main() {

	flag.Usage = func() {
		
		fmt.Println("\nUsage: program_build_name --find=file_or_folder_name [--path=starting_path]")
		fmt.Println("Description: search for files or directories by name starting from a specified path (default is current directory).")
		fmt.Println()
		flag.PrintDefaults()
		fmt.Println()
	}

	pathUsage := "Enter the starting path"
	path := flag.String("path", ".", pathUsage)

	nameUsage := "Enter the name of the file/folder to find"
	name := flag.String("find", "", nameUsage)

	flag.Parse()

	if *name == "" {
		fmt.Println("\033[31m", nameUsage, "\033[0m")
		return
	}

	fmt.Println("Path:", *path)
	fmt.Println("Find:", *name)

	_, errorPath := os.Stat(*path)

	if errorPath != nil {

		fmt.Println("Error!!", errorPath)

		//*path, _ = os.Getwd()
		return
	}

	fmt.Println()

	startTime := time.Now()

	total, err := finder.WalkDirConcurrent(*path, *name)

	defer fmt.Println()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("\n\033[34mTotal matches\033[0m\033[1;32m", total, "\033[0m")

	timeElapsed := time.Since(startTime)
	fmt.Printf("\n\033[1;37mTime Elapsed: %v\033[0m\n", timeElapsed)

}