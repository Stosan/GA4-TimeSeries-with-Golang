from google.analytics.data_v1beta import BetaAnalyticsDataClient
from google.analytics.data_v1beta.types import (
    DateRange,
    Dimension,
    Metric,
    RunReportRequest,
)
from google.oauth2 import service_account

def sample_run_report(property_id="YOUR-GA4-PROPERTY-ID"):
    """Runs a simple report on a Google Analytics 4 property."""
    # TODO(developer): Uncomment this variable and replace with your
    #  Google Analytics 4 property ID before running the sample.
    # property_id = "YOUR-GA4-PROPERTY-ID"
    # Create credentials object from the provided credentials path.
    credentials =service_account.Credentials.from_service_account_file("./radadspdtest-381223-0871747aa9a9.json")

    # Using a default constructor instructs the client to use the credentials
    # specified in GOOGLE_APPLICATION_CREDENTIALS environment variable.
    client = BetaAnalyticsDataClient(
    credentials=credentials
    )

    request = RunReportRequest(
        property=f"properties/{property_id}",
        metrics=[Metric(name="newUsers")],
        date_ranges=[DateRange(start_date="2023-02-21", end_date="2023-03-20")],
    )
    response = client.run_report(request)

    print("Report result:")
    for row in response.rows:
        print(row.metric_values[0].value)


sample_run_report(property_id="352702775")