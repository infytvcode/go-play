package main

type Dsp struct {
	AvgTraffic        int         `json:"avg_traffic"`
	BidRate           int         `json:"bid_rate"`
	Dsp               string      `json:"dsp"`
	ECpm              int         `json:"e_cpm"`
	Earning           string      `json:"earning"`
	Excluded          bool        `json:"excluded"`
	ID                int         `json:"id"`
	Impression        string      `json:"impression"`
	NoBid             string      `json:"no_bid"`
	Payout            string      `json:"payout"`
	Requests          string      `json:"requests"`
	Response          string      `json:"response"`
	ServingRate       int         `json:"serving_rate"`
	Spent             int         `json:"spent"`
	Ssp               string      `json:"ssp"`
	TimeOut           string      `json:"time_out"`
	TrafficMultiplier int         `json:"traffic_multiplier"`
	Type              string      `json:"type"`
	ViewRate          int         `json:"view_rate"`
	Win               string      `json:"win"`
	WinRate           int         `json:"win_rate"`
	Publishers        []Publisher `json:"publishers"`
}

type Publisher struct {
	AvgTraffic        int    `json:"avg_traffic"`
	BidRate           int    `json:"bid_rate"`
	Dsp               string `json:"dsp"`
	Dsps              []Dsp  `json:"dsps"`
	ECpm              int    `json:"e_cpm"`
	Earning           int    `json:"earning"`
	Excluded          bool   `json:"excluded"`
	ID                int    `json:"id"`
	Impression        string `json:"impression"`
	NoBid             string `json:"no_bid"`
	Payout            int    `json:"payout"`
	Requests          string `json:"requests"`
	Response          string `json:"response"`
	ServingRate       int    `json:"serving_rate"`
	Spent             string `json:"spent"`
	Ssp               string `json:"ssp"`
	TimeOut           string `json:"time_out"`
	TrafficMultiplier string `json:"traffic_multiplier"`
	Type              string `json:"type"`
	ViewRate          int    `json:"view_rate"`
	Win               string `json:"win"`
	WinRate           int    `json:"win_rate"`
}

type Overview struct {
	Label string `json:"label"`
	Value int    `json:"value"`
}

type Data struct {
	Data     []struct{} `json:"data"`
	Overview []Overview `json:"overview"`
}
