<script lang="ts">
  import CloseIcon from '../icons/CloseIcon.svelte'
  import CorrectIcon from '../icons/CorrectIcon.svelte'
  import ErrorIcon from '../icons/ErrorIcon.svelte'
  import { scale } from 'svelte/transition'

  export let style: string = ''
  export let status: 'error' | 'success' = 'error'
  export let message = ''
  export let handleClose: () => void = () => {}
</script>

<div
  transition:scale
  {style}
  class={`absolute left-1/2 -translate-x-1/2 transition-all duration-300 ease-linear z-50 w-fit`}
>
  <div
    class={`red-500 px-4 py-2 mx-4 flex items-center text-white rounded-xl ${
      status === 'error' ? 'bg-red-600' : 'bg-green-600'
    }`}
  >
    {#if status === 'error'}
      <ErrorIcon />
    {:else if status === 'success'}
      <CorrectIcon />
    {/if}
    <p class="font-display pr-1 pl-2 text-sm lg:text-base">
      {message}
    </p>
    <div on:click={handleClose} class="cursor-pointer">
      <CloseIcon color="white" />
    </div>
  </div>
</div>
