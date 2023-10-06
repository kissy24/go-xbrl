package goxbrl

import (
	"encoding/xml"
	"io"
	"os"
)

type XBRLData struct {
	XMLName  xml.Name `xml:"xbrl"`
	JpcrpCor string   `xml:"jpcrp_cor,attr"`
	Xlink    string   `xml:"xlink,attr"`

	SchemaRef xml.Name `xml:"schemaRef"`
	Context   xml.Name `xml:"context"`
}

func ReadXBRL(filePath string) (*XBRLData, error) {
	// XBRLデータの構造体を初期化
	xbrlData := &XBRLData{}

	// XBRLファイルを開く
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// ファイルの内容を読み込む
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// XMLデコード
	if err := xml.Unmarshal(data, xbrlData); err != nil {
		return nil, err
	}

	return xbrlData, nil
}
