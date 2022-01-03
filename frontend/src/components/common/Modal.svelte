<script lang="ts">
  import { createEventDispatcher } from 'svelte'
  import { fly, scale } from 'svelte/transition'
  import Container from './Container.svelte'
  import CloseIcon from '../icons/CloseIcon.svelte'

  let eventDispatch = createEventDispatcher()

  const onClose = () => {
    eventDispatch('close')
  }
</script>

<div
  transition:fly
  on:click={onClose}
  class="absolute top-0 z-50 backdrop-blur-md bg-black bg-opacity-50 w-full h-full"
>
  <div transition:scale class="flex justify-center items-center w-full h-full">
    <Container
      on:click={(e) => {
        e.stopImmediatePropagation()
      }}
      className="m-8 max-w-[600px] relative w-full h-fit bg-opacity-100 flex flex-col items-center gap-6 py-8"
    >
      <div on:click={onClose} class="absolute right-2 top-2 cursor-pointer">
        <CloseIcon />
      </div>
      <slot />
    </Container>
  </div>
</div>
