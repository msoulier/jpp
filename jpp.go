package main

import (
    "encoding/json"
    "os"
    "bufio"
    "io"
    )

func main() {
    objmap := make(map[string]string)
    buffer := make([]byte, 4096)
    reader := bufio.NewReader(os.Stdin)

    if nbytes, err := reader.Read(buffer); err != nil {
        if err != io.EOF {
            panic(err)
        }
    } else {
        if err := json.Unmarshal(buffer[:nbytes], &objmap); err != nil {
            panic(err)
        }
    }
    if b, err := json.MarshalIndent(objmap, "", "\t"); err != nil {
        panic(err)
    } else {
        os.Stdout.Write(b)
        os.Stdout.WriteString("\n")
    }
}
