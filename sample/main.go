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
	fmt.Println("XMLName:", xbrlData.XMLName)
	fmt.Println("Xlink:", xbrlData.Xlink)
	fmt.Println("JpcrpCor:", xbrlData.JpcrpCor)
	fmt.Println("Link:", xbrlData.Link)
	fmt.Println("Xbrldi:", xbrlData.Xbrldi)
	fmt.Println("Xbrli:", xbrlData.Xbrli)
	fmt.Println("Xsi:", xbrlData.Xsi)

	fmt.Println("SchemaRef HRef:", xbrlData.SchemaRef.HRef)
	fmt.Println("SchemaRef Type:", xbrlData.SchemaRef.Type)

	fmt.Println("Context Identifier:", xbrlData.Context.Entity.Identifier.Scheme)
}
