package table

import (
	"fmt"
	"strconv"
	"strings"
)

type sTitle struct {
	name   string
	key    string
	maxLen int
}

type STable struct {
	title []sTitle
	data  []map[string]string
}

func New() *STable {
	obj := new(STable)
	obj.data = make([]map[string]string, 0)
	return obj
}

func getStrLen(str string) int {
	count := 0
	preIndex := 0
	lastLength := 0
	for index, _ := range str {
		lastLength = index - preIndex
		if lastLength == 3 {
			lastLength = 2
		}
		count += lastLength
		preIndex = index
	}
	lastLength = len(str) - preIndex
	if lastLength == 3 {
		lastLength = 2
	}
	count += lastLength
	return count
}

func printLen(str string, maxLen int) {
	space := maxLen - getStrLen(str)
	str += strings.Repeat(" ", space)
	fmt.Print(str)
}

func (obj *STable) ShowAll() {
	obj.Show(0)
}

// lineNumber==0: SHow all lines
func (obj *STable) Show(lineNumber int) {
	for i, _ := range obj.title {
		title := &obj.title[i]
		title.maxLen = getStrLen(title.name)
		for _, line := range obj.data {
			strLen := getStrLen(line[title.key])
			if strLen > title.maxLen {
				title.maxLen = strLen
			}
		}
		title.maxLen += 2
	}

	//Print title
	printLen("#", 5)
	for _, title := range obj.title {
		printLen(title.name, title.maxLen)
	}
	fmt.Print("\n")

	//Print data
	for i, line := range obj.data {
		printLen(fmt.Sprintf("%03d", i+1), 5)
		for _, title := range obj.title {
			v := line[title.key]
			printLen(v, title.maxLen)
		}
		fmt.Print("\n")
		if lineNumber > 0 && (i+1) == lineNumber {
			break
		}
	}
}

func (obj *STable) AddTitle(name string, key string) {
	var title sTitle
	title.name = name
	title.key = key
	obj.title = append(obj.title, title)
}

func (obj *STable) AddLine(line map[string]string) {
	obj.data = append(obj.data, line)
}

func (obj *STable) Sort(key string, ascending bool) {
	for i, _ := range obj.data {
		for j, _ := range obj.data {
			if j <= i {
				continue
			}
			v0, _ := strconv.ParseFloat(obj.data[i][key], 32)
			v1, _ := strconv.ParseFloat(obj.data[j][key], 32)
			if (v1 > v0 && ascending == false) || (v1 < v0 && ascending == true) {
				temp := obj.data[i]
				obj.data[i] = obj.data[j]
				obj.data[j] = temp
			}
		}
	}
}

