package services

import (
	"fmt"
)

type LibraryManager interface { 

	AddBook(book models.Book)
	BorrowBook(bookID, memberID int) error 
	RemoveBook(bookID int) 
	ReturnBook(bookID,memberID int)  error 
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book 
}
// this initialize the Libary contruct ot store all book and memeber
type Library struct {
	Books map[int]models.Book
	Members map[int]models.Members
}

// now let's implement the Library struct

func newLibrary() *Library { 
	return &Library{
		Books: make(map[int]models.Book)
		Members: make(map[int]models.Member)
	}
}

// now let's implement each of the mothods in the library struct
func (l *Library) AddBook(book models.Book) {
	 l.Books[book.ID] = book
}

func (l *Library) RemoveBook(bookID int) {
	delete(l.Books, book.ID)
}



// implement BorrowBook 
func (l *Library) BorrowBook(bookID, memberID int) error {
	book, ok := l.Books[bookID]
	if !ok {
		return fmt.Errorf("The book with ID %d is not found", bookID)
	}
	member, ok := l.Members[memberID]
	if !ok {
		return fmt.Errorf("The member with an ID %d is not found", memberID)
	}

	if book.Status != "Available" {
		return fmt.Errorf("The book is borrowed")


	}
	// all good, now let's borrow the book 
	book.Status = "Borrowed"
	// update the map
	l.Books[bookID] = book
	// now add it to the borrowed list 

	member.BorrowedBooks = append(member.BorrowedBooks, bookID)
	l.Members[memberID] = member
	
	return nil 
}
func (l *Library) ReturnBook(bookID, memberID int) error {
	book, ok := l.Books[bookID]
	if !ok {
		return fmt.Errorf("the book with ID %d is not found", bookID)
	}

	member, ok := l.Members[memberID]
	if !ok {
		return fmt.Errorf("the member with ID %d is not found", memberID)
	}

	if book.Status != "Borrowed" {
		return fmt.Errorf("the book is not currently borrowed")
	}

	// Mark the book as available
	book.Status = "Available"
	l.Books[bookID] = book

	// Remove bookID from member's borrowed list
	for i, id := range member.BorrowedBooks {
		if id == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			break
		}
	}

	l.Members[memberID] = member

	return nil
}

// implementing available book listing 
func (l *Library) ListAvailableBooks() []models.Book {
	available := []models.Book
	for _, book := range l.Books {
		if book.Status == "Available" {
			available = append(available, book)
		}
	}
	return available
}

// borrowed book listing 
func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, ok := l.Members[memberID]
	if !ok {
		return nil
	}

	var borrowed []models.Book
	for _, id := range member.BorrowedBooks {
		borrowed = append(borrowed, l.Books[id])
	}
	return borrowed
}


