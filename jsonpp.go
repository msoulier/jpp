package main

import (
    "encoding/json"
    "os"
    "bufio"
    "io"
    "io/ioutil"
    )

func main() {
    objmap := make(map[string]string)
    var reader *bufio.Reader = nil

    if len(os.Args) > 1 {
        if os.Args[1] == "-" {
            reader = bufio.NewReader(os.Stdin)
        } else {
            // Open the file.
            if infile, err := os.Open(os.Args[1]); err != nil {
                panic(err)
            } else {
                reader = bufio.NewReader(infile)
            }
        }
    } else {
        reader = bufio.NewReader(os.Stdin)
    }

    if buffer, err := ioutil.ReadAll(reader); err != nil {
        if err != io.EOF {
            panic(err)
        }
    } else {
        if err := json.Unmarshal(buffer, &objmap); err != nil {
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
