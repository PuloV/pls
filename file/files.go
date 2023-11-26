package file

// Type used to sort collection of files
type Files []File

func (s Files) Len() int      { return len(s) }
func (s Files) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// Type used to sort the files by their size
type BySize struct{ Files }

func (s BySize) Less(i, j int) bool { return s.Files[i].Size > s.Files[j].Size }
