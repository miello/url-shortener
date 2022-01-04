<script lang="ts">
  import Container from '../components/common/Container.svelte'
  import { apiClient } from '../utils/apiClient'
  import ResultModal from '../modules/ResultModal.svelte'
  import Input from '../components/common/Input.svelte'
  import Button from '../components/common/Button.svelte'
  import type { EventInput } from '../types/Event'
  import { UpdateAlert } from '../stores/AlertStores'
  import type { AxiosError } from 'axios'

  let url = ''
  let openModal = false
  let result = ''
  let isLoad = false

  const handleSubmit = async (e: EventInput) => {
    e.preventDefault()
    isLoad = true
    try {
      const res = await apiClient.post<{ url: string }>('/short', { url })
      result = res.data.url
      openModal = true
    } catch (err) {
      const axiosErr = err as AxiosError
      UpdateAlert({
        status: 'error',
        message: axiosErr.response.data?.error && axiosErr.message,
      })
    }
    isLoad = false
  }
</script>

<div class="w-full h-full flex flex-col justify-center items-center mx-5">
  <h1 class="mb-10 lg:text-6xl md:text-5xl sm:text-4xl text-3xl font-medium">
    URL Shortener
  </h1>
  <form on:submit={handleSubmit} class="w-full flex justify-center">
    <Container className="mb-4 max-w-[500px] w-full">
      <div class="flex mb-4 items-center">
        <span class="mr-2 font-display lg:text-xl font-semibold sm:text-md"
          >URL:
        </span>
        <Input bind:value={url} required={true} label="Your URL" />
      </div>
      <div class="flex justify-center">
        <Button
          disabled={isLoad}
          type="submit"
          className="bg-yellow-400 color-black px-3 py-2 rfont-display font-bold"
        >
          Shorten
        </Button>
      </div>
    </Container>
  </form>
  {#if openModal}
    <ResultModal
      on:close={() => {
        openModal = false
      }}
      resultUrl={result}
    />
  {/if}
</div>
