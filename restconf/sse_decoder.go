package restconf

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

const (
	sseDataPrefix = "data: "
)

// we only have to decode whatever server is sending.  so far it's just "data: " fields
func decodeSse(in io.Reader) <-chan []byte {
	events := make(chan []byte)
	r := bufio.NewReader(in)
	go func() {
		defer close(events)
		var buff bytes.Buffer
		send := func() {
			if buff.Len() > 0 {
				orig := buff.Bytes()
				dup := make([]byte, len(orig))
				copy(dup, orig)
				events <- dup
				buff.Reset()
			}
		}
		for {
			line, err := r.ReadBytes('\n')
			size := len(line)
			if size <= 1 {
				send()
			} else if strings.HasPrefix(string(line), sseDataPrefix) {
				end := size
				if line[end-1] == '\n' {
					end--
				}
				chunk := line[len(sseDataPrefix):end]
				buff.Write(chunk)
			}
			if err != nil {
				// EOF or other; stream is no longer
				send()
				return
			}
		}
	}()
	return events
}
