export async function getCsrfToken() {
  const res = await fetch(`/api/csrf`)
  return res.headers.get(`x-csrf-token`)
}

export async function getUser() {
  const res = await fetch(`/api/user`)
  return res.json()
}
