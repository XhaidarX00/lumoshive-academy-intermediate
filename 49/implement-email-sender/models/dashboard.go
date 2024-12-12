package models

type Summary struct {
	TotalUser    int
	TotalSales   float64
	TotalOrder   int
	TotalProduct int
}

type Revenue struct {
	Month        string
	TotalEarning float64
}
