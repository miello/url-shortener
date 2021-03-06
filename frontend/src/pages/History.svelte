<script lang="ts">
  import type { AxiosError, AxiosResponse } from 'axios'

  import { onMount } from 'svelte'

  import Button from '../components/common/Button.svelte'
  import Container from '../components/common/Container.svelte'
  import Loading from '../components/common/Loading.svelte'
  import UserGuard from '../components/hoc/UserGuard.svelte'
  import ConfirmModal from '../modules/ConfirmModal.svelte'
  import EditModal from '../modules/EditModal.svelte'
  import Pagination from '../modules/Pagination.svelte'
  import { UpdateAlert } from '../stores/AlertStores'
  import type { IHistoryRequest } from '../types/history'
  import type { ExpiresTimeType } from '../types/time'
  import { apiClient } from '../utils/apiClient'

  const TABLE_HEADER = [
    'Date Created',
    'Expires',
    'Shorten ID',
    'Original',
    'Action',
  ]
  let promise: Promise<AxiosResponse<IHistoryRequest>> = null
  let openModal = false

  let current = 1
  let allPages = 1

  let currentId = ''
  let originalUrl = ''
  let rawExpires: ExpiresTimeType = 'None'
  let editModal = false

  const updateHistory = (newVal: number = current) => {
    promise = apiClient
      .get<IHistoryRequest>(`/short/history?pages=${newVal}`)
      .then((val) => {
        current = val.data.current
        allPages = val.data.all_pages
        return val
      })
      .catch((err) => {
        const axiosErr = err as AxiosError<{ error: string }>
        UpdateAlert({
          status: 'error',
          message: axiosErr,
        })
        return err
      })
  }

  onMount(() => {
    updateHistory()
  })

  const handleDelete = async () => {
    try {
      await apiClient.delete(`/short/${currentId}`)
      updateHistory()
      UpdateAlert({
        status: 'success',
        message: 'Delete Successfully',
      })
    } catch (err) {
      const axiosErr = err as AxiosError<{ error: string }>
      UpdateAlert({
        status: 'error',
        message: axiosErr,
      })
    }
    openModal = false
  }

  const handleNavigate = (e: CustomEvent<number>) => {
    updateHistory(e.detail)
  }

  const handleClose = () => {
    currentId = ''
    openModal = false
  }
</script>

<UserGuard>
  <div class="w-full h-full flex items-center overflow-auto py-3">
    <div class="w-full px-3">
      <div class="flex justify-center">
        <h4 class="font-bold font-display text-2xl mb-4">History</h4>
      </div>
      <Container
        className="max-w-[900px] overflow-auto mx-auto md:max-h-full max-h-[300px]"
      >
        <table class="m-0 p-0 border-4 border-solid border-black w-full">
          {#if promise !== null}
            {#await promise}
              <div class="flex justify-center items-center h-full">
                <Loading />
              </div>
            {:then value}
              <tr class="border-2 border-solid border-black font-semibold">
                {#each TABLE_HEADER as topic}
                  <th class="border-2 border-solid border-black p-2 h-fit"
                    >{topic}</th
                  >
                {/each}
              </tr>
              {#each value.data.data as { original, shorten_id, created_at, expires, raw_expires }}
                <tr>
                  <td
                    class="border-2 border-solid border-black px-2 text-center font-medium"
                    >{new Date(created_at).toLocaleString('en-GB')}</td
                  >
                  <td
                    class="border-2 border-solid border-black px-2 text-center font-medium"
                    >{expires && raw_expires !== 'None'
                      ? new Date(expires).toLocaleString('en-GB')
                      : '-'}</td
                  >
                  <td
                    class="border-2 border-solid border-black px-2 text-center"
                    >{shorten_id}</td
                  >
                  <td
                    id="original-url"
                    class="border-2 border-solid border-black px-2 overflow-auto max-w-[150px] text-center"
                    >{original}</td
                  >
                  <td
                    class="border-2 py-2 border-solid border-black px-2 m-auto"
                  >
                    <div class="flex  justify-evenly">
                      <Button
                        className="bg-[#BDFF00] py-1 font-semibold"
                        on:click={() => {
                          currentId = shorten_id
                          originalUrl = original
                          console.log(raw_expires)
                          rawExpires = raw_expires
                          editModal = true
                        }}>Edit</Button
                      >
                      <Button
                        className="bg-red-500 py-1 font-semibold ml-1"
                        on:click={() => {
                          currentId = shorten_id
                          openModal = true
                        }}>Delete</Button
                      >
                    </div>
                  </td>
                </tr>
              {/each}
            {/await}
          {/if}
        </table>
      </Container>
      <Pagination {current} {allPages} on:change={handleNavigate} />
    </div>
  </div>
  {#if openModal}
    <ConfirmModal
      on:close={handleClose}
      on:submit={handleDelete}
      message="Are you sure that you want to delete?"
    />
  {/if}
  {#if editModal}
    <EditModal
      shortId={currentId}
      original={originalUrl}
      expires={rawExpires}
      on:close={() => (editModal = false)}
      on:submit={() => {
        updateHistory()
        editModal = false
      }}
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
