package internal

import "fmt"

var kb float64 = 1024
var mb float64 = kb * kb
var gb float64 = mb * kb
var tb float64 = gb * kb
var pb float64 = tb * kb
var eb float64 = pb * kb

func OutputFmt(size int64, path string, isHumanFmt bool) string {
	sl := "B"
	if !isHumanFmt {
		return fmt.Sprintf("%d"+sl+"\t%s", size, path)
	}

	fs := float64(size)
	switch {
	case fs >= kb && fs < mb:
		sl = "KB"
		fs /= kb
	case fs >= mb && fs < gb:
		sl = "MB"
		fs /= mb
	case fs >= gb && fs < tb:
		sl = "GB"
		fs /= gb
	case fs >= tb && fs < pb:
		sl = "TB"
		fs /= tb
	case fs >= pb && fs < eb:
		sl = "PB"
		fs /= pb
	case fs >= eb:
		sl = "EB"
		fs /= eb
	}

	format := "%.1f" + sl + "\t%s"
	return fmt.Sprintf(format, fs, path)
}
