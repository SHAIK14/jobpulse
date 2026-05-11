package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"
)

type Details struct {
	Id      int    `json:"id"`
	Company string `json:"company"`

	Joblink   string `json:"joblink"`
	Contact   string `json:"contact"`
	Email     string `json:"email"`
	Title     string `json:"title"`
	Status    string `json:"status"`
	Dmsent_at string `json:"dmsent_at"`
}

func loadData() []Details {
	bytes, err := os.ReadFile("jobpulse.json")
	if err != nil {
		return []Details{}
	}
	var details []Details
	json.Unmarshal(bytes, &details)
	return details

}

func saveData(jsonData []Details) {
	bytes, _ := json.Marshal(jsonData)
	os.WriteFile("jobpulse.json", bytes, 0644)

}

func addDetails(company, joblink, contact, email, title, status, dmsent_at string) {

	data := Details{
		Company:   company,
		Joblink:   joblink,
		Contact:   contact,
		Email:     email,
		Title:     title,
		Status:    status,
		Dmsent_at: dmsent_at,
	}
	jsonData := loadData()
	data.Dmsent_at = time.Now().Format("2006-01-02")
	if len(jsonData) == 0 {
		data.Id = 1
	} else {
		data.Id = jsonData[len(jsonData)-1].Id + 1
	}

	jsonData = append(jsonData, data)
	saveData(jsonData)

}

func listDetails() {
	data := loadData()
	if len(data) == 0 {
		fmt.Println("add atleast one company using the 'add' cmd")
		os.Exit(1)
	}
	for _, d := range data {
		fmt.Println(d.Id, d.Company, d.Joblink, d.Contact, d.Title, d.Status, d.Dmsent_at)
	}

}

func outreach() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	companyPtr := addCmd.String("company", "", "name of the company")
	joblinkPtr := addCmd.String("joblink", "", "link of the JD")
	contactPtr := addCmd.String("contact", "", "name or linkedin of the contact")
	emailPtr := addCmd.String("email", "", "email of the contact")
	titlePtr := addCmd.String("title", "", "job title of the contact")
	statusPtr := addCmd.String("status", "DMed", "status of the cold mail")
	dmSentAtPtr := addCmd.String("date", "", "date of the dm sent")

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("expected 'add' or 'list' commands ")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		addDetails(*companyPtr, *joblinkPtr, *contactPtr, *emailPtr, *titlePtr, *statusPtr, *dmSentAtPtr)

	case "list":
		listCmd.Parse(os.Args[2:])
		listDetails()

	default:
		fmt.Println("expected 'add' or 'list' commands ")
		os.Exit(1)

	}
}
