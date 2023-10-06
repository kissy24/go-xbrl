package main

import (
	"fmt"

	goxbrl "github.com/kissy24/go-xbrl"
)

func main() {
	filePath := "sample.xbrl"
	xbrlData, err := goxbrl.ReadXBRL(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// データのアクセス例:
	fmt.Println("TagName:", xbrlData.XMLName)
	fmt.Println("TagName:", xbrlData.Xlink)
	fmt.Println("SchemaRef:", xbrlData.SchemaRef)
	fmt.Println("Context:", xbrlData.Context)
}
