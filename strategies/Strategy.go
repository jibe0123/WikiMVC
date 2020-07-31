package strategies

import "github.com/jibe0123/WikiMVC/models"

type ExportStrategy interface {
	// Methods for exporting articles
	Export(dataToExport models.Article) error
}

