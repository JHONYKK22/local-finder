package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
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
	fmt.Println("Name:", *name)

	_, errorPath := os.Stat(*path)

	if errorPath != nil {

		fmt.Println("Error!!", errorPath)

		//*path, _ = os.Getwd()
		return
	}

	fmt.Println()

	startTime := time.Now()

	total, err := WalkDirConcurrent(*path, *name)

	defer fmt.Println()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("\n\033[34mTotal matches\033[0m\033[1;32m", total, "\033[0m")

	timeElapsed := time.Since(startTime)
	fmt.Printf("\n\033[1;37mTime Elapsed: %v\033[0m\n", timeElapsed)

}

func skipDir(d fs.DirEntry) (skipDir bool) {

	name := d.Name()

	_, exists := namesToSkip[name]

	if exists {
		skipDir = true
	}

	return
}

func WalkDirConcurrent(root string, nameToFind string) (int, error) {

	nameToFind = strings.ToLower(nameToFind)

	var wg sync.WaitGroup

	dirs := make(chan string, 100)
	results := make(chan string, 100)

	workerCount := 20

	for i := 0; i < workerCount; i++ {
		go func() {

			for dir := range dirs {

				entries, err := os.ReadDir(dir)

				if err == nil {

					for _, entry := range entries {

						full := filepath.Join(dir, entry.Name())

						if strings.Contains(strings.ToLower(entry.Name()), nameToFind) {
							results <- full
						}

						if entry.IsDir() && !skipDir(entry) {
							wg.Add(1)
							// Prevent deadlock
							go func(p string) {
								dirs <- p
							}(full)
						}
					}
				}
				wg.Done()
			}
		}()
	}

	wg.Add(1)
	go func() {
		dirs <- root
	}()

	go func() {
		wg.Wait()
		close(dirs)
		close(results)
	}()

	total := 0

	for r := range results {
		fmt.Println(r)
		total++
	}

	return total, nil
}

var namesToSkip = map[string]string{

	// windows system folders
	//"Program Files" : "",
	"Program Files (x86)": "",
	"ProgramData":         "",
	"AppData":             "",
	"Temp":                "",
	"Windows":             "",

	//Linux system folders
	"system":       "",
	"library":      "",
	//"applications": "",
	//"usr":          "",
	"bin":          "",
	"sbin":         "",
	"var":          "",
	"dev":          "",
	"proc":         "",
	"run":          "",

	//dev folders
	".git":         "",
	".venv":        "",
	".vscode":      "",
	".idea":        "",
	".gradle":      "",
	"node_modules": "",
	"postgres":     "",
	"build":        "",
	"target":       "",
	"dist":         "",
	"venv":         "",
	"__pycache__":  "",
}
