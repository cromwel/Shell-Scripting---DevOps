package main

import (
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "os"
)


// Appender holds the details about the XML
type Appender struct {
	XMLName  xml.Name `xml:"appender"`
	Text     string   `xml:",chardata"`
	Name     string   `xml:"name,attr"`
	Class    string   `xml:"class,attr"`
	MinLevel string   `xml:"minLevel"` // WARN
	Filter   struct {
		Text  string `xml:",chardata"`
		Class string `xml:"class,attr"`
		Level string `xml:"level"` // ERROR
	} `xml:"filter"`
}

func main() {
	// open file
    xmlFile, err := os.Open("initial.xml")
    if err != nil {
        fmt.Println(err)
	}

	// read file
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// unmarshal file data to struct
	var appender Appender
	xml.Unmarshal(byteValue, &appender)


	// change values
	appender.Class = "com"
	appender.Filter.Class = "hsk"
	appender.Filter.Level = "ERROR"

	// decode struct to bytes
	byteValue, _ = xml.Marshal(appender)

	// write to file
	ioutil.WriteFile("final.xml", byteValue, os.ModePerm)

	
	defer xmlFile.Close()
}