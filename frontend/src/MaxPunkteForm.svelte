<script>
  export let hvMax;
  export let lvMax;
  export let hvGewichtung;
  export let lvGewichtung;
  export let onSubmit;
</script>

<form on:change|preventDefault={onSubmit} class="mb-4">
  <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
    <div class="form-control">
      <label class="label" for="hvMax">HV-Max-Punkte</label>
      <input
        id="hvMax"
        class="input input-bordered"
        type="number"
        step="0.5"
        bind:value={hvMax}
        required
      />
    </div>
    <div class="form-control">
      <label class="label" for="hvGewichtung">HV-Gewichtung in %</label>
      <input
        id="hvGewichtung"
        class="input input-bordered"
        type="number"
        bind:value={hvGewichtung}
        required
      />
    </div>
    <div class="form-control">
      <label class="label" for="lvMax">LV-Max-Punkte</label>
      <input
        id="lvMax"
        class="input input-bordered"
        type="number"
        step="0.5"
        bind:value={lvMax}
        required
      />
    </div>
    <div class="form-control">
      <label class="label" for="lvGewichtung">LV-Gewichtung in %</label>
      <input
        id="lvGewichtung"
        class="input input-bordered"
        type="number"
        bind:value={lvGewichtung}
        required
      />
    </div>
  </div>
</form>
