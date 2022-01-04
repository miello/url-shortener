<script lang="ts">
import { createEventDispatcher } from "svelte";

  import Chip from "../components/common/Chip.svelte"
  import ArrowLeft from "../components/icons/ArrowLeft.svelte"
  import ArrowRight from "../components/icons/ArrowRight.svelte"

  export let current = 4
  export let allPages = 10

  let dispatchEvent = createEventDispatcher()

  const handleChange = (newCurrent: number) => {
    if(newCurrent < 1 || newCurrent > allPages) return
    current = newCurrent
    dispatchEvent('change', newCurrent)
  }

  $: (() => {
    console.log(current)
  })()

  const handleNext = () => {
    if(current >= allPages) return
    handleChange(current + 1)
  }

  const handlePrev = () => {
    if(current <= 1) return
    handleChange(current - 1)
  }
</script>

<div class="flex items-center justify-center mt-5">
  <div class="flex items-center mr-2">
    <div class={`${current > 1 && "cursor-pointer"} ${current <= 1 && "opacity-0"}`} on:click={() => handlePrev()}>
      <ArrowLeft />
    </div>
    <Chip className="bg-[#ABA9FC] ml-2 mr-2">1</Chip>
    {#if current >= 4 && current <= allPages - 3}
      <Chip className="bg-[#C4C4C4] mr-2 w-3 h-3" />
      <Chip className="bg-[#C4C4C4] mr-2 w-3 h-3" />
      <Chip className="bg-[#C4C4C4] mr-2 w-3 h-3" />
    {/if}


    <div class={`${current < allPages && "cursor-pointer"} ${current >= allPages && "opacity-0"}`} on:click={() => handleNext()}>
      <ArrowRight />
    </div>
  </div>
  <select bind:value={current} class="font-display px-2" on:change={(ev) => handleChange(+ev.currentTarget.value)}>
    {#each Array(allPages) as _, idx}
      <option value={idx + 1}>{idx + 1}</option>
    {/each}
  </select>
</div>