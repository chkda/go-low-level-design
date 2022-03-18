package main

type BookReservation struct {
	CreationDate string
	Status BookStatus
	BookItemBarcode string
	MemberID string
}

type BookLending struct {
	CreationDate string
	DueDate string
	BookItemBarcode string
	MemberID string
}

type Fine struct {
	CreationDate string
	BookItemBarcode string
	MemberID string
}

