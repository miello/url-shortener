<script lang="ts">
  export let className = ''
  export let rounded = true
</script>

<section
  class={`p-4 backdrop-blur-md drop-shadow-md bg-white box-border ${
    rounded && 'rounded-xl'
  } bg-opacity-50 ${className}`}
  on:click
>
  <slot />
</section>
