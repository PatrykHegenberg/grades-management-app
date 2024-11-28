<script>
  export let notenspiegel;
</script>

<h3 class="text-2xl font-bold mt-8 mb-4">Notenspiegel</h3>
<div class="overflow-x-auto">
  <table class="table w-full">
    <thead>
      <tr>
        {#each [1, 2, 3, 4, 5, 6] as note}
          <th>{note}</th>
        {/each}
      </tr>
    </thead>
    <tbody>
      <tr>
        {#each [1, 2, 3, 4, 5, 6] as note}
          <td>{notenspiegel[note] || 0}</td>
        {/each}
      </tr>
    </tbody>
  </table>
</div>
