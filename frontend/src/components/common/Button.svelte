<script lang="ts">
  export let className = ''
  export let type: 'submit' | 'reset' | 'button' = 'button'
  export let disabled = false
</script>

<button
  {disabled}
  class={`font-display drop-shadow-lg px-4 py-2 rounded-xl ${className} disabled:bg-gray-500 transition-all`}
  {type}
  on:click><slot /></button
>
