import useSWR from 'swr'

export async function getCsrfToken() {
  const res = await fetch(`/api/csrf`)
  return res.headers.get(`x-csrf-token`)
}

export const fetcher = (url: string) =>
  fetch(url).then((r) => {
    return r.json()
  })
