package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/user"
	"text/tabwriter"
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
	usr, _ := user.Current()

	bytes, err := os.ReadFile(usr.HomeDir + "/.jobpulse.json")

	if err != nil {
		return []Details{}
	}
	var details []Details
	json.Unmarshal(bytes, &details)
	return details

}

func saveData(jsonData []Details) {
	usr, _ := user.Current()
	bytes, _ := json.Marshal(jsonData)

	os.WriteFile(usr.HomeDir+"/.jobpulse.json", bytes, 0644)

}

func turnicate(s string, n int) string {
	if len(s) > n {
		return s[:n] + "..."
	}
	return s
}

func addDetails(company, joblink, contact, email, title, status, dmsent_at string) {

	if company == "" || joblink == "" || contact == "" {
		fmt.Println("company, joblink and contact are required")
		os.Exit(1)
	}

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
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tCOMPANY\tJOBLINK\tCONTACT\tTITLE\tSTATUS\tDATE")
	for _, d := range data {

		fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\t%s\t%s\n", d.Id, d.Company, turnicate(d.Joblink, 30), turnicate(d.Contact, 25), d.Title, d.Status, d.Dmsent_at)
	}
	w.Flush()

}

func updateDetails(id int, status string) {
	data := loadData()
	if len(data) == 0 {
		fmt.Println("add atleast one company using the 'add' cmd")
		os.Exit(1)
	}
	for i, d := range data {
		if d.Id == id {
			data[i].Status = status
		}
	}
	saveData(data)

}
func deleteDetail(id int) {
	data := loadData()
	if len(data) == 0 {
		fmt.Println("add atleast one company using the 'add' cmd")
		os.Exit(1)
	}
	newData := []Details{}
	for _, d := range data {
		if d.Id != id {
			newData = append(newData, d)
		}
	}
	saveData(newData)
	fmt.Println("deleted")

}
func showDetails(id int) {
	data := loadData()
	for _, d := range data {
		if d.Id == id {
			fmt.Printf("ID:      %d\n", d.Id)
			fmt.Printf("Company: %s\n", d.Company)
			fmt.Printf("Joblink: %s\n", d.Joblink)
			fmt.Printf("Contact: %s\n", d.Contact)
			fmt.Printf("Title:   %s\n", d.Title)
			fmt.Printf("Status:  %s\n", d.Status)
			fmt.Printf("Date:    %s\n", d.Dmsent_at)

		}
	}
}

func followUp() {
	data := loadData()
	count := 0
	printed := false
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	for _, d := range data {
		if d.Status == "DMed" {
			dmDate, _ := time.Parse("2006-01-02", d.Dmsent_at)
			days := time.Since(dmDate).Hours() / 24

			if days >= 3 {
				if !printed {

					fmt.Fprintln(w, "ID\tCOMPANY\tJOBLINK\tCONTACT\tTITLE\tSTATUS\tDATE")
					printed = true
				}

				fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\t%s\t%s\n", d.Id, d.Company, d.Joblink, d.Contact, d.Title, d.Status, d.Dmsent_at)
				count++

			}

		}

	}
	w.Flush()
	if count == 0 {
		fmt.Println("nothing needs followup today")
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

	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	idPtr := updateCmd.Int("id", 0, "id of the outreach")
	statusUpdatePtr := updateCmd.String("status", "", "status of the outreach")

	deleteCmd := flag.NewFlagSet("del", flag.ExitOnError)
	delId := deleteCmd.Int("id", 0, "id of the outreact ")

	followupCmd := flag.NewFlagSet("followup", flag.ExitOnError)

	showCmd := flag.NewFlagSet("show", flag.ExitOnError)
	showIdPtr := showCmd.Int("id", 0, "id of the outreach")

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
	case "update":
		updateCmd.Parse(os.Args[2:])
		updateDetails(*idPtr, *statusUpdatePtr)
	case "del":
		deleteCmd.Parse(os.Args[2:])
		deleteDetail(*delId)
	case "followup":
		followupCmd.Parse(os.Args[2:])
		followUp()
	case "template":
		TemplatesRouter()
	case "show":
		showCmd.Parse(os.Args[2:])
		showDetails(*showIdPtr)

	default:
		fmt.Println("expected 'add' or 'list' commands ")
		os.Exit(1)

	}
}
