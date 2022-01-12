<script lang="ts">
  import Alert from '../components/common/Alert.svelte'
  import { AlertStores, CancelAlert } from '../stores/AlertStores'

  let AlertKeys: string[] = []

  AlertStores.subscribe((val) => {
    AlertKeys = Object.keys(val).splice(0, 10)
  })
</script>

{#each AlertKeys as key, idx (key)}
  <Alert
    handleClose={() => CancelAlert(+key)}
    status={$AlertStores[key].status}
    message={$AlertStores[key].message}
    style={`top: ${3 * idx}rem`}
  />
{/each}
