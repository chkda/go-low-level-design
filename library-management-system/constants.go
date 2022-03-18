package main

type BookFormat uint

const (
	HARDCOVER BookFormat = iota + 1
	PAPERBACK
	AUDIO_BOOK
	EBOOK
	NEWSPAPER
	MAGAZINE
	JOURNAL
)

type BookStatus uint

const (
	AVAILABLE BookStatus = iota + 1
	RESERVED
	LOANED
	LOST
)

type ReservationStatus uint

const (
	WAITING ReservationStatus = iota + 1
	PENDING
	CANCELED
	NONE
)

type AccountStatus uint

const (
	ACTIVE AccountStatus = iota + 1
	CLOSED 
	BLACKLISTED 
)

type Address struct {
	Street string
	City string
	State string
	Pincode  string
}

type Person struct {
	Name string
	Address Address
	Email string
	Phone string
}

const MAX_BOOKS_ISSUED_PER_USER uint = 5
const MAX_LENDING_DAYS uint = 10
