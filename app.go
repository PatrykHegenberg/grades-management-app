package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/jung-kurt/gofpdf"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

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

func (a *App) OpenSaveDialog() (string, error) {
	return runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		DefaultFilename: "bewertungen.pdf",
		Filters: []runtime.FileFilter{
			{DisplayName: "PDF Files (*.pdf)", Pattern: "*.pdf"},
		},
	})
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

	// Bewertungen exportieren
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(0, 10, "Bewertungen", "", 1, "C", false, 0, "")
	pdf.Ln(5)

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
		if bewertung.Gewertet {
			pdf.CellFormat(27, 10, bewertung.Vorname, "1", 0, "", false, 0, "")
			pdf.CellFormat(27, 10, bewertung.Nachname, "1", 0, "", false, 0, "")
			pdf.CellFormat(27, 10, strconv.FormatFloat(bewertung.HvPunkte, 'f', 2, 64), "1", 0, "", false, 0, "")
			pdf.CellFormat(27, 10, strconv.FormatInt(int64(bewertung.HvNote), 10), "1", 0, "", false, 0, "")
			pdf.CellFormat(27, 10, strconv.FormatFloat(bewertung.LvPunkte, 'f', 2, 64), "1", 0, "", false, 0, "")
			pdf.CellFormat(27, 10, strconv.FormatInt(int64(bewertung.LvNote), 10), "1", 0, "", false, 0, "")
			pdf.CellFormat(27, 10, strconv.FormatInt(int64(bewertung.GesamtNote), 10), "1", 0, "", false, 0, "")
			pdf.Ln(-1)
		}
	}

	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(0, 10, "Notenspiegel", "", 1, "C", false, 0, "")
	pdf.Ln(5)

	notenspiegel := a.GetNotenspiegel()

	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(30, 10, "Note", "1", 0, "", false, 0, "")
	pdf.CellFormat(30, 10, "Anzahl", "1", 0, "", false, 0, "")
	pdf.Ln(-1)

	pdf.SetFont("Arial", "", 11)
	for note := 1; note <= 6; note++ {
		anzahl := notenspiegel[note]
		pdf.CellFormat(30, 10, strconv.Itoa(note), "1", 0, "", false, 0, "")
		pdf.CellFormat(30, 10, strconv.Itoa(anzahl), "1", 0, "", false, 0, "")
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
