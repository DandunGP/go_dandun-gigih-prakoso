package main

import (
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	Id       string `json:"id" form:"id"`
	Title    string `json:"title" form:"title"`
	Price    string `json:"price" form:"price"`
	Category string `json:"category" form:"category"`
}

var users []User

func Menu() {
	fmt.Println("===BOOK MANAGEMENT===")
	fmt.Println("1. Get all books")
	fmt.Println("2. Get book by ID")
	fmt.Println("3. Create a books")
	fmt.Println("4. Update a books")
	fmt.Println("5. Delete a books")
	fmt.Println("6. Exit")
}

func GetAllBookController() {
	fmt.Println("AllBooks")
	for i := 0; i < len(users); i++ {
		fmt.Println("===")
		fmt.Println("ID : " + users[i].Id)
		fmt.Println("Title : " + users[i].Title)
		fmt.Println("Price : " + users[i].Price)
		fmt.Println("Category : " + users[i].Category)
		fmt.Println("===")
	}
}

func GetBookController() {
	var userGet []User
	var id string
	fmt.Print("Enter ID : ")
	fmt.Scan(&id)

	for i := 0; i < len(users); i++ {
		if users[i].Id == id {
			userGet = append(userGet, users[i])
		}
	}

	if userGet != nil {
		fmt.Println("Title : " + userGet[0].Title)
		fmt.Println("Price : " + userGet[0].Price)
		fmt.Println("Category : " + userGet[0].Category)
	} else {
		fmt.Println("Record not found!")
	}
}

func CreateBookController() {
	userNew := User{}
	var title, price, category string

	userNew.Id = uuid.New().String()
	fmt.Print("Enter title : ")
	fmt.Scan(&title)
	userNew.Title = title
	fmt.Print("Enter price : ")
	fmt.Scan(&price)
	userNew.Price = price
	fmt.Print("Enter category : ")
	fmt.Scan(&category)
	userNew.Category = category

	users = append(users, userNew)
}

func UpdateBookController() {
	var id, title, price, category string

	fmt.Print("Enter ID : ")
	fmt.Scan(&id)

	for i := 0; i < len(users); i++ {
		if users[i].Id == id {
			fmt.Print("Enter title: ")
			fmt.Scan(&title)
			users[i].Title = title
			fmt.Print("Enter price: ")
			fmt.Scan(&price)
			users[i].Price = price
			fmt.Print("Enter category: ")
			fmt.Scan(&category)
			users[i].Category = category

			fmt.Println("the ID : " + id)
			fmt.Println("Book updated!")
		}
	}
}

func DeleteBookController() {
	var id string

	fmt.Print("Enter ID : ")
	fmt.Scan(&id)

	for i := 0; i < len(users); i++ {
		if users[i].Id == id {
			users = append(users[:i], users[i+1:]...)
			fmt.Println("Book deleted!")
		}
	}
}

func main() {
	var choose int
	i := 0
	for i != 6 {
		Menu()
		fmt.Print("Choose your menu : ")
		fmt.Scan(&choose)
		switch choose {
		case 1:
			GetAllBookController()
		case 2:
			GetBookController()
		case 3:
			CreateBookController()
		case 4:
			UpdateBookController()
		case 5:
			DeleteBookController()
		case 6:
			i = 6
		default:
			fmt.Println("Input invalid!")
		}
	}
	fmt.Println("Byee...")
}
