package strategies

import "github.com/jibe0123/WikiMVC/models"

type Context struct {
	Strategy ExportStrategy
	Article models.Article
}

func (c* Context) setStrategy(strategy ExportStrategy){
	c.Strategy = strategy
}

func (c* Context) executeStrategy(strategy ExportStrategy, article models.Article) error {
	return strategy.Export(article)
}
