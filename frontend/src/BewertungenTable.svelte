<script>
  import BewertungRow from "./BewertungRow.svelte";
  export let bewertungen;
  export let onToggleWertung;
</script>

<div class="max-h-[500px] overflow-y-auto overflow-x-auto">
  <table class="table w-full">
    <!-- <div class="overflow-x-auto overflow-y auto"> -->
    <!--   <table class="table w-full table-container"> -->
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
      {#each bewertungen as bewertung (bewertung.id)}
        <BewertungRow {bewertung} {onToggleWertung} />
      {/each}
    </tbody>
  </table>
</div>

<!-- <style> -->
<!--   .table-container { -->
<!--     max-height: 400px; -->
<!--     overflow-y: auto; -->
<!--   } -->
<!-- </style> -->
