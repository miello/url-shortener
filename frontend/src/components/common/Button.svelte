<script lang="ts">
  export let className: string = ""
  export let type: 'submit' | 'reset' | 'button' = "button"
</script>

<button class={`drop-shadow-lg px-4 py-2 ${className}`} type={type} on:click><slot /></button>
