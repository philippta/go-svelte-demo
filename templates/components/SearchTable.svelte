<script>
  export let data;
  export let filter;

  $: header = data[0];
  $: rows = data
    .slice(1)
    .filter((row) => row.some((cell) => !filter || cell.includes(filter)));
</script>

Search: <input bind:value={filter} />

<table>
  <thead>
    <tr>
      {#each header as cell}
        <td>{cell}</td>
      {/each}
    </tr>
  </thead>
  {#each rows as row}
    <tr>
      {#each row as cell}
        <td>{cell}</td>
      {/each}
    </tr>
  {/each}
</table>
