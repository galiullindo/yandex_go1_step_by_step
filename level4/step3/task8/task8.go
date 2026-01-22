package main

import "math/rand/v2"

// Агрегация.
// При агрегации один объект содержит в себе другой как свою часть.
// Однако внутренний объект не абсолютно зависим от внешнего.
// Внутренний объект может существовать и сам по себе, без внешнего.

// Композиция.
// При композиции один объект - это часть другого и не может существовать сам по себе.

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

type Report struct {
	ReportID int
	Date     string
	User
}

var ids = make(map[int]struct{})

func getUniqueId() int {
	for {
		id := rand.Int()
		if _, found := ids[id]; !found {
			return id
		}
	}
}

func CreateReport(user User, reportDate string) Report {
	return Report{ReportID: getUniqueId(), Date: reportDate, User: user}
}

func GenerateUserReports(users []User, reportDate string) []Report {
	reports := make([]Report, 0)
	for _, user := range users {
		report := CreateReport(user, reportDate)
		reports = append(reports, report)
	}
	return reports
}
