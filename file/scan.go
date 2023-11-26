package file

import (
	"fmt"
	"io/fs"
	"os"
	"sync"
)

func ScanFiles(files []string) (scanedFiles []File) {
	for _, filePath := range files {
		scanedFiles = append(scanedFiles, scanFilePath(filePath, nil))
	}
	return
}

func scanFilePath(filePath string, parent *File) (f File) {
	osFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening:", err)
		return
	}
	fileStat, err := osFile.Stat()

	if err != nil {
		fmt.Println("Error getting stats:", err)
		return
	}

	f.Path = filePath
	f.Parent = parent

	if fileStat.IsDir() {
		f.IsDir = true
		f.Subfiles = []File{}

		files, err := osFile.Readdir(0)
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}

		subfilesChan := make(chan File)
		wg := sync.WaitGroup{}
		wg.Add(len(files))

		go func() {
			wg.Wait()
			close(subfilesChan)
		}()

		for _, file := range files {
			go func(d fs.FileInfo, p *File) {
				defer wg.Done()
				newPath := fmt.Sprintf("%s/%s", filePath, d.Name())
				subfilesChan <- scanFilePath(newPath, p)
			}(file, &f)
		}

		for file := range subfilesChan {
			if file.Path != "" {
				f.Subfiles = append(f.Subfiles, file)
				f.Size = f.Size + file.Size
			}
		}

	} else {
		f.Path = filePath
		f.Size = fileStat.Size()
		f.IsDir = false
	}

	return
}
