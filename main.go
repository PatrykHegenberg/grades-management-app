// main.go
package main

import (
	"context"
	"embed"
	"fmt"
	"strconv"

	"github.com/jung-kurt/gofpdf"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

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

type App struct {
	ctx         context.Context
	bewertungen []Bewertung
	maxPunkte   MaxPunkte
}

func NewApp() *App {
	return &App{
		bewertungen: make([]Bewertung, 0),
		maxPunkte: MaxPunkte{
			HvMax:        0.00,
			HvGewichtung: 0.00,
			LvMax:        0.00,
			LvGewichtung: 0.00,
		},
	}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetBewertungen() []Bewertung {
	return a.bewertungen
}

func (a *App) GetMaxPunkte() MaxPunkte {
	return a.maxPunkte
}

func (a *App) ToggleWertung(id int) Bewertung {
	var updatedBewertung Bewertung
	for i, bewertung := range a.bewertungen {
		if bewertung.ID == id {
			a.bewertungen[i].Gewertet = !bewertung.Gewertet
			updatedBewertung = a.bewertungen[i]
			break
		}
	}
	return updatedBewertung
}

func (a *App) AddBewertung(vorname, nachname string, hvPunkte, lvPunkte float64) bool {
	if !a.validateName(vorname, nachname) {
		return false
	}

	hvProzent := 100.00 / a.maxPunkte.HvMax * hvPunkte
	lvProzent := 100.00 / a.maxPunkte.LvMax * lvPunkte
	hvNote := setNote(hvProzent)
	lvNote := setNote(lvProzent)
	gesamtProzent := hvProzent*a.maxPunkte.HvGewichtung/100 + lvProzent*a.maxPunkte.LvGewichtung/100
	gesamtNote := setNote(gesamtProzent)

	bewertung := Bewertung{
		ID:            len(a.bewertungen) + 1,
		Vorname:       vorname,
		Nachname:      nachname,
		HvPunkte:      hvPunkte,
		HvProzent:     hvProzent,
		HvNote:        int(hvNote),
		LvPunkte:      lvPunkte,
		LvProzent:     lvProzent,
		LvNote:        int(lvNote),
		GesamtProzent: gesamtProzent,
		GesamtNote:    int(gesamtNote),
		Gewertet:      true,
	}

	a.bewertungen = append(a.bewertungen, bewertung)
	return true
}

func (a *App) SetMaxPunkte(hvMax, lvMax, hvGewichtung, lvGewichtung float64) bool {
	if !checkGewichtung(lvGewichtung, hvGewichtung) {
		return false
	}

	a.maxPunkte = MaxPunkte{
		HvMax:        hvMax,
		LvMax:        lvMax,
		HvGewichtung: hvGewichtung,
		LvGewichtung: lvGewichtung,
	}
	return true
}

func (a *App) ExportBewertungen(path string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(27, 10, "Vorname", "1", 0, "", false, 0, "")
	pdf.CellFormat(27, 10, "Nachname", "1", 0, "", false, 0, "")
	pdf.CellFormat(27, 10, "HV-Punkte", "1", 0, "", false, 0, "")
	pdf.CellFormat(27, 10, "HV-Note", "1", 0, "", false, 0, "")
	pdf.CellFormat(27, 10, "LV-Punkte", "1", 0, "", false, 0, "")
	pdf.CellFormat(27, 10, "LV-Note", "1", 0, "", false, 0, "")
	pdf.CellFormat(27, 10, "Gesamtnote", "1", 0, "", false, 0, "")
	pdf.Ln(-1)

	pdf.SetFont("Arial", "", 11)
	for _, bewertung := range a.bewertungen {
		pdf.CellFormat(27, 10, bewertung.Vorname, "1", 0, "", false, 0, "")
		pdf.CellFormat(27, 10, bewertung.Nachname, "1", 0, "", false, 0, "")
		pdf.CellFormat(27, 10, strconv.FormatFloat(bewertung.HvPunkte, 'f', 2, 64), "1", 0, "", false, 0, "")
		pdf.CellFormat(27, 10, strconv.FormatInt(int64(bewertung.HvNote), 10), "1", 0, "", false, 0, "")
		pdf.CellFormat(27, 10, strconv.FormatFloat(bewertung.LvPunkte, 'f', 2, 64), "1", 0, "", false, 0, "")
		pdf.CellFormat(27, 10, strconv.FormatInt(int64(bewertung.LvNote), 10), "1", 0, "", false, 0, "")
		pdf.CellFormat(27, 10, strconv.FormatInt(int64(bewertung.GesamtNote), 10), "1", 0, "", false, 0, "")
		pdf.Ln(-1)
	}

	err := pdf.OutputFileAndClose(path)
	if err != nil {
		fmt.Println("Fehler beim Exportieren der Bewertungen:", err)
		return err
	}

	runtime.EventsEmit(a.ctx, "export-complete")
	return nil
}

func (a *App) GetNotenspiegel() map[int]int {
	notenspiegel := make(map[int]int)

	for _, bewertung := range a.bewertungen {
		notenspiegel[bewertung.GesamtNote]++
	}

	return notenspiegel
}

func (a *App) validateName(vorname, nachname string) bool {
	for _, bewertung := range a.bewertungen {
		if bewertung.Nachname == nachname && bewertung.Vorname == vorname {
			return false
		}
	}
	return true
}

func setNote(prozent float64) float64 {
	switch {
	case prozent <= 22:
		return 6.00
	case prozent <= 49:
		return 5.00
	case prozent <= 64:
		return 4.00
	case prozent <= 79:
		return 3.00
	case prozent <= 94:
		return 2.00
	default:
		return 1.00
	}
}

func checkGewichtung(lv, hv float64) bool {
	sum := hv/100 + lv/100
	return sum == 1
}

func (a *App) OpenSaveDialog() (string, error) {
	return runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		DefaultFilename: "bewertungen.pdf",
		Filters: []runtime.FileFilter{
			{DisplayName: "PDF Files (*.pdf)", Pattern: "*.pdf"},
		},
	})
}

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := NewApp()

	err := wails.Run(&options.App{
		Title:     "Bewertungen",
		Width:     1024,
		Height:    768,
		MinWidth:  1024,
		MinHeight: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})
	if err != nil {
		fmt.Println("Error:", err)
	}
}
