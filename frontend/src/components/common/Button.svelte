<script lang="ts">
  export let className: string = ""
  export let type: 'submit' | 'reset' | 'button' = "button"
  export let disabled = false
</script>

<button {disabled} class={`drop-shadow-lg px-4 py-2 rounded-xl ${className} disabled:bg-gray-500 transition-all`} type={type} on:click><slot /></button>
