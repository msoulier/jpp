package main

import (
    "os"
    "bufio"
    "bytes"
    "io"
    "github.com/Jeffail/gabs/v2"
    "github.com/charmbracelet/log"
    "fmt"
    "flag"
)

const (
    BUFSIZE int = 4096
)

var (
    debug bool = false
)

func init() {
    flag.BoolVar(&debug, "debug", false, "Debug logging")
    flag.Parse()
}

func main() {
    var buffer bytes.Buffer
    reader := bufio.NewReader(os.Stdin)
    bytes_count := 0
    if debug {
        log.SetLevel(log.DebugLevel)
    }

    // Read all input data, grow as required
    for {
        tempbuf := make([]byte, BUFSIZE)
        nbytes, err := reader.Read(tempbuf)
        if err != nil {
            if err != io.EOF {
                panic(err)
            } else {
                break
            }
        }
        log.Debugf("read %d bytes", nbytes)
        bytes_count += nbytes
        buffer.Write(tempbuf)
        if nbytes < BUFSIZE {
            break
        }
    }
    j, err := gabs.ParseJSON(buffer.Bytes()[:bytes_count])
    if err != nil {
        panic(err)
    }
    fmt.Printf(j.StringIndent("", "    "))
}
