<script lang="ts">
  import type { AxiosError } from 'axios'
  import { createEventDispatcher } from 'svelte'
  import Button from '../components/common/Button.svelte'
  import Input from '../components/common/Input.svelte'
  import Modal from '../components/common/Modal.svelte'
  import { EXPIRES_TIME } from '../constant/time'
  import { UpdateAlert } from '../stores/AlertStores'
  import type { ExpiresTimeType } from '../types/time'
  import { apiClient } from '../utils/apiClient'

  const eventDispatch = createEventDispatcher()

  export let shortId = ''
  export let original = ''
  export let expires: ExpiresTimeType = 'None'

  const onClose = () => {
    eventDispatch('close')
  }

  const onSubmit = async () => {
    try {
      await apiClient.put(`/short/${shortId}`, {
        url: original,
        expires,
      })
      UpdateAlert({
        status: 'success',
        message: 'Edit Successfully',
      })
      eventDispatch('submit')
    } catch (err) {
      const axiosErr = err as AxiosError<{ error: string }>
      UpdateAlert({
        status: 'error',
        message: axiosErr,
      })
    }
  }
</script>

<Modal on:close={onClose}>
  <h3 class="font-display text-2xl font-bold">Edit</h3>
  <section id="containerEdit">
    <span class="mr-2 font-display lg:text-xl font-semibold sm:text-lg text-sm"
      >Short ID:
    </span>
    <span class="font-display font-semibold lg:text-xl sm:text-lg text-sm"
      >{shortId}</span
    >
    <span class="mr-2 font-display lg:text-xl font-semibold sm:text-lg text-sm"
      >Original:
    </span>
    <Input
      bind:value={original}
      required={true}
      label="Your URL"
      className="font-display lg:text-xl sm:text-lg text-sm"
    />
    <span class="mr-2 font-display font-semibold lg:text-xl sm:text-lg text-sm"
      >Expires:
    </span>
    <select
      class="pl-1 pr-3 max-w-[200px] rounded-lg box-content px-2.5 py-1 hover:border-black border-2 lg:text-xl sm:text-lg text-sm"
      bind:value={expires}
    >
      {#each EXPIRES_TIME as expire}
        <option value={expire}>{expire}</option>
      {/each}
    </select>
  </section>
  <div class="flex justify-evenly w-full">
    <Button className="bg-green-500 text-white" on:click={onSubmit}>Edit</Button
    >
    <Button className="bg-red-500 text-white" on:click={onClose}>Cancel</Button>
  </div>
</Modal>

<style>
  #containerEdit {
    display: grid;
    grid-template-columns: auto auto;
    row-gap: 0.75rem;
    column-gap: 0.5rem;
    align-items: center;
  }
</style>
