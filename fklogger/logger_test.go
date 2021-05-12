package fklogger

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/fkmatsuda/go-utils/fkfile"
	"github.com/fkmatsuda/go-utils/fksystem"
)

var (
	tempFiles []string
	lorem     []string
)

func initTestLog() string {
	file, err := ioutil.TempFile(os.TempDir(), "log")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	filePath := file.Name()
	dir := filepath.Dir(filePath)
	filename := filepath.Base(filePath)

	tempFiles = append(tempFiles, filePath)

	RegisterAppender(NewFileAppender(dir, filename, filename, 1024))

	return filePath

}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	tempFiles = make([]string, 0)
	lorem = []string{
		"Contrary to popular belief, Lorem Ipsum is not simply random text. It has roots in a piece of classical Latin literature from 45 BC, making it over 2000 years old. Richard McClintock, a Latin professor at Hampden-Sydney College in Virginia, looked up one of the more obscure Latin words, consectetur, from a Lorem Ipsum passage, and going through the cites of the word in classical literature, discovered the undoubtable source. Lorem Ipsum comes from sections 1.10.32 and 1.10.33 of \"de Finibus Bonorum et Malorum\" (The Extremes of Good and Evil) by Cicero, written in 45 BC. This book is a treatise on the theory of ethics, very popular during the Renaissance. The first line of Lorem Ipsum, \"Lorem ipsum dolor sit amet..\", comes from a line in section 1.10.32.",
	}
}

func teardown() {
	for _, fileName := range tempFiles {
		os.Remove(fileName)
		var err error
		for i := 1; err == nil; i++ {
			rollingFileName := fmt.Sprintf("%s.%d.gz", fileName, i)
			_, err = os.Stat(rollingFileName)
			if err == nil {
				os.Remove(rollingFileName)
			}
		}
	}
}

func TestInfo(t *testing.T) {
	fileName := initTestLog()
	Info(lorem[0])

	textFile, err := fkfile.ReadTextFile(fileName)
	if err != nil {
		t.Error(err)
	}
	r := regexp.MustCompile(`(INFO:\s+\d{4}/\d{2}/\d{2}\s+\d{2}:\d{2}:\d{2}\s+)(.*)`)
	m := r.FindStringSubmatch(textFile)
	if m == nil || len(m) != 3 {
		t.Errorf("Log content is invalid")
	} else if m[2] != lorem[0] {
		t.Errorf("Expected log: \"%s\"\nbut was\"%s\"", lorem[0], textFile)
	}

}

func TestError(t *testing.T) {
	fileName := initTestLog()
	Error(lorem[0])

	textFile, err := fkfile.ReadTextFile(fileName)
	if err != nil {
		t.Error(err)
	}
	r := regexp.MustCompile(`(ERROR:\s+\d{4}/\d{2}/\d{2}\s+\d{2}:\d{2}:\d{2}\s+)(.*?)(\s+Stack\s+trace\s+=>.*?:\s+.*?\s+.*logger_test\.go:\d+.*)`)
	m := r.FindStringSubmatch(textFile)
	if m == nil || len(m) != 4 {
		t.Errorf("Log content is invalid:\n\"%s\"", textFile)
	} else if m[2] != lorem[0] {
		t.Errorf("Expected log: \"%s\"\nbut was\"%s\"", lorem[0], textFile)
	}

}

func TestRolling(t *testing.T) {
	fileName := initTestLog()
	Error(lorem[0])
	Error(lorem[0])

	dir := filepath.Dir(fileName)
	file := filepath.Base(fileName)

	rollingFile := fmt.Sprintf("%s%s%s.1.gz", dir, fksystem.DirSeparator(), file)

	_, err := os.Stat(rollingFile)

	if err != nil && errors.Is(err, os.ErrNotExist) {
		t.Error("Rolling file must exist")
	}

}
func TestNoRolling(t *testing.T) {
	fileName := initTestLog()
	Info(lorem[0])
	Info(lorem[0])

	dir := filepath.Dir(fileName)
	file := filepath.Base(fileName)

	rollingFile := fmt.Sprintf("%s%s%s.1.gz", dir, fksystem.DirSeparator(), file)

	_, err := os.Stat(rollingFile)

	if err == nil && errors.Is(err, os.ErrExist) {
		t.Error("Rolling file must not exist")
	}

}
