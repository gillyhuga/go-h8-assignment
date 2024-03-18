package main

import (
	"fmt"
	"os"
)

type Friend struct {
	Name       string
	Address    string
	Occupation string
	Reason     string
}

var friendData = []Friend{
	{
		Name:       "Ahmad",
		Address:    "Jakarta",
		Occupation: "Programmer",
		Reason:     "Untuk meningkatkan keterampilan",
	},
	{
		Name:       "Budi",
		Address:    "Bandung",
		Occupation: "Programmer",
		Reason:     "Mendapatkan pemahaman lebih dalam tentang pemrograman",
	},
	{
		Name:       "Cici",
		Address:    "Bekasi",
		Occupation: "Programmer",
		Reason:     "Ingin belajar mengembangkan web service menggunakan Go",
	},
}

func main() {
	args := os.Args
	absentNumber := args[1]
	ShowFriendData(absentNumber)
}

func ShowFriendData(absentNumber string) {
	var absent int
	_, err := fmt.Sscanf(absentNumber, "%d", &absent)
	if err != nil {
		fmt.Println("Must be an integer")
		return
	}

	if absent < 1 || absent > len(friendData) {
		fmt.Println("No friend data")
		return
	}

	friend := friendData[absent-1]
	fmt.Println("Name:", friend.Name)
	fmt.Println("Address:", friend.Address)
	fmt.Println("Occupation:", friend.Occupation)
	fmt.Println("Reason for Choosing Golang Class:", friend.Reason)
}
