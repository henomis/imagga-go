package response

import "io"

type Usage struct {
	Status `json:"status"`
	Result UsageResult `json:"result"`
}

type UsageResult struct {
	BillingPeriodEnd   string         `json:"billing_period_end"`
	BillingPeriodStart string         `json:"billing_period_start"`
	Concurrency        Concurrency    `json:"concurrency"`
	Daily              map[string]int `json:"daily"`
	DailyFor           string         `json:"daily_for"`
	DailyProcessed     int            `json:"daily_processed"`
	DailyRequests      int            `json:"daily_requests"`
	Monthly            map[string]int `json:"monthly"`
	LastUsage          int            `json:"last_usage"`
	MonthlyLimit       int            `json:"monthly_limit"`
	MonthlyProcessed   int            `json:"monthly_processed"`
	MonthlyRequests    int            `json:"monthly_requests"`
	TotalProcessed     int            `json:"total_processed"`
	TotalRequests      int            `json:"total_requests"`
	Weekly             map[string]int `json:"weekly"`
	WeeklyProcessed    int            `json:"weekly_processed"`
	WeeklyRequests     int            `json:"weekly_requests"`
}

type Concurrency struct {
	Max int `json:"max"`
	Now int `json:"now"`
}

func (u *Usage) Decode(body io.ReadCloser) error {
	return decode(body, u)
}

func (u *Usage) SetBody(body io.ReadCloser) {}
