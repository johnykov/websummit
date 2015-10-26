package main

import (
    "encoding/csv"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "strconv"
    "strings"
)

type Json struct {
    RecordID int64  `json:"id"`
    Bio      string `json:"bio"`
    Name     string `json:"name"`
    Company  string `json:"company"`
    Country  string `json:"country"`
    Avatar   string `json:"avatar_url"`
    Medium   string `json:"medium_image"`
    Career   string `json:"career"`
}
type Attendees struct {
    Fruits []Json `json:"attendees"`
}


func main() {
    folder := "csv"
    os.Mkdir(folder, 0777)
    //    json2csv("respone0.json", folder)
    files, _ := ioutil.ReadDir("./responses")
    for _, f := range files {
        json2csv(f.Name(), folder)
//        fmt.Println(f.Name()[:len(f.Name()) - 5] + " done")
    }
}

func removeNewLines(src string) string {
    return strings.Replace(strings.Replace(src, "\r\n", "", -1), "\n", "", -1)
}

func json2csv(filename string, folder string) {
    // reading data from JSON File
    data, err := ioutil.ReadFile("./responses/" + filename)
    if err != nil {
        fmt.Println(err)
    }

    // Unmarshal JSON data
    var d Attendees
    err = json.Unmarshal([]byte(data), &d)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(d.Fruits[0])
    // Create a csv file
    f, err := os.Create("./" + folder + "/" + filename[:len(filename) - 5] + ".csv")
    if err != nil {
        fmt.Println(err)
    }
    defer f.Close()
    // Write Unmarshaled json data to CSV file
    w := csv.NewWriter(f)


    for _, obj := range d.Fruits {
        var record []string
        record = append(record, strconv.FormatInt(obj.RecordID, 10))
        record = append(record, obj.Name)
        record = append(record, obj.Career)
        record = append(record, obj.Company)
        record = append(record, obj.Country)
        record = append(record, removeNewLines(obj.Bio))
        record = append(record, obj.Avatar)
        record = append(record, obj.Medium)
        w.Write(record)
    }
    w.Flush()
}
