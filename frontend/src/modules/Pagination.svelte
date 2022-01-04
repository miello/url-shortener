<script lang="ts">
  import { createEventDispatcher } from 'svelte'

  import Chip from '../components/common/Chip.svelte'
  import ArrowLeft from '../components/icons/ArrowLeft.svelte'
  import ArrowRight from '../components/icons/ArrowRight.svelte'

  export let current = 4
  export let allPages = 10
  let innerChip: number[] = []

  let dispatchEvent = createEventDispatcher()

  const handleChange = (newCurrent: number) => {
    if (newCurrent < 1 || newCurrent > allPages) return
    current = newCurrent
    dispatchEvent('change', newCurrent)
  }

  $: (() => {
    let newChip = []
    if (current === 1) {
      for (let i = 1; i <= 3; ++i) {
        if (i <= allPages && i !== 1 && i !== allPages) {
          newChip.push(i)
        }
      }
    } else if (current === allPages) {
      for (let i = allPages - 2; i <= allPages; ++i) {
        if (i >= 1 && i !== 1 && i !== allPages) {
          newChip.push(i)
        }
      }
    } else {
      for (let i = current - 1; i <= current + 1; ++i) {
        if (i >= 1 && i <= allPages && i !== 1 && i !== allPages) {
          newChip.push(i)
        }
      }
    }
    innerChip = [...newChip]
  })()

  const handleNext = () => {
    if (current >= allPages) return
    handleChange(current + 1)
  }

  const handlePrev = () => {
    if (current <= 1) return
    handleChange(current - 1)
  }
</script>

<div class="flex items-center justify-center mt-5">
  <div class="flex items-center mr-2">
    <div
      class={`${current > 1 && 'cursor-pointer'} ${
        current <= 1 && 'opacity-0'
      }`}
      on:click={() => handlePrev()}
    >
      <ArrowLeft />
    </div>
    <Chip
      className={`${
        current === 1 ? 'bg-white' : 'bg-[#ABA9FC]'
      } mr-2 cursor-pointer`}
      on:click={() => handleChange(1)}>1</Chip
    >
    {#if (current >= 4 && allPages >= 6) || (current >= allPages - 1 && allPages >= 5)}
      <Chip className="bg-[#C4C4C4] mr-2 w-3 h-3" />
      <Chip className="bg-[#C4C4C4] mr-2 w-3 h-3" />
      <Chip className="bg-[#C4C4C4] mr-2 w-3 h-3" />
    {/if}
    {#each innerChip as val}
      <Chip
        className={`${
          current === val ? 'bg-white' : 'bg-[#ABA9FC]'
        } mr-2 cursor-pointer`}
        on:click={() => handleChange(val)}>{val}</Chip
      >
    {/each}
    {#if (current <= allPages - 3 && allPages >= 6) || (current <= 2 && allPages >= 5)}
      <Chip className="bg-[#C4C4C4] mr-2 w-3 h-3" />
      <Chip className="bg-[#C4C4C4] mr-2 w-3 h-3" />
      <Chip className="bg-[#C4C4C4] mr-2 w-3 h-3" />
    {/if}
    {#if allPages > 1}
      <Chip
        className={`${
          current === allPages ? 'bg-white' : 'bg-[#ABA9FC]'
        } mr-2 cursor-pointer`}
        on:click={() => handleChange(allPages)}>{allPages}</Chip
      >
    {/if}
    <div
      class={`${current < allPages && 'cursor-pointer'} ${
        current >= allPages && 'opacity-0'
      }`}
      on:click={() => handleNext()}
    >
      <ArrowRight />
    </div>
  </div>
  <select
    bind:value={current}
    class="font-display px-2"
    on:change={(ev) => handleChange(+ev.currentTarget.value)}
  >
    {#each Array(allPages) as _, idx}
      <option value={idx + 1}>{idx + 1}</option>
    {/each}
  </select>
</div>
