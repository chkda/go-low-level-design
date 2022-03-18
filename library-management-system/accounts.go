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