package poppler

import (
	"encoding/xml"
	"testing"
)

const testPDF = "./test/test.pdf"

func loadTestFile(t *testing.T) *Document {
	doc, err := NewFromFile(testPDF, "")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	return doc
}

func TestNewFromFile(t *testing.T) {
	loadTestFile(t)
}

func TestGetTitle(t *testing.T) {
	doc := loadTestFile(t)

	r := doc.GetTitle()
	t.Logf("Title: %s", r)
	// if r == "" {
	// 	t.Fail()
	// }
}

func TestGetPDFVersion(t *testing.T) {
	doc := loadTestFile(t)

	r := doc.GetPDFVersion()
	t.Logf("PDF Version: %s", r)
	// if r == "" {
	// 	t.Fail()
	// }
}

func TestGetTotalPages(t *testing.T) {
	doc := loadTestFile(t)

	r := doc.GetTotalPages()
	t.Logf("Total Pages: %d", r)
	if r == 0 {
		t.Fail()
	}
}

func TestGetAuthor(t *testing.T) {
	doc := loadTestFile(t)

	r := doc.GetAuthor()
	t.Logf("Author: %s", r)
	// if r == "" {
	// 	t.Fail()
	// }
}

func TestGetSubject(t *testing.T) {
	doc := loadTestFile(t)

	r := doc.GetSubject()
	t.Logf("Subject: %s", r)
	// if r == "" {
	// 	t.Fail()
	// }
}

func TestGetKeywords(t *testing.T) {
	doc := loadTestFile(t)

	r := doc.GetKeywords()
	t.Logf("Keywords: %s", r)
	// if r == "" {
	// 	t.Fail()
	// }
}

func TestGetCreator(t *testing.T) {
	doc := loadTestFile(t)

	r := doc.GetCreator()
	t.Logf("Creator: %s", r)
	// if r == "" {
	// 	t.Fail()
	// }
}

func TestGetProducer(t *testing.T) {
	doc := loadTestFile(t)

	r := doc.GetProducer()
	t.Logf("Producer: %s", r)
	// if r == "" {
	// 	t.Fail()
	// }
}

func TestGetMetadata(t *testing.T) {
	doc := loadTestFile(t)

	r := doc.GetMetadata()
	t.Logf("Metadata: %s", r)

	// Don't care about the data, just that it's valid XML
	data := struct{}{}
	if err := xml.Unmarshal([]byte(r), &data); err != nil {
		t.Logf("Metadata Error: %s", err)
		t.Fail()
	}
}

func TestGetCreatedDate(t *testing.T) {
	doc := loadTestFile(t)

	r := doc.GetCreatedDate()
	t.Logf("Created Date: %s", r)
}

func TestGetModificationDate(t *testing.T) {
	doc := loadTestFile(t)

	r := doc.GetModificationDate()
	t.Logf("Modification Date: %s", r)
}

func TestIsLinearized(t *testing.T) {
	doc := loadTestFile(t)

	r := doc.IsLinearized()
	t.Logf("Is Linearized: %v", r)
}
