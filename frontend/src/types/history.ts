interface IShortenHistory {
  original: string
  shorten_id: string
  created_at: string
  expires: string
}

export interface IHistoryRequest {
  all_pages: number
  count: number
  current: number
  data: IShortenHistory[]
}
