package models

type Weather struct {
	Status Status
}

type Status struct {
	Water int
	Wind  int
}

type ViewWeather struct {
	Water       int
	StatusWater string
	Wind        int
	StatusWind  string
}
