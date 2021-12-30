<script lang="ts">
  export let className: string = ''
  export let rounded: boolean = true
</script>

<section
  class={`p-4 backdrop-blur-md drop-shadow-md bg-white ${
    rounded && 'rounded-xl'
  } bg-opacity-50 ${className}`}
  on:click
>
  <slot />
</section>
