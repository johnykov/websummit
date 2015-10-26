package main

import (
    "net/http"
    "io"
    "os"
    "strconv"
    "fmt"
)

func main() {
    folder := "responses"
    os.Mkdir(folder, 0777)
    maxLimit := 185

    for i := 0; i < maxLimit; i++ {
        strNr := strconv.Itoa(i)
        download("https://api.cilabs.net/v1/conferences/ws15/info/attendees?page=" + strNr, "respone" + strNr + ".json", folder)
    }
}

func download(address string, filename string, folder string) {
    resp, err := http.Get(address)
    check(err)
    defer resp.Body.Close()
    out, err := os.Create(folder + "/" + filename)
    check(err)
    fmt.Println("downloaded " + filename)
    defer out.Close()
    io.Copy(out, resp.Body)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}