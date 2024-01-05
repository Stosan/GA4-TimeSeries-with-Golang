package campaigns

import (

	"log"
	"time"

	"google.golang.org/api/analyticsdata/v1beta"

)

func Traffic_AdsPresent(client *analyticsdata.Service, propertyID string,startDate string,endDate string) (map[string]string, error) {
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

	return sessionsPerDay,nil
}


func Traffic_PriorAds(client *analyticsdata.Service, propertyID string,startDate string,endDate string) (map[string]string, error) {
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

	return sessionsPerDay,nil
}


// google analytics for radioadspread
func GA4_Analytics_Summary(client *analyticsdata.Service, propertyID string,duration,durationprior map[string]string) (map[string]interface{}, error) {

	// convstartDate, err := time.Parse("2006-01-02", duration["start"])
	// if err != nil {
	// 	log.Fatalf("Failed to parse start date: %v", err)
	// }
	// duration["start"] = convstartDate.Format("2006-01-02")
	// startDate :=duration["start"]


	// convendDate, err := time.Parse("2006-01-02", duration["end"])
	// if err != nil {
	// 	log.Fatalf("Failed to parse end date: %v", err)
	// }
	// duration["end"] = convendDate.Format("2006-01-02")
	// endDate :=duration["end"]
resp,_:=Traffic_AdsPresent(client,propertyID,duration["start"],duration["end"])
resp_prior,_:=Traffic_PriorAds(client,propertyID,durationprior["start"],durationprior["end"])
web_analytics := map[string]interface{}{
	"before":resp_prior,
	"after":resp,
}
	return web_analytics,nil
}
