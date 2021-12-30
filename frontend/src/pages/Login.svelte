<script lang="ts">
  import Container from '../components/common/Container.svelte'
  import Input from '../components/common/Input.svelte'
  import type { EventInput } from '../types/Event'
  import { apiClient } from '../utils/apiClient'
  import { useNavigate } from 'svelte-navigator'
  import { UpdateAlert } from '../stores/AlertStores'
  import type { AxiosError } from 'axios'
  import Button from '../components/common/Button.svelte'
  import { GetUser } from '../stores/UserStores'

  export let username: string = ''
  export let password: string = ''
  let loading = false

  let navigate = useNavigate()
  const onSubmit = async (e: EventInput) => {
    e.preventDefault()
    loading = true
    try {
      await apiClient.post('/auth/login', { username, password })

      navigate('/')
      UpdateAlert({
        status: 'success',
        message: 'Login Successfully',
      })

      await GetUser()
    } catch (err) {
      const axiosErr = err as AxiosError
      UpdateAlert({
        status: 'error',
        message: axiosErr.response?.data?.error || axiosErr.message,
      })
    }
    loading = false
  }
</script>

<div class="flex justify-center items-center w-full">
  <Container className="h-fit w-full max-w-[500px] mx-5">
    <form on:submit={onSubmit} class="flex flex-col ">
      <h3 class="font-display text-3xl text-center mb-5 font-bold">Login</h3>
      <div id="container" class="mb-5">
        <span class="font-display">Username</span>
        <Input label="username" bind:value={username} required={true} />
        <span class="font-display">Password</span>
        <Input
          type="password"
          label="password"
          bind:value={password}
          required={true}
        />
      </div>
      <Button
        disabled={loading}
        type="submit"
        className="bg-yellow-500 text-md font-display mb-5">Login</Button
      >
      <hr class="h-[2px] bg-black mb-5" />
      <h6 class="font-display text-md text-center mb-3 font-normal ">
        New Here ?
      </h6>
      <Button
        className="bg-lime-500 text-md font-display text-white"
        on:click={() => navigate('/register')}>Create Account</Button
      >
    </form>
  </Container>
</div>

<style>
  #container {
    display: grid;
    grid-template-columns: min-content auto;
    row-gap: 0.75rem;
    column-gap: 0.75rem;
    align-items: center;
  }
</style>
