import Main from '../pages/Main.svelte'
import Login from '../pages/Login.svelte'
import Register from '../pages/Register.svelte'
import Profile from '../pages/Profile.svelte'
import NotFound from '../pages/NotFound.svelte'
import History from '../pages/History.svelte'
import type { RouteInstance } from 'svelte-navigator'

interface IRoute extends RouteInstance {
  Components: typeof Main
}

export default [
  { path: '/', Components: Main },
  { path: '/login', Components: Login },
  { path: '/register', Components: Register },
  { path: '/user/profile', Components: Profile },
  { path: '/user/history', Components: History },
  { path: '*', Components: NotFound },
] as IRoute[]
