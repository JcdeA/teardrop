import React from 'react'
import useSWR from 'swr'
import { fetcher } from '../utils'
import {
  SessionContextValue,
  SessionProviderProps,
  UseSessionOptions,
} from './types'

const SessionContext = React.createContext(undefined)

export function useSession<R extends boolean>(
  options?: UseSessionOptions<R>,
): SessionContextValue<R> {
  // @ts-expect-error Satisfy TS if branch on line below
  const value: SessionContextValue<R> = React.useContext(SessionContext)
  if (!value && process.env.NODE_ENV !== `production`) {
    throw new Error(
      `[auth]: \`useSession\` must be wrapped in a <SessionProvider />`,
    )
  }

  const { required, onUnauthenticated } = options ?? {}

  const requiredAndNotLoading = required && value.status === `unauthenticated`

  React.useEffect(() => {
    if (requiredAndNotLoading) {
      const url = `/api/auth/signin?${new URLSearchParams({
        error: `SessionRequired`,
        callbackUrl: window.location.href,
      })}`
      if (onUnauthenticated) onUnauthenticated()
      else window.location.href = url
    }
  }, [requiredAndNotLoading, onUnauthenticated])

  if (requiredAndNotLoading) {
    return { data: value.data, status: `loading` } as const
  }

  return value
}

export async function getSession() {
  const res = await fetch(`/api/session`)
  const data = await res.json()
  if (!res.ok) throw data

  return data
}

/**
 * Provider to wrap the app in to make session data available globally.
 * Can also be used to throttle the number of requests to the endpoint
 * `/api/auth/session`.
 *
 * [Documentation](https://next-auth.js.org/getting-started/client#sessionprovider)
 */
export function SessionProvider(props: SessionProviderProps) {
  const { children } = props

  /**
   * If session was `null`, there was an attempt to fetch it,
   * but it failed, but we still treat it as a valid initial value.
   */

  const hasInitialSession = props.session !== undefined

  const [session, setSession] = React.useState(() => {
    if (hasInitialSession) return props.session
  })

  const { error, isValidating } = useSWR(`/api/auth/session`, fetcher, {
    onSuccess: (data) => {
      setSession(data)
    },
  })

  const loading = (!error && !session) || isValidating
  console.log(loading)
  const value: any = React.useMemo(
    () => ({
      data: session,
      status: loading
        ? `loading`
        : session
        ? `authenticated`
        : `unauthenticated`,
    }),
    [session, loading],
  )

  return (
    <SessionContext.Provider value={value}>{children}</SessionContext.Provider>
  )
}
