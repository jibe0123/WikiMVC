package strategies

import "github.com/jibe0123/WikiMVC/models"

type CsvFile struct{
	Name string
}

func (c *CsvFile) Export(dataToExport models.Article) error {
	println(dataToExport.Title)
	// do logic for export
	return nil
}