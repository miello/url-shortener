<script lang="ts">
  import axios from "axios";
  import Container from "../../components/common/Container.svelte";

  type EventInput = Event & {
    currentTarget: EventTarget & HTMLFormElement;
  };

  let url: string = "";
  const handleSubmit = async (e: EventInput) => {
    e.preventDefault()
    try {
      const res = await axios.post('http://localhost:8080/api/short', { url: url.toString() })
      console.log(res.data)
    } catch(err) {
      console.log(err)
    }
  };
</script>

<div class="h-full flex flex-col justify-center items-center mx-5">
  <h1 class="mb-10 lg:text-6xl md:text-5xl sm:text-4xl text-3xl font-medium">URL Shortener</h1>
  <form on:submit={handleSubmit} class="w-full flex justify-center">
    <Container _class="mb-4 max-w-[500px] w-full">
      <div class="flex mb-4 items-center">
        <span class="mr-2 font-display text-lg font-semibold">URL : </span>
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
  <!-- <Container _class="max-w-[500px] w-full">
    <h5 class="text-xl text-center mb-2">History</h5>
    <div class="grid grid-cols-2">
      <div>Hello</div>
      <div>Hello</div>
    </div>
  </Container> -->
</div>
