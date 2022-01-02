<script lang="ts">
  export let className: string = ''
  export let src: string = ''
  export let name: string = ''
</script>

<div
  class={`rounded-full bg-[#C4C4C4] ${className} w-10 h-10 cursor-pointer select-none`}
  on:click
>
  {#if !src}
    <div class="flex justify-center items-center h-full">
      <h6 class="font-display font-semibold text-xl">
        {name.length !== 0 ? name[0].toUpperCase() : ''}
      </h6>
    </div>
  {:else}
    <img alt="profile" {src} />
  {/if}
</div>
