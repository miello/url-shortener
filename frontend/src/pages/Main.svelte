<script lang="ts">
  import Container from "../components/common/Container.svelte";
  import { apiClient } from '../utils/apiClient';
  import Modal from '../components/common/ResultModal.svelte'
  import Alert from "../components/common/Alert.svelte";
  
  type EventInput = Event & {
    currentTarget: EventTarget & HTMLFormElement;
  };

  let url: string = "";
  let openModal: boolean = false
  let result: string = ""
  
  const handleSubmit = async (e: EventInput) => {
    e.preventDefault()
    try {
      const res = await apiClient.post<{url: string}>("/api/short", { url })
      result = res.data.url
      openModal = true
    } catch(err) {
      console.log(err)
    }
  };
</script>

<div class="w-full h-full flex flex-col justify-center items-center mx-5">
  <h1 class="mb-10 lg:text-6xl md:text-5xl sm:text-4xl text-3xl font-medium">URL Shortener</h1>
  <form on:submit={handleSubmit} class="w-full flex justify-center">
    <Container className="mb-4 max-w-[500px] w-full">
      <div class="flex mb-4 items-center">
        <span class="mr-2 font-display lg:text-xl font-semibold sm:text-md">URL : </span>
        <input bind:value={url} label="Your url" class="px-2.5 py-1 flex-1 rounded-xl" required />
      </div>
      <div class="flex justify-center">
        <button
          type="submit"
          class="bg-yellow-400 color-black px-3 py-2 rounded-xl drop-shadow-md font-display font-bold">
          Shorten
        </button>
      </div>
    </Container>
  </form>
  {#if openModal}
    <Modal on:close={() => {openModal = false}} resultUrl={result} />
  {/if}
  <Alert />
  <!-- <Container _class="max-w-[500px] w-full">
    <h5 class="text-xl text-center mb-2">History</h5>
    <div class="grid grid-cols-2">
      <div>Hello</div>
      <div>Hello</div>
    </div>
  </Container> -->
</div>
