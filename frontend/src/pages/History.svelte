<script lang="ts">
  import type { AxiosError, AxiosResponse } from 'axios'

  import { onMount } from 'svelte'

  import Button from '../components/common/Button.svelte'
  import Container from '../components/common/Container.svelte'
  import Loading from '../components/common/Loading.svelte'
  import UserGuard from '../components/hoc/UserGuard.svelte'
  import ConfirmModal from '../modules/ConfirmModal.svelte'
import Pagination from '../modules/Pagination.svelte'
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
  let currentId = ''
  let openModal = false
  let current = 1
  let allPages = 1

  const updateHistory = (newVal: number = current) => {
    promise = apiClient.get<IHistoryRequest>(`/short/history?pages=${newVal}`).then((val) => {
      current = val.data.current
      allPages = val.data.all_pages
      return val
    }).catch((err) => {
      const axiosErr = err as AxiosError<{ error: string }>
      UpdateAlert({
        status: 'error',
        message: axiosErr.response.data?.error && axiosErr.message,
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
        message: axiosErr.response.data?.error && axiosErr.message,
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
  <div class="w-full h-full flex items-center overflow-auto">
    <div class="w-full px-3">
      <div class="flex justify-center">
        <h4 class="font-bold font-display text-2xl mb-4">History</h4>
      </div>
      <Container className="max-w-[800px] overflow-auto mx-auto">
        <table class="m-0 p-0 border-4 border-solid border-black w-full">
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
                    >{new Date(created_at).toLocaleString("en-GB")}</td
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
                    >{expires ? new Date(expires).toLocaleString("en-GB") : '-'}</td
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
      <Pagination current={current} allPages={allPages} on:change={handleNavigate} />
    </div>
  </div>
  {#if openModal}
    <ConfirmModal
      on:close={handleClose}
      on:submit={handleDelete}
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
