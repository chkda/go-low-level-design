package main

type Account struct {
	ID string
	Password string
	AccountStatus
	Person
}

type Librarian struct {
	Account
}

type Member struct {
	Account
	MembershipDate string
	TotalBooksCheckedOut int
}

func (l Librarian) AddBook(book BookItem) bool {

}

func (l Librarian) RemoveBook(book BookItem) bool {

}

func (l Librarian) EditBook(book BookItem) bool {

}

func (l Librarian) RegisterUser(user Member) bool {

}

func (l Librarian) UpdateUser(user Member) bool {

}

func (l Librarian) BlockUser(user Member) bool {

}

func (l Librarian) CheckOutBook(book BookItem) bool {

}

func (l Librarian) ReturnBook(book BookItem) bool {

}

func (l Librarian) ReserveBook(book BookItem) BookStatus {

}

func (l Librarian) CancelReservation(book BookItem) BookStatus {

}

func (m Member) CheckOutBook(book BookItem) bool {

}

func (m Member) ReturnBook(book BookItem) bool {

}

func (m Member) ReserveBook(book BookItem) BookStatus {

}

func (m Member) CancelReservation(book  BookItem) BookStatus {

}