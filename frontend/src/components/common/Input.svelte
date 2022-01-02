<script lang="ts">
  export let value = ''
  export let label = ''
  export let className = ''
  export let placeHolder = ''
  export let required = false
  export let type = 'text'

  const setType = (node: HTMLInputElement) => {
    node.type = type
  }
</script>

<input
  bind:value
  use:setType
  {label}
  placeholder={placeHolder}
  class={`box-content px-2.5 py-1 flex-1 rounded-xl hover:border-black border-2 ${className}`}
  {required}
/>
