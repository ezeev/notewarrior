package search

import (
	"log"
	"strings"

	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/analysis/analyzer/keyword"
	"github.com/blevesearch/bleve/analysis/lang/en"
	"github.com/blevesearch/bleve/mapping"
)

func OpenIndex(path string) (bleve.Index, error) {
	index, err := bleve.Open(path)
	if err != nil {
		if strings.Contains(err.Error(), "path does not exist") {
			mapping, err := buildNoteIndexMapping()
			if err != nil {
				return nil, err
			}
			//mapping := bleve.NewIndexMapping()
			index, err = bleve.New(path, mapping)
			log.Printf("Created new index using path: %s\n", path)
			if err != nil {
				return nil, err
			}
		}
	}
	return index, nil
}

func buildNoteIndexMapping() (mapping.IndexMapping, error) {

	// en text analysis
	enTextFieldMapping := bleve.NewTextFieldMapping()
	enTextFieldMapping.Analyzer = en.AnalyzerName

	// keyword analysis
	keywordFieldMapping := bleve.NewTextFieldMapping()
	keywordFieldMapping.Analyzer = keyword.Name

	//add the mappings
	nwmapping := bleve.NewDocumentMapping()

	nwmapping.AddFieldMappingsAt("headings", enTextFieldMapping)
	nwmapping.AddFieldMappingsAt("body", enTextFieldMapping)
	nwmapping.AddFieldMappingsAt("tags", keywordFieldMapping)

	indexMapping := bleve.NewIndexMapping()
	indexMapping.AddDocumentMapping("note", nwmapping)

	indexMapping.TypeField = "type"
	indexMapping.DefaultAnalyzer = "en"

	return indexMapping, nil

}

func IndexNote(note *NoteDoc, index bleve.Index) error {
	return index.Index(note.Path, note)
}
