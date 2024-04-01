package main

type Sensor struct {
	Sensor    string `json:"sensor"`
	Tipo      string `json:"tipo"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

type Alert struct {
	Alert     string `json:"alert"`
	Tipo      string `json:"tipo"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}