package main

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
)

func handleForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")
	email := r.FormValue("email")
	interest := r.FormValue("interest")

	message := fmt.Sprintf("New Form Submission\n\nName: %s %s\nEmail: %s\nInterest: %s", firstName, lastName, email, interest)

	err := smtp.SendMail("your-email@gmail.com", "New Form Submission", message)
	if err != nil {
		log.Println("Error sending email:", err)
		http.Error(w, "Error sending email", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Form submitted successfully"))
}

func main() {
	http.HandleFunc("/submit", handleForm)

	fmt.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
