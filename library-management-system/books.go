package main

import (
	"fmt"
)

type Book struct {
	ISBN string
	Authors []string
	Title string
	Subject string
	Publisher string
	Pages string
	Language string
}

type BookItem struct {
	Book
	BarCode string
	RefOnly bool
	Borrowed bool
	DueDate string
	Price float32
	Format BookFormat
	Status BookStatus
	PurchaseDate string
	PublicationDate string
	PlacedAt Rack
}

type Rack struct {
	Number uint
	LocationID string
}

func (b BookItem) CheckOut() bool {
}