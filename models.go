package main

type Bewertung struct {
	Vorname       string  `json:"vorname"`
	Nachname      string  `json:"nachname"`
	ID            int     `json:"id"`
	HvPunkte      float64 `json:"hvPunkte"`
	HvProzent     float64 `json:"hvProzent"`
	HvNote        int     `json:"hvNote"`
	LvPunkte      float64 `json:"lvPunkte"`
	LvProzent     float64 `json:"lvProzent"`
	LvNote        int     `json:"lvNote"`
	GesamtProzent float64 `json:"gesamtProzent"`
	GesamtNote    int     `json:"gesamtNote"`
	Gewertet      bool    `json:"gewertet"`
}

type MaxPunkte struct {
	HvMax        float64 `json:"hvMax"`
	LvMax        float64 `json:"lvMax"`
	HvGewichtung float64 `json:"hvGewichtung"`
	LvGewichtung float64 `json:"lvGewichtung"`
}
