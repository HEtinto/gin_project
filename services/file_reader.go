package services

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

type Reader struct {
	file *os.File
}

type StringChannel struct {
	strings chan string
	done    chan bool
}

func NewStringChannel() *StringChannel {
	return &StringChannel{
		strings: make(chan string),
		done:    make(chan bool),
	}
}

// check channel is closed
func (c *StringChannel) IsClosed() bool {
	select {
	case <-c.done:
		return true
	default:
		return false
	}
}

// NewReader 新建一个Reader
func NewReader(filePath string) (*Reader, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return &Reader{file: file}, nil
}

// Close file handles
func (r *Reader) Close() error {
	if r.file != nil {
		r.file.Close()
	}
	return nil
}

// Read file by line
func (r *Reader) ReadLines() (sc *StringChannel, err error) {
	sc = NewStringChannel()
	// 将文件指针重置到文件开头
	r.file.Seek(0, io.SeekStart)
	reader := bufio.NewReader(r.file)
	go func() {
		defer func() {
			fmt.Println("done info send.")
			close(sc.strings)
			sc.done <- true
		}()
		for {
			line, err := reader.ReadString('\n')
			if err == nil {
				sc.strings <- line
			} else {
				if err == io.EOF {
					break // 当遇到文件结束时跳出循环
				} else {
					fmt.Println("Error reading line", err)
					break // 遇到其他错误时也跳出循环
				}
			}
		}
	}()
	return
}

// Filter the file lines by regex pattern
func (r *Reader) FilterLines(pattern string) (strSlice []string, err error) {
	sc, err := r.ReadLines()
	if err == nil {
		for {
			if sc.IsClosed() {
				break
			}
			// re := regexp.MustCompile(pattern)
			str := <-sc.strings
			matched, err := regexp.MatchString(pattern, str)
			if err != nil {
				fmt.Printf("Error matching regex: %v\n", err)
				continue
			}
			fmt.Printf("matched: %v\n", matched)
			if matched {
				strSlice = append(strSlice, str)
			}
		}
	}
	return strSlice, err
}
