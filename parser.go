package goxbrl

import (
	"encoding/xml"
	"io"
	"os"
)

type XBRLData struct {
	XMLName  xml.Name `xml:"xbrl"`
	JpcrpCor string   `xml:"jpcrp_cor,attr"`
	Link     string   `xml:"link,attr"`
	Xbrldi   string   `xml:"xbrldi,attr"`
	Xbrli    string   `xml:"xbrli,attr"`
	Xlink    string   `xml:"xlink,attr"`
	Xsi      string   `xml:"xsi,attr"`

	SchemaRef SchemaRef `xml:"schemaRef"`
	Contexts  []Context `xml:"context"`
}

type SchemaRef struct {
	HRef string `xml:"href,attr"`
	Type string `xml:"type,attr"`
}

type Context struct {
	Entity   Entity   `xml:"entity"`
	Period   Period   `xml:"period"`
	Scenario Scenario `xml:"scenario"`
}

type Entity struct {
	Identifier string `xml:"identifier"`
}

type Period struct {
	Instant string `xml:"instant"`
}

type Scenario struct {
	ExplicitMember string `xml:"explicitMember"`
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
