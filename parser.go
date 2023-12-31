package goxbrl

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
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

	IndependentAuditorsReportConsolidatedTextBlock     string `xml:"IndependentAuditorsReportConsolidatedTextBlock"`
	AuditFirm1Consolidated                             string `xml:"AuditFirm1Consolidated"`
	CPA1AuditFirm1Consolidated                         string `xml:"CPA1AuditFirm1Consolidated"`
	CPA2AuditFirm1Consolidated                         string `xml:"CPA2AuditFirm1Consolidated"`
	KeyAuditMattersConsolidatedTextBlock               string `xml:"KeyAuditMattersConsolidatedTextBlock"`
	OverviewKAMConsolidatedTextBlock                   string `xml:"OverviewKAMConsolidatedTextBlock"`
	ShortDescriptionKAMConsolidated                    string `xml:"ShortDescriptionKAMConsolidated"`
	DescriptionIncludingReasonKAMConsolidatedTextBlock string `xml:"DescriptionIncludingReasonKAMConsolidatedTextBlock"`
	ReferenceKAMConsolidated                           string `xml:"ReferenceKAMConsolidated"`
	Reference3KAMConsolidated                          string `xml:"Reference3KAMConsolidated"`
	AuditorsResponseKAMConsolidatedTextBlock           string `xml:"AuditorsResponseKAMConsolidatedTextBlock"`
	OtherInformationConsolidatedTextBlock              string `xml:"OtherInformationConsolidatedTextBlock"`
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
	if filepath.Ext(filePath) != "xbrl" {
		return nil, fmt.Errorf("Error: %s", "The extension is not xbrl.")
	}

	xbrlData := &XBRLData{}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if err := xml.Unmarshal(data, xbrlData); err != nil {
		return nil, err
	}

	return xbrlData, nil
}
