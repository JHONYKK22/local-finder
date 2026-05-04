package finder

import (
	"path/filepath"
	"strings"
	"sync"
	"io/fs"
	"os"
	"fmt"
)
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
