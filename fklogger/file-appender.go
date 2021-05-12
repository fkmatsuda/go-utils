package fklogger

import (
	"errors"
	"fmt"
	"os"

	"github.com/fkmatsuda/go-utils/fkfile"
	"github.com/fkmatsuda/go-utils/fksystem"
)

type fileAppenderWriter struct {
	dir, file string
	size      uint
}

func (w fileAppenderWriter) fullFilePath() string {
	return fmt.Sprintf("%s%s%s", w.dir, fksystem.DirSeparator(), w.file)
}

func (w fileAppenderWriter) fileRotate() error {

	err := w.moveLogFileToNext(0)
	if err != nil {
		return err
	}

	return nil

}

func (w fileAppenderWriter) indexFileName(idx int) string {
	fullFilePath := w.fullFilePath()
	if idx <= 0 {
		return fullFilePath
	}
	return fmt.Sprintf("%s.%d.gz", fullFilePath, idx)
}

func (w fileAppenderWriter) moveLogFileToNext(idx int) error {

	maxLogFiles := 32

	currentFileName := w.indexFileName(idx)

	_, err := os.Stat(currentFileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if idx >= maxLogFiles {
		os.Remove(currentFileName)
		return nil
	}

	err = w.moveLogFileToNext(idx + 1)
	if err != nil {
		return err
	}

	nextFileName := w.indexFileName(idx + 1)
	if idx == 0 {
		return fkfile.CompressFile(currentFileName, nextFileName, true)
	}

	return os.Rename(currentFileName, nextFileName)

}

func (w fileAppenderWriter) Write(p []byte) (n int, err error) {

	n = 0
	err = fksystem.EnsureDir(w.dir)
	if err != nil {
		return
	}

	fullFilePath := w.fullFilePath()

	stat, err := os.Stat(fullFilePath)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			return
		}
		err = nil
	}
	if int64(w.size) < stat.Size() {
		err = w.fileRotate()
		if err != nil {
			return
		}
	}

	f, err := os.OpenFile(fullFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return
	}
	defer f.Close()

	n, err = f.Write(p)
	if err == nil {
		err = f.Sync()
	}

	return

}

func NewFileAppender(logDir, logFile, logFileErr string, rollingSize uint) Appender {

	appender := initAppender(fileAppenderWriter{logDir, logFile, rollingSize}, fileAppenderWriter{logDir, logFileErr, rollingSize})

	return appender
}
