package dto

type Data struct {
	Humidity     string `json:"humidity"`
	Temperature  string `json:"temperature"`
	SoilMoisture string `json:"soil_moisture"`
	SoilPH       string `json:"soil_ph"`
	Gas          string `json:"gas"`
	Langitude    string `json:"langitude"`
	Latitude     string `json:"latitude"`
	Image        string `json:"image"`
}
