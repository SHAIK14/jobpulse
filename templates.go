package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
)

type Templates struct {
	Role string `json:"role"`
	Body string `json:"body"`
}

func loadTemplates() []Templates {
	bytes, err := os.ReadFile("templates.json")
	if err != nil {
		return []Templates{}
	}
	var data []Templates
	json.Unmarshal(bytes, &data)
	return data

}
func saveTemplates(jsonData []Templates) {
	bytes, _ := json.Marshal(jsonData)
	os.WriteFile("templates.json", bytes, 0644)

}

func addTemplate(role, body string) {
	data := loadTemplates()

	for i, d := range data {
		if d.Role == role {
			data[i].Body = body
			saveTemplates(data)
			return

		}
	}
	data = append(data, Templates{Role: role, Body: body})
	saveTemplates(data)

}

func showTemplates() {
	data := loadTemplates()
	if len(data) == 0 {
		fmt.Println("add atleast one template using the 'add' cmd")
		os.Exit(1)
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "ROLE\tTEMPLATE")
	for _, d := range data {
		fmt.Fprintf(w, "%s\t%s\n", d.Role, d.Body)
	}
	w.Flush()

}

func updateTemplate(role, body string) {
	data := loadTemplates()
	for i, d := range data {
		if d.Role == role {

			data[i].Body = body
		}

	}
	saveTemplates(data)

}

func TemplatesRouter() {

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	rolePtr := addCmd.String("role", "", "type like CTO,CEO,HR")
	bodyPtr := addCmd.String("body", "", "body of the template")

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	updateRolePtr := updateCmd.String("role", "", "role to update")
	updateBodyPtr := updateCmd.String("body", "", "new body")

	if len(os.Args) < 2 {
		fmt.Println("expected 'add' or 'list' commands ")
		os.Exit(1)

	}
	switch os.Args[2] {
	case "add":
		addCmd.Parse(os.Args[3:])
		addTemplate(*rolePtr, *bodyPtr)
	case "list":
		listCmd.Parse(os.Args[3:])
		showTemplates()
	case "update":
		updateCmd.Parse(os.Args[3:])
		updateTemplate(*updateRolePtr, *updateBodyPtr)

	}

}
