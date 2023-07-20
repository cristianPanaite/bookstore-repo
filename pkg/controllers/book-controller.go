package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cristianPanaite/go-bookstore/pkg/models"
	"github.com/cristianPanaite/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, &CreateBook)
	createdBook := CreateBook.CreateBook()
	res, err := json.Marshal(createdBook)

	if err != nil {
		fmt.Println("Error in marshalling")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Printf("Error while parsing")
		return
	}

	bookDetails, _ := models.GetBookById(ID)
	res, err := json.Marshal(bookDetails)

	if err != nil {
		fmt.Println("Error in marshalling")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, err := json.Marshal(books)

	if err != nil {
		fmt.Println("Error in marshalling")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Printf("Error while parsing")
		return
	}

	UpdateBook := &models.Book{}
	utils.ParseBody(r, &UpdateBook)
	updatedBook := UpdateBook.UpdateBook(ID)
	res, err := json.Marshal(updatedBook)

	if err != nil {
		fmt.Println("Error in marshalling")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Printf("Error while parsing")
		return
	}

	deletedBook := models.DeleteBook(ID)

	res, err := json.Marshal(deletedBook)
	if err != nil {
		fmt.Println("Error in marshalling")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
