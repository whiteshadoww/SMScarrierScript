package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"io/ioutil"
)

type SMSCarrier struct {
	MCC         string `json:"MCC"`
	MNC         string `json:"MCN"`
	ISO         string `json:"ISO"`
	Country     string `json:"Country"`
	CountryCode string `json:"Country Code"`
	Network     string `json:"Network"`
}

func main() {

	var smsCarriers []SMSCarrier

	c := colly.NewCollector()
	//parse Table body that contains the data
	c.OnHTML("#mncmccTable tbody", func(e *colly.HTMLElement) {

		//loop through all rows
		e.ForEach("tr", func(i int, e *colly.HTMLElement) {

			//get data based on index of column
			smsCarriers = append(smsCarriers, SMSCarrier{
				MCC:         e.ChildText("td:nth-child(1)"),
				MNC:         e.ChildText("td:nth-child(2)"),
				ISO:         e.ChildText("td:nth-child(3)"),
				Country:     e.ChildText("td:nth-child(4)"),
				CountryCode: e.ChildText("td:nth-child(5)"),
				Network:     e.ChildText("td:nth-child(6)"),
			})
		})
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.StatusCode)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	err := c.Visit("http://mcc-mnc.com/")
	if err != nil {
		fmt.Println(err)
	}

	file, _ := json.MarshalIndent(smsCarriers, "", " ")
	_ = ioutil.WriteFile("sms_carriers.json", file, 0644)
}
