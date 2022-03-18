package main

type Catalog struct {
	BookTitles map[string]bool
	BookAuthors map[string][]string
}