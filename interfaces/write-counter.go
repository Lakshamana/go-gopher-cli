package interfaces

import (
	"fmt"
	"strings"

	"github.com/dustin/go-humanize"
)

type WriteCounter struct {
  Total uint64
  progressCounter int
}

func NewWriteCounter() *WriteCounter {
  return &WriteCounter{}
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
  n := len(p)
  wc.Total += uint64(n)
  wc.PrintProgress()
  return n, nil
}

func (wc *WriteCounter) PrintProgress() {
 fmt.Printf("\r%s", strings.Repeat(" ", 35))

  steps := [4]string{"|", "/", "-", "\\"}
  fmt.Printf("\r> Downloading file (%s) %s", humanize.Bytes(wc.Total), steps[wc.progressCounter % 4])
  wc.progressCounter++
}
