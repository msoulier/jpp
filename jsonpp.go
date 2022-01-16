package main

import (
    "encoding/json"
    "os"
    "bufio"
    "io"
    "io/ioutil"
    )

func main() {
    objmap := make(map[string]interface{})
    var reader *bufio.Reader = nil

    if len(os.Args) > 1 {
        if os.Args[1] == "-h" {
            os.Stdout.WriteString("Usage: jsonpp [input file, - or blank for stdin]\n")
            os.Exit(0)
        }
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
    /*
     * To pull a value out it is done like so:
     * details_map := objmap["details"].(map[string]interface{})
     * note: value is an interface{} at this point
     * active_calls := details_map["active_calls"]
     */
}
