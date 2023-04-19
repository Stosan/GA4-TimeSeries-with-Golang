package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/api/analyticsdata/v1beta"
	"google.golang.org/api/option"
)

func main() {


	ctx := context.Background()

	// Replace with your private key file name and property ID.
	privateKey := "radadspdtest-381223-0871747aa9a9.json"
	propertyID := "352702775"

	// Create a new analyticsdata client.
	client, err := analyticsdata.NewService(ctx, option.WithCredentialsFile(privateKey))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Set the date range for the report.
	startDate := "2023-02-21"
	endDate := "2023-03-20"

	// Create a request to get sessions per day.
	request := &analyticsdata.RunReportRequest{
		Property: "properties/" + propertyID,
		DateRanges: []*analyticsdata.DateRange{
			{
				StartDate: startDate,
				EndDate:   endDate,
			},
		},
		Dimensions: []*analyticsdata.Dimension{
			{
				Name: "date",
			},
		},
		Metrics: []*analyticsdata.Metric{
			{
				Name: "sessions",
			},
		},
	}

	// Run the report and print the result.
	response, err := client.Properties.RunReport("properties/"+propertyID, request).Do()
	if err != nil {
		log.Fatalf("Failed to get report: %v", err)
	}

	sessionsPerDay := make(map[string]string)
	for _, row := range response.Rows {
		dateStr := row.DimensionValues[0].Value
		date, err := time.Parse("20060102", dateStr)
		if err != nil {
			log.Fatalf("Failed to parse date: %v", err)
		}
		sessions := row.MetricValues[0]

		sessionsPerDay[date.Format("2006-01-02")] = sessions.Value
	}

	fmt.Printf("%v\n", sessionsPerDay)
}
