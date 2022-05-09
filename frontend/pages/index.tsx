import type { NextPage } from 'next'
import Head from 'next/head'
import Link from 'next/link'
import Image from 'next/image'
import { useEffect, useState } from 'react'
import useSWR, { mutate } from 'swr'
import styles from '../styles/Home.module.css'
import { getCsrfToken, getUser } from '../utils'

const fetcher = (url: string) =>
  fetch(url).then((r) => {
    return r.json()
  })

const Home: NextPage = () => {
  const { data: projects, error } = useSWR<
    | {
      id: string
      name: string
      default_branch: string
      create_at: string
      update_at: string
    }[]
    | {
      documentation_url: string
      message: string
      data: any
      status: number
    }
  >(`/api/projects`, fetcher)
  const [user, setUser] = useState<{ admin: boolean; image: string, git: string }>()

  useEffect(() => {
    getUser().then((user) => setUser(user))
  }, [user, setUser])

  return (
    <>
      <nav className="w-full py-3 max-w-4xl px-4 items-center mx-auto flex justify-between">
        <div>
          {` `}
          <span className="text-2xl font-bold">
            <Link href="/">Teardrop</Link>
          </span>
        </div>

        <div>
          {user && <img className="w-8 h-8 rounded-full" src={user.image} />}
        </div>
      </nav>
      <main className="max-w-3xl mx-auto mt-8 px-4">
        <div className="flex gap-2 justify-between items-center">
          <h2 className="text-2xl font-bold mb-4">My Projects</h2>
          {user && user!.admin && (
            <button
              className="rounded-lg px-2 py-1 text-sm bg-gray-100 my-2"
              onClick={async () => {
                const csrf = await getCsrfToken()

                fetch(`/api/projects`, {
                  method: `POST`,
                  headers: {
                    'x-csrf-token': csrf ? csrf : ``,
                    'content-type': `application/json`,
                    Accept: `application/json`,
                  },
                  body: JSON.stringify({
                    name: `teardrop-test`,
                    git: `https://github.com/jcdea/website`,
                    defaultBranch: `main`,
                  }),
                })

                await mutate(`/api/projects`)
              }}
            >
              New project
            </button>
          )}
        </div>

        <div className="grid grid-cols-2 gap-2">
          {Array.isArray(projects) &&
            projects.map((proj) => (
              <div key={proj.id} className="rounded-xl border p-2 px-3 max-w-lg ">

                <Link href={`/api/projects/${proj.id}`}>
                  <a className="font-semibold">
                    Name:{` `}
                    {proj.name}
                  </a>
                </Link>

                <span className="text-sm text-gray-400">
                  Id:{` `}
                  {proj.id}
                </span>

              </div>
            ))}
        </div>
      </main>
    </>
  )
}

export default Home
