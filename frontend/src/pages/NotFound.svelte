<script>
  import { Link } from 'svelte-navigator'
  import Button from '../components/common/Button.svelte'
  import Container from '../components/common/Container.svelte'
</script>

<div class="flex justify-center items-center w-full h-full">
  <Container
    className="flex flex-col justify-center items-center max-w-[400px] w-full h-fit p-5"
  >
    <h1 class="font-display mb-6 font-bold text-2xl">404 Not Found</h1>
    <div class="mb-6 max-w-[300px] font-display text-md">
      Oops. You might accidentally access to expired or non-exists url
    </div>
    <Link to="/">
      <Button
        className="bg-red-600 border-none m-auto text-white rounded-2xl text"
      >
        Back to main page
      </Button>
    </Link>
  </Container>
</div>
