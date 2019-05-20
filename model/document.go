package model

type DocumentFile struct {
	FileName     string `json:"fileName"`
	Consolidates []struct {
		Overdue           string `json:"overdue"`
		Credit            string `json:"credit"`
		Dispute           string `json:"dispute"`
		SalesOrganization string `json:"salesOrganization"`
		CustomerCode      string `json:"customerCode"`
		CustomerName      string `json:"customerName"`
		Currency          string `json:"currency"`
		Docs              []struct {
			Division          string `json:"division"`
			SalesOrganization string `json:"salesOrganization"`
			CustomerCode      string `json:"customerCode"`
			Number            string `json:"number"`
			ReferenceNumber   string `json:"referenceNumber"`
			BillingNumber     string `json:"billingNumber"`
			IsDispute         bool   `json:"isDispute"`
			IssuedDate        string `json:"issuedDate"`
			DueDate           string `json:"dueDate"`
			TotalAmount       string `json:"totalAmount"`
		} `json:"docs"`
		NotOverdue string `json:"notOverdue,omitempty"`
	} `json:"consolidates"`
}
