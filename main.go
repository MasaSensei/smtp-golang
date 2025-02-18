package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"server/smtp"
)

type FormData struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Interest  string `json:"interest"`
	Message   string `json:"message"`
}

func handleForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var formData FormData
	err := json.NewDecoder(r.Body).Decode(&formData)
	if err != nil {
		log.Println("Error decoding form data:", err)
		http.Error(w, "Failed to decode form data", http.StatusBadRequest)
		return
	}

	message := fmt.Sprintf(
		"New Form Submission\n\nName: %s %s\nEmail: %s\nInterest: %s\nMessage: %s",
		formData.FirstName, formData.LastName, formData.Email, formData.Interest, formData.Message,
	)

	emailTo := os.Getenv("EMAIL_TO")
	if emailTo == "" {
		log.Println("EMAIL_TO environment variable is not set")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = smtp.SendEmail(emailTo, "New Form Submission", message)
	if err != nil {
		log.Println("Error sending email:", err)
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Form submitted successfully"})
}

func main() {

	smtp.InitSMTP()

	http.HandleFunc("/submit", handleForm)

	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
