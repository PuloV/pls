package file

import (
	"fmt"
	"strings"
)

var ShowConvertedSize bool
var ShowType bool
var ShowSizePercentage bool
var ShowFilesCount bool
var ShowFilesPercentage bool

type File struct {
	Path     string
	Size     int64
	IsDir    bool
	Subfiles []File
	Parent   *File
}

func (f File) Files() (filesCount int) {
	if f.IsDir {
		for _, file := range f.Subfiles {
			filesCount += file.Files()
		}
	} else {
		filesCount = 1
	}
	return
}

func (f File) String() string {
	fileInfo := []string{
		f.Path,
	}
	if ShowConvertedSize {
		fileInfo = append(fileInfo, f.ConvertedSize())
	}
	if ShowType {
		fileInfo = append(fileInfo, f.Type())
	}
	if ShowSizePercentage {
		fileInfo = append(fileInfo, f.SizePercentage())
	}
	if ShowFilesCount {
		fileInfo = append(fileInfo, f.FilesCount())
	}
	if ShowFilesPercentage {
		fileInfo = append(fileInfo, f.FilesPercentage())
	}

	return strings.Join(fileInfo, " | ")
}

func (f File) Type() string {
	if f.IsDir {
		return "dir"
	} else {
		return "file"
	}
}

func (f File) SizePercentage() string {
	if f.Parent != nil {
		persentage := float64(f.Size) / float64(f.Parent.Size) * 100
		return fmt.Sprintf("%.2f%s", persentage, "%")
	} else {
		return "100&"
	}
}

func (f File) FilesPercentage() string {
	if f.Parent != nil {
		persentage := float64(f.Files()) / float64(f.Parent.Files()) * 100
		return fmt.Sprintf("%.2f%s", persentage, "%")
	} else {
		return "100%"
	}
}

func (f File) FilesCount() string {
	return fmt.Sprintf("%d", f.Files())
}

func (f File) ConvertedSize() string {
	return detectFileSize(float64(f.Size), []string{"B", "KB", "MB", "GB", "TB"})
}

func detectFileSize(size float64, labels []string) string {
	if size < 1000 || len(labels) == 1 {
		return fmt.Sprintf("%f %s", size, labels[0])
	} else {
		return detectFileSize(size/1024, labels[1:])
	}
}
