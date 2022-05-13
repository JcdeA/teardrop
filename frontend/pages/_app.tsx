import '../styles/globals.css'
import type { AppProps } from 'next/app'
import { SessionProvider } from '../session'

function MyApp({ Component, pageProps: { session, ...pageProps } }: AppProps) {
  return (
    <SessionProvider session={session}>
      <Component {...pageProps} />
    </SessionProvider>
  )
}

export default MyApp
