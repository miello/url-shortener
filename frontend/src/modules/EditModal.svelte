<script lang="ts">
  import type { AxiosError } from 'axios'
  import { createEventDispatcher } from 'svelte'
  import Button from '../components/common/Button.svelte'
  import Input from '../components/common/Input.svelte'
  import Modal from '../components/common/Modal.svelte'
  import { UpdateAlert } from '../stores/AlertStores'
  import { apiClient } from '../utils/apiClient'

  const eventDispatch = createEventDispatcher()

  export let shortId = ''
  export let original = ''

  const onClose = () => {
    eventDispatch('close')
  }

  const onSubmit = async () => {
    try {
      await apiClient.put(`/short/${shortId}`, {
        url: original,
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
        message: axiosErr.response.data?.error && axiosErr.message,
      })
    }
  }
</script>

<Modal on:close={onClose}>
  <h3 class="font-display text-2xl font-bold">Edit</h3>
  <section id="containerEdit">
    <span class="mr-2 font-display lg:text-xl font-semibold sm:text-md"
      >Shorten ID:
    </span>
    <span class="font-display lg:text-xl font-semibold sm:text-md"
      >{shortId}</span
    >
    <span class="mr-2 font-display lg:text-xl font-semibold sm:text-md"
      >Original:
    </span>
    <Input
      bind:value={original}
      required={true}
      label="Your URL"
      className="font-display"
    />
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
    grid-template-columns: max-content auto;
    row-gap: 0.75rem;
    column-gap: 0.75rem;
    align-items: center;
  }
</style>
