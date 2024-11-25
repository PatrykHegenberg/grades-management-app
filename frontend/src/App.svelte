<script>
  import { onMount } from 'svelte';
  import { AddBewertung, GetBewertungen, GetMaxPunkte, SetMaxPunkte, ToggleWertung, ExportBewertungen, GetNotenspiegel } from '../wailsjs/go/main/App';
  import { OpenSaveDialog } from '../wailsjs/go/main/App';
  import 'bulma/css/bulma.min.css';

  let bewertungen = [];
  let maxPunkte = {
    hvMax: 0,
    lvMax: 0,
    hvGewichtung: 0,
    lvGewichtung: 0
  };
  
  let vorname = '';
  let nachname = '';
  let hvPunkte = 0;
  let lvPunkte = 0;
  let hvMax = 0;
  let lvMax = 0;
  let hvGewichtung = 0;
  let lvGewichtung = 0;

  let exportPath = '';
  let notenspiegel = {}

  onMount(async () => {
    await loadData();
    await loadNotenspiegel();
  });

  async function loadData() {
    bewertungen = await GetBewertungen();
    maxPunkte = await GetMaxPunkte();
    if (maxPunkte.hvMax === 0) {
      hvMax = maxPunkte.hvMax;
      lvMax = maxPunkte.lvMax;
      hvGewichtung = maxPunkte.hvGewichtung;
      lvGewichtung = maxPunkte.lvGewichtung;
    }
  }

  async function handleAddBewertung() {
    const success = await AddBewertung(vorname, nachname, hvPunkte, lvPunkte);
    if (success) {
      await loadData();
      resetForm();
      await loadNotenspiegel();
    } else {
      alert('Name existiert bereits!');
    }
  }

  async function loadNotenspiegel() {
    notenspiegel = await GetNotenspiegel();
  }

  async function handleSetMaxPunkte() {
    const success = await SetMaxPunkte(hvMax, lvMax, hvGewichtung, lvGewichtung);
    if (!success) {
      alert('Ungültige Gewichtung! Die Summe muss 100% ergeben.');
    } else {
      await loadData();
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
      console.error('Fehler beim Öffnen des Dateiauswahldialogs:', error);
    }
  }

  async function handleExport() {
    if (!exportPath) {
      alert('Bitte wählen Sie einen Speicherpfad aus.');
      return;
    }
    await ExportBewertungen(exportPath);
    alert('Export abgeschlossen!');
  }

  function resetForm() {
    vorname = '';
    nachname = '';
    hvPunkte = 0;
    lvPunkte = 0;
  }
</script>

<div class="container is-widescreen">
  <div class="card">
    <header class="card-header">
      <p class="card-header-title">Arbeit</p>
    </header>

    <div class="card-content">
      <div class="content">
        <h1 class="title">Bewertungen</h1>

        {#if maxPunkte.hvMax === 0}
          <form on:submit|preventDefault={handleSetMaxPunkte}>
            <div class="field is-grouped">
              <div class="control">
                <input class="input" type="number" bind:value={hvMax} placeholder="HV-Max-Punkte" required>
              </div>
              <div class="control">
                <input class="input" type="number" bind:value={hvGewichtung} placeholder="HV-Gewichtung in %" required>
              </div>
              <div class="control">
                <input class="input" type="number" bind:value={lvMax} placeholder="LV-Max-Punkte" required>
              </div>
              <div class="control">
                <input class="input" type="number" bind:value={lvGewichtung} placeholder="LV-Gewichtung in %" required>
              </div>
              <div class="control">
                <button type="submit" class="button is-primary">Setzen</button>
              </div>
            </div>
          </form>
        {/if}
        <form on:submit|preventDefault={handleAddBewertung}>
          <div class="field is-grouped">
            <div class="control">
              <input class="input" type="text" bind:value={vorname} placeholder="Vorname" required>
            </div>
            <div class="control">
              <input class="input" type="text" bind:value={nachname} placeholder="Nachname" required>
            </div>
            <div class="control">
              <input class="input" type="number" bind:value={hvPunkte} placeholder="HV-Punkte" required>
            </div>
            <div class="control">
              <input class="input" type="number" bind:value={lvPunkte} placeholder="LV-Punkte" required>
            </div>
            <div class="control">
              <button type="submit" class="button is-primary">Hinzufügen</button>
            </div>
          </div>
        </form>

        <div class="table-container">
          <table class="table is-hoverable is-fullwidth">
            <thead>
              <tr>
                <th>Gewertet</th>
                <th>Vorname</th>
                <th>Nachname</th>
                <th>HV-Punkte</th>
                <th>HV-Prozent</th>
                <th>HV-Note</th>
                <th>LV-Punkte</th>
                <th>LV-Prozent</th>
                <th>LV-Note</th>
                <th>Gesamt-Prozent</th>
                <th>Gesamt-Note</th>
              </tr>
            </thead>
            <tbody>
              {#each bewertungen as bewertung}
                <tr>
                  <td><input type="checkbox" checked={bewertung.gewertet} on:change={() => handleToggleWertung(bewertung.id)}></td>
                  <td>{bewertung.vorname}</td>
                  <td>{bewertung.nachname}</td>
                  <td>{bewertung.hvPunkte.toFixed(2)}</td>
                  <td>{bewertung.hvProzent.toFixed(2)}</td>
                  <td>{bewertung.hvNote}</td>
                  <td>{bewertung.lvPunkte.toFixed(2)}</td>
                  <td>{bewertung.lvProzent.toFixed(2)}</td>
                  <td>{bewertung.lvNote}</td>
                  <td>{bewertung.gesamtProzent.toFixed(2)}</td>
                  <td>{bewertung.gesamtNote}</td>
                </tr>
              {/each}
            </tbody>
          </table>
        </div>
        <h2 class="title">Notenspiegel</h2>
        <table class="table is-hoverable is-fullwidth">
          <thead>
            <tr>
              {#each [1,2,3,4,5,6] as note}
                <th>{note}</th>
              {/each}
            </tr>
          </thead>
          <tbody>
            <tr>
              {#each [1,2,3,4,5,6] as note}
                {#if notenspiegel[note]}
                  <td>{notenspiegel[note]}</td>
                {:else}
                  <td>0</td>
                {/if}
              {/each}
            </tr>
          </tbody>
        </table>

          <div class="field is-grouped">
            <div class="control">
              <input type="text" class="input" bind:value={exportPath} readonly>
            </div>
            <div class="control">
              <button class="button is-info" on:click={selectExportPath}>Pfad auswählen</button>
            </div>
            <div class="control">
              <button class="button is-info" on:click={handleExport}>Exportieren</button>
            </div>
          </div>
      </div> <!-- content -->
    </div> <!-- card-content -->
    
  </div> <!-- card -->
</div> <!-- container -->
