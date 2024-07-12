package services

import (
	"fmt"

	"github.com/hpcloud/tail"
	"github.com/pkg/errors"
)

type FileWatcher struct {
	watcher *tail.Tail
}

type FileWatch interface {
	Open(fileName string) error
	Close() error
	GetOneNewLine() (string, error)
}

// Get a FileWatcher obj
func NewFileWatcher() *FileWatcher {
	return &FileWatcher{}
}

// Open file which you want to watch
func (fileWatcher *FileWatcher) Open(fileName string) error {
	t, err := tail.TailFile(fileName, tail.Config{
		Location: &tail.SeekInfo{Offset: 0, Whence: 2}, Follow: true, ReOpen: false})
	if err != nil {
		return err
	}
	fileWatcher.watcher = t
	return nil
}

// Close file which you open
func (fileWatcher *FileWatcher) Close() error {
	if fileWatcher.watcher != nil {
		fileWatcher.watcher.Stop()
	}
	fmt.Println("FileWatcher closed finish...")
	return nil
}

// Get a new line from the watch file
func (fileWatcher *FileWatcher) GetOneNewLine() (string, error) {
	if fileWatcher.watcher == nil {
		return "No have new line", errors.New("Not open the file")
	}
	select {
	case line := <-fileWatcher.watcher.Lines:
		return line.Text, nil
	default:
		return "No have new line", nil
	}
}
