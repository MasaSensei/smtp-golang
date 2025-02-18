package main

import (
	"fmt"
	"log"
	"net/http"
	"server/smtp"
)

func handleForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Ambil data dari form
	firstName := r.FormValue("first-name")
	lastName := r.FormValue("last-name")
	email := r.FormValue("email")
	interest := r.FormValue("type")

	// Format isi email
	message := fmt.Sprintf("New Form Submission\n\nName: %s %s\nEmail: %s\nInterest: %s", firstName, lastName, email, interest)

	// Kirim email
	err = smtp.SendEmail("your-email@gmail.com", "New Form Submission", message)
	if err != nil {
		log.Println("Error sending email:", err)
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Form submitted successfully!"))
}

func main() {
	http.HandleFunc("/submit", handleForm)

	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
