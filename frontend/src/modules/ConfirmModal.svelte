<script lang="ts">
  import { createEventDispatcher } from 'svelte'
  import Button from '../components/common/Button.svelte'
  import Modal from '../components/common/Modal.svelte'

  export let message = ''
  let eventDispatch = createEventDispatcher()

  const onClose = () => {
    eventDispatch('close')
  }

  const onSubmit = () => {
    eventDispatch('submit')
  }
</script>

<Modal on:close={onClose}>
  <h3 class="font-display text-2xl font-bold">{message}</h3>
  <div class="w-full flex justify-evenly px-4">
    <Button className="bg-green-500 text-white" on:click={onSubmit}
      >Confirm</Button
    >
    <Button className="bg-red-500 text-white" on:click={onClose}>Cancel</Button>
  </div>
</Modal>
