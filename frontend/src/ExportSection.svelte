<script>
  export let exportPath;
  export let onSelectPath;
  export let onExport;
</script>

<div class="flex flex-wrap gap-2 mt-4">
  <input
    type="text"
    class="input input-bordered flex-grow"
    bind:value={exportPath}
    readonly
  />
  <button class="btn btn-info" on:click={onSelectPath}>Pfad auswählen</button>
  <button class="btn btn-info" on:click={onExport}>Exportieren</button>
</div>
