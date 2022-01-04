<script lang="ts">
  export let className = ""
</script>

<div class={`w-8 h-8 rounded-full flex justify-center items-center font-display ${className}`} on:click>
  <slot />
</div>