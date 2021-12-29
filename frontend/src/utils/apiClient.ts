import axios from "axios"
import { BACKEND_API } from "./env"

export const apiClient = axios.create({
  baseURL: BACKEND_API,
  timeout: 60000,
  withCredentials: true,
})
