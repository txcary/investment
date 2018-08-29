package utils

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"time"
)

const (
	LogLevelError = iota
	LogLevelWarn
	LogLevelInfo
	LogLevelDebug
)

var logLevel int

func SetLogLevel(level int) {
	if level == LogLevelError ||
		level == LogLevelInfo ||
		level == LogLevelWarn ||
		level == LogLevelDebug {
		logLevel = level
	} else {
		Error("Log Level Number not correct!")
	}
}

func init() {
	log.SetPrefix("[" + filepath.Base(os.Args[0]) + "]")
	log.SetFlags(log.Ldate | log.Ltime)
}

func HttpDownload(url string, out string) {
	os.Remove(out)
	resp, err := http.Get(url)
	Check(err)
	defer resp.Body.Close()
	file, err := os.Create(out)
	defer file.Close()
	Check(err)
	io.Copy(file, resp.Body)
}

func Slash() string {
	var ostype = runtime.GOOS
	if ostype == "windows" {
		return "\\"
	}
	if ostype == "linux" {
		return "/"
	}
	return "/"
}

func Gopath() string {
	return os.Getenv("GOPATH")
}

func TmpPath() string {
	return os.Getenv("TMP")
}

func Pwd() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}

func logPrint(prefix string, msg string) {
	pc, _, _, _ := runtime.Caller(2)
	function := runtime.FuncForPC(pc)
	file, line := function.FileLine(pc)
	file = path.Base(file)
	funcname := function.Name()
	log.Printf("[%s][%s:%04d:%s()] %s\n", prefix, file, line, funcname, msg)
}

func Warn(msg string) {
	if logLevel >= LogLevelWarn {
		logPrint("WRN", msg)
	}
}

func Info(msg string) {
	if logLevel >= LogLevelInfo {
		logPrint("INF", msg)
	}
}

func Debug(msg string) {
	if logLevel >= LogLevelDebug {
		logPrint("DBG", msg)
	}
}

func Error(msg string) {
	logPrint("ERR", msg)
}

func Check(err error) bool {
	if err != nil {
		logPrint("ERROR", err.Error())
		return true
	}
	return false
}

func CheckNil(ptr interface{}) bool {
	if reflect.TypeOf(ptr).Kind() == reflect.Ptr {
		if reflect.ValueOf(ptr).IsNil() {
			logPrint("ERROR", "nil pointer")
			return true
		}
	}
	return false
}

func TimeString() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func TmpFile() string {
	return (TmpPath() + Slash() + TimeString())
}

func GbkToUtf8(str string) (string, error) {
	s := []byte(str)
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return "", e
	}
	return string(d), nil
}

func Utf8ToGbk(str string) (string, error) {
	s := []byte(str)
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return "", e
	}
	return string(d), nil
}

func Exec(name string, arg ...string) (string, string, error) {
	cmd := exec.Command(name, arg...)
	stdout, err := cmd.StdoutPipe()
	Check(err)
	defer stdout.Close()
	stderr, err := cmd.StderrPipe()
	Check(err)
	defer stderr.Close()

	err = cmd.Start()
	Check(err)

	outb, err := ioutil.ReadAll(stdout)
	Check(err)
	errb, err := ioutil.ReadAll(stderr)
	Check(err)
	return string(outb), string(errb), err
}

func Average(args ...float64) float64 {
	var total float64 = 0
	numb := float64(len(args))
	for _, arg := range args {
		total += arg / numb
	}
	return total
}

func Max(args ...float64) float64 {
	var max float64 = args[0]
	for _, arg := range args {
		if arg > max {
			max = arg
		}
	}
	return max
}

func Min(args ...float64) float64 {
	var min float64 = args[0]
	for _, arg := range args {
		if arg < min {
			min = arg
		}
	}
	return min
}

func StringCount(str string) int {
	return strings.Count(str, "") - 1
}

func UtcStringToYear(str string) int {
	layout := "2006-01-02T15:04:05.000Z"
	t, err := time.Parse(layout, str)
	Check(err)

	return t.Year()
}
