<script>
  export let vorname;
  export let nachname;
  export let hvPunkte;
  export let lvPunkte;
  export let onSubmit;
  export let disabled = false;
</script>

<form on:submit|preventDefault={onSubmit} class="mb-4">
  <div class="flex flex-wrap gap-2">
    <input
      class="input input-bordered"
      type="text"
      bind:value={vorname}
      placeholder="Vorname"
      required
    />
    <input
      class="input input-bordered"
      type="text"
      bind:value={nachname}
      placeholder="Nachname"
      required
    />
    <input
      class="input input-bordered"
      type="number"
      bind:value={hvPunkte}
      placeholder="HV-Punkte"
      required
    />
    <input
      class="input input-bordered"
      type="number"
      bind:value={lvPunkte}
      placeholder="LV-Punkte"
      required
    />
    <button type="submit" class="btn btn-primary" {disabled}>
      Hinzufügen
    </button>
  </div>
</form>

