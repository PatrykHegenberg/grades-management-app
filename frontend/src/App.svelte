<script>
  import { onMount } from "svelte";
  import {
    AddBewertung,
    GetBewertungen,
    GetMaxPunkte,
    SetMaxPunkte,
    ToggleWertung,
    ExportBewertungen,
    GetNotenspiegel,
  } from "../wailsjs/go/main/App";
  import { OpenSaveDialog } from "../wailsjs/go/main/App";
  import MaxPunkteForm from "./MaxPunkteForm.svelte";
  import BewertungForm from "./BewertungForm.svelte";
  import BewertungenTable from "./BewertungenTable.svelte";
  import Notenspiegel from "./Notenspiegel.svelte";
  import ExportSection from "./ExportSection.svelte";
  import ThemeSwitcher from "./ThemeSwitcher.svelte";

  let bewertungen = [];
  let maxPunkte = {
    hvMax: 0,
    lvMax: 0,
    hvGewichtung: 0,
    lvGewichtung: 0,
  };

  let vorname = "";
  let nachname = "";
  let hvPunkte = 0;
  let lvPunkte = 0;
  let hvMax = 0;
  let lvMax = 0;
  let hvGewichtung = 0;
  let lvGewichtung = 0;

  let exportPath = "";
  let notenspiegel = {};

  onMount(async () => {
    await loadData();
    await loadNotenspiegel();
  });

  async function loadData() {
    bewertungen = await GetBewertungen();
    maxPunkte = await GetMaxPunkte();
    hvMax = maxPunkte.hvMax;
    lvMax = maxPunkte.lvMax;
    hvGewichtung = maxPunkte.hvGewichtung;
    lvGewichtung = maxPunkte.lvGewichtung;
  }

  async function handleAddBewertung() {
    if (maxPunkte.hvMax !== 0 && maxPunkte.hvGewichtung !== 0) {
      const success = await AddBewertung(vorname, nachname, hvPunkte, lvPunkte);
      if (success) {
        await loadData();
        resetForm();
        await loadNotenspiegel();
      } else {
        alert("Name existiert bereits!");
      }
    } else {
      alert("Max Punkte muss befüllt sein");
    }
  }

  async function loadNotenspiegel() {
    notenspiegel = await GetNotenspiegel();
  }

  let isMaxPunkteValid = false;

  function validateMaxPunkte() {
    if (hvMax > 0 && lvMax === 0) {
      isMaxPunkteValid = hvGewichtung === 100;
    } else if (hvMax > 0 && lvMax > 0) {
      isMaxPunkteValid = hvGewichtung + lvGewichtung === 100;
    } else {
      isMaxPunkteValid = false;
    }
  }

  $: {
    validateMaxPunkte();
  }

  async function handleSetMaxPunkte() {
    validateMaxPunkte();

    if (isMaxPunkteValid) {
      const success = await SetMaxPunkte(
        hvMax,
        lvMax,
        hvGewichtung,
        lvGewichtung,
      );
      if (success) {
        await loadData();
      } else {
        alert("Es gab einen Fehler beim Setzen der Max-Punkte.");
      }
    }
  }

  async function handleToggleWertung(id) {
    await ToggleWertung(id);
    await loadData();
  }
  async function selectExportPath() {
    try {
      const result = await OpenSaveDialog();
      if (result) {
        exportPath = result;
      }
    } catch (error) {
      console.error("Fehler beim Öffnen des Dateiauswahldialogs:", error);
    }
  }

  async function handleExport() {
    if (!exportPath) {
      alert("Bitte wählen Sie einen Speicherpfad aus.");
      return;
    }
    await ExportBewertungen(exportPath);
    alert("Export abgeschlossen!");
  }

  function resetForm() {
    vorname = "";
    nachname = "";
    hvPunkte = 0;
    lvPunkte = 0;
  }
</script>

<div class="container mx-auto p-4">
  <!-- <ThemeSwitcher /> -->
  <div class="card bg-base-100 shadow-xl">
    <div class="card-body">
      <div class="flex justify-between items-center mb-6">
        <h1 class="text-3xl font-bold">Klassenarbeit-Bewertungssystem</h1>
        <ThemeSwitcher />
      </div>
      <!-- <h1 class="text-3xl font-bold text-center mb-6"> -->
      <!--   Klassenarbeit-Bewertungssystem -->
      <!-- </h1> -->

      <div class="divider"></div>
      <h3 class="text-2xl font-bold mb-4">Bewertungen</h3>

      <MaxPunkteForm
        bind:hvMax
        bind:lvMax
        bind:hvGewichtung
        bind:lvGewichtung
        onSubmit={handleSetMaxPunkte}
      />

      <BewertungForm
        bind:vorname
        bind:nachname
        bind:hvPunkte
        bind:lvPunkte
        onSubmit={handleAddBewertung}
        disabled={!isMaxPunkteValid}
      />

      <BewertungenTable {bewertungen} onToggleWertung={handleToggleWertung} />

      <Notenspiegel {notenspiegel} />

      <ExportSection
        bind:exportPath
        onSelectPath={selectExportPath}
        onExport={handleExport}
      />
    </div>
  </div>
</div>
