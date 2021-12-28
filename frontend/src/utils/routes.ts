import Main from "../pages/Main.svelte";
import Login from "../pages/Login.svelte";
import type { RouteInstance } from "svelte-navigator";

interface IRoute extends RouteInstance {
  Components: typeof Main;
}

export default [
  { path: "/", Components: Main },
  { path: "/login", Components: Login },
] as IRoute[];
