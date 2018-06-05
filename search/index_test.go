package search

import (
	"testing"

	"github.com/blevesearch/bleve"
)

func TestIndexNote(t *testing.T) {

	idx, err := OpenIndex("test.bleve")
	//defer idx.Close()
	if err != nil {
		t.Error(err)
	}

	note := NoteDoc{
		Path:     "/tmp/test.md",
		Body:     "this is test body",
		Type:     "note",
		Tags:     []string{"tag1", "tag2", "tag3"},
		Headings: make([]string, 3),
	}
	note.Headings[0] = "heading 1"
	note.Headings[1] = "heading 2"
	note.Headings[2] = "heading 3"

	err = IndexNote(&note, idx)
	if err != nil {
		t.Error(err)
	}
	//now query it
	query := bleve.NewQueryStringQuery("test")
	searchRequest := bleve.NewSearchRequest(query)
	searchRequest.Fields = []string{"headings", "body"}
	searchRequest.AddFacet("tags", bleve.NewFacetRequest("tags", 10))
	searchResult, err := idx.Search(searchRequest)
	if err != nil {
		t.Error(err)
	}
	for _, v := range searchResult.Hits {
		t.Log(v.ID)
		for _, f := range v.Fields {
			t.Logf("%s", f)
		}
	}
	for _, v := range searchResult.Facets {
		t.Logf("facet: %s\n", v.Field)
		for _, f := range v.Terms {
			t.Logf("\t%s : %d\n", f.Term, f.Count)
		}
	}

}
