import type { NextPage } from 'next'
import Head from 'next/head'
import Link from 'next/link'
import Image from 'next/image'
import { useEffect, useState } from 'react'
import useSWR, { mutate } from 'swr'
import styles from '../styles/Home.module.css'
import { fetcher, getCsrfToken } from '../utils'
import Nav from '../components/Nav'
import { useSession } from '../session'

const Home: NextPage = () => {
  const { data: projects, error } = useSWR<
    | {
      id: string
      name: string
      git: string
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

  const { data: session } = useSession()
  return (
    <>
      <Nav />
      <main className="max-w-4xl mx-auto mt-8 px-4">
        <div className="flex gap-2 justify-between items-center">
          <h2 className="text-2xl font-bold mb-4">My Projects</h2>
          {session && session.user!.admin && (
            <button
              className="rounded-lg px-4 py-1 text-sm border my-2"
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
              New
            </button>
          )}
        </div>

        <div className="grid lg:grid-cols-2 gap-2">
          {Array.isArray(projects) &&
            projects.map((proj) => (
              <div
                key={proj.id}
                className="rounded-xl border p-2 px-3 max-w-lg "
              >
                <Link href={`/projects/${proj.id}`}>
                  <a className="font-semibold text-lg">{proj.name}</a>
                </Link>
                <div className="">
                  <div className="text-sm text-gray-600">{proj.git}</div>

                  <div className="text-xs text-gray-400">
                    ID:{` `}
                    {proj.id}
                  </div>
                </div>
              </div>
            ))}
        </div>
      </main>
    </>
  )
}

export default Home
