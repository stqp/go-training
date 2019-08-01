package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

const (
	secInOneMonth = 60 * 60 * 24 * 30
	secInOneYear  = secInOneMonth * 12
)

func l(item *Issue) {
	fmt.Printf("#%-5d %s %9.9s %.55s\n", item.Number, item.CreatedAt.Format("2006-01-03"), item.User.Login, item.Title)
}

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	items := result.Items
	sort.Slice(items, func(i, j int) bool {
		return items[i].CreatedAt.After(items[j].CreatedAt)
	})

	printOneMonthBanner := true
	printOneYearBanner := true
	printOtherBanner := true

	for _, i := range items {

		if i.CreatedAt.AddDate(0, 1, 0).After(time.Now()) {
			if printOneMonthBanner {
				fmt.Println("\n==== now - createdAt < oneMonth ====")
				printOneMonthBanner = false
			}
		} else if i.CreatedAt.AddDate(1, 0, 0).After(time.Now()) {
			if printOneYearBanner {
				fmt.Println("\n==== now - createdAt < oneYear ====")
				printOneYearBanner = false
			}
		} else {
			if printOtherBanner {
				fmt.Println("\n==== now - createdAt > oneYear ====")
				printOtherBanner = false
			}
		}
		l(i)

	}
}
