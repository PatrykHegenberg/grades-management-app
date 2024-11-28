<script>
  export let bewertung;
  export let onToggleWertung;
</script>

<tr>
  <td>
    <input
      type="checkbox"
      class="toggle"
      checked={bewertung.gewertet}
      on:change={() => onToggleWertung(bewertung.id)}
    />
  </td>
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
