package utils

import (
	"crypto/rand"
	"fmt"
	"io"
	"strconv"
	"time"
)

func GeneratePassword() string {
	pass := strconv.FormatInt(time.Now().UnixMilli(), 10)
	pass = pass[0:6]
	return pass
}

func GenerateRegNo(index int) string {
	var complaintId string
	if index < 9 {
		incrementIndex := index + 1
		complaintId = fmt.Sprintf("0000%d", incrementIndex)
	}
	if index >= 9 && index < 99 {
		incrementIndex := index + 1
		complaintId = fmt.Sprintf("000%d", incrementIndex)
	}
	if index >= 99 && index < 999 {
		incrementIndex := index + 1
		complaintId = fmt.Sprintf("00%d", incrementIndex)
	}
	if index >= 999 && index < 9999 {
		incrementIndex := index + 1
		complaintId = fmt.Sprintf("0%d", incrementIndex)
	}
	if index >= 9999 {
		incrementIndex := index + 1
		complaintId = fmt.Sprintf("%d", incrementIndex)
	}

	return complaintId
}

func NumberGenerator(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = numberTable[int(b[i])%len(numberTable)]
	}
	return string(b)
}

var numberTable = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func CalculateAge(birthdate, today time.Time) int {
	today = today.In(birthdate.Location())
	ty, tm, td := today.Date()
	today = time.Date(ty, tm, td, 0, 0, 0, 0, time.UTC)
	by, bm, bd := birthdate.Date()
	birthdate = time.Date(by, bm, bd, 0, 0, 0, 0, time.UTC)
	if today.Before(birthdate) {
		return 0
	}
	age := ty - by
	anniversary := birthdate.AddDate(age, 0, 0)
	if anniversary.After(today) {
		age--
	}

	return age
}
