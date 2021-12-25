<script lang="ts">
  import { fly } from 'svelte/transition'
  export let reason: string = ""
  export let type: "success" | "error" = "error"
  
  let isOpen: boolean = false

</script>

{#if isOpen}
  <div transition:fly class={`absolute left-1/2 top-10 red-500 -translate-x-1/2 px-4 py-2 text-white rounded-xl ${type === "error" ? "bg-red-600": "bg-green-600"}`} >
    {reason}
  </div>
{/if}