<script lang="ts">
  import { createEventDispatcher, onDestroy } from 'svelte'
  import { fly } from 'svelte/transition'
  import Button from '../components/common/Button.svelte'
  import Modal from '../components/common/Modal.svelte'
  import Copy from '../components/icons/Copy.svelte'

  const dispatch = createEventDispatcher()
  export let resultUrl = ''
  let isCopied = false
  let prev = 0

  const onClose = () => {
    dispatch('close')
  }

  const onCopy = () => {
    navigator.clipboard.writeText(resultUrl).then(() => {
      isCopied = true
      if (prev === -1) {
        window.clearTimeout(prev)
      }
      prev = window.setTimeout(() => {
        isCopied = false
        prev = -1
      }, 5000)
    })
  }

  onDestroy(() => {
    window.clearTimeout(prev)
  })
</script>

<Modal on:close={onClose}>
  <h3 class="font-display text-2xl font-bold">Here you go</h3>
  <section class="flex items-center w-full max-w-[400px]">
    <div
      class="min-h-[35px] relative ml-2 flex-1 pl-2 py-1 pr-10 border-black border-solid border-2 rounded-lg font-display"
    >
      <p class="font-display font-light sm:text-lg text-sm">{resultUrl}</p>
      <div
        class="absolute cursor-pointer right-1 top-1/2 -translate-y-1/2"
        on:click={onCopy}
      >
        <Copy />
        {#if isCopied}
          <div
            transition:fly
            class="absolute -top-9 bg-black right-1/2 translate-x-1/2 text-white px-3 py-1 text-sm rounded-xl"
          >
            Copied
          </div>
        {/if}
      </div>
    </div>
  </section>
  <Button
    on:click={onClose}
    className="bg-red-500 font-display text-white rounded-lg hover:bg-red-700 transition-colors"
    >Close</Button
  >
</Modal>
