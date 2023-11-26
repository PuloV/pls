package main

import (
	"flag"
	"fmt"
	"os"
	"sort"

	"github.com/PuloV/pls/file"
)

var showLargestFiles int
var showDepthFiles int
var withVersion bool

const INDEPTH_PREFIX = "    |"
const FILE_POINTER = " - > "
const CLI_VERSION = "0.0.2"

func init() {
	flag.IntVar(&showLargestFiles, "l", -1, "shows the largest max files")
	flag.IntVar(&showDepthFiles, "d", 1, "shows data in depth")
	flag.BoolVar(&withVersion, "v", false, "show cli version")
	flag.BoolVar(&file.ShowConvertedSize, "s", false, "show file size")
	flag.BoolVar(&file.ShowType, "t", false, "show file type")
	flag.BoolVar(&file.ShowSizePercentage, "sp", false, "show size percentage ")
	flag.BoolVar(&file.ShowFilesCount, "fc", false, "show files count version")
	flag.BoolVar(&file.ShowFilesPercentage, "fp", false, "show files count percentage ")
}

func main() {
	flag.Parse()

	if withVersion {
		fmt.Println("Version:", CLI_VERSION)
		return
	}

	wd, _ := os.Getwd()
	argsToList := flag.Args()
	listedDirs := []string{wd}
	if len(argsToList) > 0 {
		listedDirs = argsToList
	}
	scanedFiles := file.ScanFiles(listedDirs)

	fmt.Println("===========================")
	for _, scannedFile := range scanedFiles {
		fmt.Printf("- %s \n", scannedFile)
		if scannedFile.IsDir {
			displayFiles(scannedFile.Subfiles, showDepthFiles, INDEPTH_PREFIX)
		}

	}
	fmt.Println("===========================")
}

func displayFiles(files []file.File, depth int, prefix string) {
	if depth == 0 {
		return
	}
	if len(files) < showLargestFiles {
		showLargestFiles = len(files)
	}

	sort.Sort(file.BySize{files})

	var presentableFiles []file.File

	if showLargestFiles > 0 {
		presentableFiles = files[0:showLargestFiles]
	} else {
		presentableFiles = files
	}

	for _, file := range presentableFiles {
		fmt.Printf("%s%s%s\n", prefix, FILE_POINTER, file)
		if file.IsDir {
			displayFiles(file.Subfiles, depth-1, fmt.Sprintf("%s %s", INDEPTH_PREFIX, prefix))
		}
	}

}
