<script lang="ts">
import { createEventDispatcher } from "svelte";
import { fly } from "svelte/transition";
import Button from "./Button.svelte";
import Container from "./Container.svelte";
import CloseIcon from '../icons/CloseIcon.svelte'
import Copy from "../icons/Copy.svelte";

const dispatch = createEventDispatcher()
export let resultUrl: string = ""
type ContainerEvent = MouseEvent & {
    currentTarget: EventTarget & HTMLDivElement;
}

const onClose = (e: ContainerEvent) => {
  e.stopImmediatePropagation()

  dispatch('close')
}

const onCopy = () => {
  navigator.clipboard.writeText(resultUrl).then(() => {

  }).catch(() => {
    
  })
}

</script>

<div transition:fly on:click={onClose} class="absolute z-50 flex justify-center items-center backdrop-blur-md bg-black bg-opacity-50 w-full h-full">
  <Container on:click={(e) => { e.stopImmediatePropagation() }} className="m-10 max-w-[600px] relative w-full h-fit bg-opacity-100 flex flex-col items-center gap-6 py-8">
    <div on:click={onClose} class="absolute right-2 top-2 cursor-pointer" onClick={onClose}>
      <CloseIcon />
    </div>
    <h3 class="font-display text-2xl font-bold">Here you go</h3>
    <section class="flex items-center w-full max-w-[400px]">
      <div class="relative ml-2 flex-1 pl-2 py-1 pr-10 border-black border-solid border-2 rounded-lg font-display">
        <p class="font-display font-light sm:text-lg text-sm">{resultUrl}</p>
        <div class="absolute cursor-pointer right-1 top-1/2 -translate-y-1/2" on:click={onCopy}>
          <Copy />
        </div>
      </div>
    </section>
    <Button on:click={onClose} className="bg-red-500 font-display text-white rounded-lg hover:bg-red-700 transition-colors">Close</Button>
  </Container>
</div>