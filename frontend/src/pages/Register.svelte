<script lang="ts">
  import { useNavigate } from "svelte-navigator";

  import Container from "../components/common/Container.svelte";
  import Input from "../components/common/Input.svelte";
  import type { EventInput } from "../types/Event";
  import { apiClient } from "../utils/apiClient";

  export let handleName: string = ""
  export let username: string = ""
  export let password: string = ""
  export let confirmPassword: string = ""

  const navigate = useNavigate()

  const onSubmit = async (e: EventInput) => {
    try {
      e.preventDefault()
      const res = await apiClient.post('/register', {
        handle: handleName,
        username,
        password
      })
      if(res.status === 200) {
        navigate('/login')
      } 
    } catch(err) {
      console.log(err)
    }
  }

</script>
<Container>
  <form on:submit={onSubmit}>
    <h3>Register</h3>
    <span>Handle Name</span>
    <Input bind:value={handleName} required={true} />
    <span>Username</span>
    <Input bind:value={username} required={true} />
    <span>Password</span>
    <Input bind:value={password} required={true} />
    <span>Confirm Password</span>
    <Input bind:value={confirmPassword} required={true} />
  </form>
</Container>