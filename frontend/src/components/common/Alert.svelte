<script lang="ts">
  import CloseIcon from '../icons/CloseIcon.svelte'
  import { AlertStores, OpenError, CancelAlert } from '../../stores/AlertStores'
  import CorrectIcon from '../icons/CorrectIcon.svelte'
  import ErrorIcon from '../icons/ErrorIcon.svelte'
</script>

<div
  class={`absolute left-1/2 ${
    $OpenError ? 'top-10' : '-top-32'
  } -translate-x-1/2 transition-all duration-300 ease-linear z-50 w-fit`}
>
  <div
    class={`red-500 px-4 py-2 mx-4 flex items-center text-white rounded-xl ${
      $AlertStores.status === 'error' ? 'bg-red-600' : 'bg-green-600'
    }`}
  >
    {#if $AlertStores.status === 'error'}
      <ErrorIcon />
    {:else if $AlertStores.status === 'success'}
      <CorrectIcon />
    {/if}
    <p class="font-display pr-1 pl-2 text-sm lg:text-base">
      {$AlertStores.message}
    </p>
    <div on:click={CancelAlert} class="cursor-pointer">
      <CloseIcon color="white" />
    </div>
  </div>
</div>
