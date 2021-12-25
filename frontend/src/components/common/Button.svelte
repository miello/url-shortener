<script lang="ts">
  export let className: string = ""

</script>

<button class={`drop-shadow-lg px-4 py-2 ${className}`} on:click><slot /></button>
