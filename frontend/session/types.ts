export interface UseSessionOptions<R extends boolean> {
  required: R
  /** Defaults to `signIn` */
  onUnauthenticated?: () => void
}

export type SessionContextValue<R extends boolean = false> = R extends true
  ?
      | { data: Session; status: 'authenticated' }
      | { data: null; status: 'loading' }
  :
      | { data: Session; status: 'authenticated' }
      | { data: null; status: 'unauthenticated' | 'loading' }

export interface DefaultSession extends Record<string, unknown> {
  user?: {
    name?: string | null
    email?: string | null
    image?: string | null
    admin?: boolean | null
  }
}

export interface Session extends Record<string, unknown>, DefaultSession {}

export interface SessionProviderProps {
  children: React.ReactNode
  session?: Session | null
  baseUrl?: string
  basePath?: string
  /**
   * A time interval (in seconds) after which the session will be re-fetched.
   * If set to `0` (default), the session is not polled.
   */
  refetchInterval?: number
  /**
   * `SessionProvider` automatically refetches the session when the user switches between windows.
   * This option activates this behaviour if set to `true` (default).
   */
  refetchOnWindowFocus?: boolean
}
