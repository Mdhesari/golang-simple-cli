package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

const StorageFileName = "Storage.xlsx"
const sheetName = "Sheet1"

type Agency struct {
	ID             int
	Name           string
	Address        string
	Telphone       string
	SubmissionDate string
	EmployeesCount int
	Region         string
}

type Option struct {
	ID     int
	Region string
}

func main() {
	var command string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please enter the command: ")
	scanner.Scan()
	command = scanner.Text()

	switch command {
	case "get":
		getAgency()
	case "create":
		createAgency()
	case "list":
		listAgencies()
	case "update":
		updateAgency()
	case "status":
		getStatus()
	}
}

func interfaceArrayToStringArray(input []interface{}) []string {
	output := make([]string, len(input))
	for i, item := range input {
		if str, ok := item.(string); ok {
			output[i] = str
		} else {
			output[i] = "N/A"
		}
	}
	return output
}

func ToAlphaString(num int) string {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if num >= 0 && num < len(alphabet) {
		return string(alphabet[num])
	}
	return ""
}

func openOrCreateFile() *excelize.File {
	f, err := excelize.OpenFile(StorageFileName)
	if err != nil {
		f = excelize.NewFile()
	}
	return f
}

func addToDataStorage(answers []interface{}) {
	f := openOrCreateFile()
	newRowIndex := answers[0].(int)

	for colIdx, cellValue := range answers {
		colName := ToAlphaString(colIdx)
		cellRef := colName + fmt.Sprint(newRowIndex)
		f.SetCellValue(sheetName, cellRef, cellValue)
	}

	if err := f.SaveAs(StorageFileName); err != nil {
		panic(err)
	}
}

func getDataStorage() ([][]string, error) {
	f, err := excelize.OpenFile(StorageFileName)
	if err != nil {
		f = excelize.NewFile()
	}

	rows, err := f.GetRows(sheetName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if err := f.Close(); err != nil {
		fmt.Println(err)
	}

	return rows, nil
}

func getAgencies(opt Option) []Agency {
	rows, err := getDataStorage()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var agencies []Agency

	for _, row := range rows {
		if opt.Region != "" && strings.ToLower(row[6]) != strings.ToLower(opt.Region) {
			continue
		}

		if opt.ID != 0 && toInt(row[0]) != opt.ID {
			continue
		}

		agencies = append(agencies, Agency{
			ID:             toInt(row[0]),
			Name:           row[1],
			Address:        row[2],
			Telphone:       row[3],
			SubmissionDate: row[4],
			EmployeesCount: toInt(row[5]),
			Region:         row[6],
		})

		if opt.ID != 0 {
			break
		}
	}

	return agencies
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func askQuestions() []interface{} {
	scanner := bufio.NewScanner(os.Stdin)

	questions := []string{"Name", "Address", "Telephone", "SubmissionDate", "EmployeesCount", "Region"}
	answers := make([]interface{}, len(questions))

	for i, field := range questions {
		fmt.Printf("Please enter the %s: ", field)
		scanner.Scan()
		answers[i] = scanner.Text()
	}

	return answers
}

func createAgency() {
	answers := askQuestions()

	rows, err := getDataStorage()
	if err != nil {
		panic(err)
	}

	answers = append([]interface{}{len(rows) + 1}, answers...)

	addToDataStorage(answers)
	fmt.Println("Your agency is created!")
}

func listAgencies() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter the region: ")
	scanner.Scan()
	region := scanner.Text()

	agencies := getAgencies(Option{Region: region})
	fmt.Println(agencies)
}

func getAgency() {
	id := askAgencyID()
	res := getAgencies(Option{ID: id})

	if len(res) > 0 {
		fmt.Println(res[0])
		return
	}

	fmt.Printf("There is no agency with %d id", id)
}

func askAgencyID() int {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter the agency id: ")
	scanner.Scan()
	return toInt(scanner.Text())
}

func updateAgency() {
	agencyID := askAgencyID()
	agencies := getAgencies(Option{ID: agencyID})

	if len(agencies) < 1 {
		fmt.Println("Invalid Agency ID!")
		os.Exit(1)
	}

	fmt.Printf("Your agency is %v\n", agencies[0])
	answers := askQuestions()
	answers = append([]interface{}{agencyID}, answers...)
	addToDataStorage(answers)
	fmt.Println("Your agency is updated!")
}

func getStatus() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter the region: ")
	scanner.Scan()
	region := scanner.Text()
	agencies := getAgencies(Option{Region: region})
	employeesCount := 0
	agenciesCount := len(agencies)

	for _, agency := range agencies {
		employeesCount += agency.EmployeesCount
	}

	fmt.Printf("There are %d employees and %d agencies in %s region.\n", employeesCount, agenciesCount, region)
}
