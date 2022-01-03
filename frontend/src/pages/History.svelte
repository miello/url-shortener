<script lang="ts">
  import type { AxiosError, AxiosResponse } from 'axios'

  import { onMount } from 'svelte'

  import Button from '../components/common/Button.svelte'
  import Container from '../components/common/Container.svelte'
  import Loading from '../components/common/Loading.svelte'
  import UserGuard from '../components/hoc/UserGuard.svelte'
  import ConfirmModal from '../modules/ConfirmModal.svelte'
  import { UpdateAlert } from '../stores/AlertStores'
  import type { IHistoryRequest } from '../types/history'
  import { apiClient } from '../utils/apiClient'

  const TABLE_HEADER = [
    'Date Created',
    'Shorten ID',
    'Original',
    'Expires',
    'Action',
  ]
  let promise: Promise<AxiosResponse<IHistoryRequest>> = null
  let currentId: string = ''
  let openModal: boolean = false

  onMount(() => {
    promise = apiClient.get('/short/history')
  })

  const handleDelete = async () => {
    try {
      await apiClient.delete(`/short/${currentId}`)
      promise = apiClient.get('/short/history')
      UpdateAlert({
        status: 'success',
        message: 'Delete Successfully',
      })
    } catch (err) {
      const axiosErr = err as AxiosError<{ error: string }>
      UpdateAlert({
        status: 'error',
        message: axiosErr.response.data?.error && axiosErr.message,
      })
    }
    openModal = false
  }

  const handleClose = () => {
    currentId = ''
    openModal = false
  }
</script>

<UserGuard>
  <div class="w-full h-full flex flex-col justify-center items-center">
    <Container className="flex flex-col items-center min-w-[800px] mx-5">
      <h4 class="font-bold font-display text-2xl mb-4">History</h4>
      <table class="m-0 p-0 border-4 border-solid border-black ">
        {#if promise !== null}
          <tr class="border-2 border-solid border-black font-semibold">
            {#each TABLE_HEADER as topic}
              <th class="border-2 border-solid border-black p-2">{topic}</th>
            {/each}
          </tr>
          {#await promise}
            <Loading />
          {:then value}
            {#each value.data.data as { original, shorten_id, created_at, expires }}
              <tr>
                <td class="border-2 border-solid border-black px-2"
                  >{new Date(created_at).toLocaleString()}</td
                >
                <td class="border-2 border-solid border-black px-2"
                  >{shorten_id}</td
                >
                <td
                  id="original-url"
                  class="border-2 border-solid border-black px-2 overflow-auto max-w-[150px]"
                  >{original}</td
                >
                <td class="border-2 border-solid border-black px-2"
                  >{!!expires ? new Date(expires).toLocaleString() : '-'}</td
                >
                <td class="border-2 border-solid border-black px-2 py-2">
                  <Button className="bg-[#BDFF00] py-1 font-semibold"
                    >Edit</Button
                  >
                  <Button
                    className="bg-red-500 py-1 font-semibold"
                    on:click={() => {
                      currentId = shorten_id
                      openModal = true
                    }}>Delete</Button
                  >
                </td>
              </tr>
            {/each}
          {/await}
        {/if}
      </table>
    </Container>
  </div>
  {#if openModal}
    <ConfirmModal
      on:close={handleClose}
      on:submit={() => handleDelete()}
      message="Are you sure that you want to delete?"
    />
  {/if}
</UserGuard>

<style>
  #original-url::-webkit-scrollbar {
    height: 2px;
  }

  #original-url::-webkit-scrollbar-thumb {
    background: #808080;
  }

  #original-url::-webkit-scrollbar {
    height: 2px;
  }

  #original-url::-webkit-scrollbar {
    height: 2px;
  }
</style>
