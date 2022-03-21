package main

type Catalog struct {
	CreationDate string
	BookTitles map[string]bool
	BookAuthors map[string][]string
	BookPubDate map[string][]string
	BookBySubject map[string][]string
}

func (c Catalog) UpdateCatalog(b BookItem) bool {

}

func (c Catalog) SearchByTitle(title string) string {

}

func (c Catalog) SearchByAuthor(author string) []string {

}

func (c Catalog) SearchBySubject(subject string) []string {

}

func (c Catalog) SearchByPubDate(date string) []string {

}

type Search interface {
	SearchByAuthor(author string) []string
	SearchBySubject(subject string) []string
	SearchByTitle(title string) string
	SearchByPubDate(date string) []string 
}

func QueryEngine() Search {
	return 
}