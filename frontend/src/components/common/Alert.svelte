<script lang="ts">
  import CloseIcon from '../icons/CloseIcon.svelte';
  import { AlertStores, OpenError, CancelAlert } from '../../stores/AlertStores'

</script>

<div class={`absolute left-1/2 ${$OpenError ? "top-10" : "-top-20"} -translate-x-1/2 transition-all duration-300 ease-linear`}>
  <div class={`red-500 px-4 py-2 flex text-white rounded-xl ${$AlertStores.status === "error" ? "bg-red-600": "bg-green-600"}`} >
    <p class="font-display pr-1">{$AlertStores.message}</p>
    <div on:click={CancelAlert} class="cursor-pointer">
      <CloseIcon color="white"  />
    </div>
  </div>
</div>