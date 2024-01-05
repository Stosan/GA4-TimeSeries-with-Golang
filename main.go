package main

import (
	"context"
	campaigns "ga4test/campaign"
	"log"
	"time"

	"google.golang.org/api/analyticsdata/v1beta"
	"google.golang.org/api/option"
)

func main() {

	ctx := context.Background()

	// Replace with your private key file name and property ID.
	privateKey := "radsp_json.json"
	propertyID := "352702775"
	// Create a new analyticsdata client.
	client, err := analyticsdata.NewService(ctx, option.WithCredentialsFile(privateKey))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	var durCurrent = map[string]string{
		"start": "2023-11-28",
		"end":   "2023-12-05",
	}
	
	// Compute durPrior dynamically
	startDate, err := time.Parse("2006-01-02", durCurrent["start"])
	if err != nil {
		log.Fatalf("Failed to parse start date: %v", err)
	}
	
	endDate, err := time.Parse("2006-01-02", durCurrent["end"])
	if err != nil {
		log.Fatalf("Failed to parse end date: %v", err)
	}
	
	lapsed_duration := endDate.Sub(startDate).Hours() / 24
	log.Printf("Duration between start and end date: %v days", int(lapsed_duration))
	
	startPrior := startDate.AddDate(0, 0, -int(lapsed_duration)).Format("2006-01-02")
	endPrior := endDate.AddDate(0, 0, -int(lapsed_duration)).Format("2006-01-02")
	
	var durPrior = map[string]string{
		"start": startPrior,
		"end":   endPrior,
	}

	
	resp, _ := campaigns.GA4_Analytics_Summary(client, propertyID,durCurrent,durPrior)
	log.Println(resp)
}
