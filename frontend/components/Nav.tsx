import Link from 'next/link'
import { useSession } from '../session'
import UserSection from './UserSection'

const Nav = () => {
  const { data: session } = useSession()
  return (
    <nav className="w-full py-3 max-w-4xl px-4 items-center mx-auto flex justify-between">
      <div>
        {` `}
        <span className="text-2xl font-bold">
          <Link href="/">Teardrop</Link>
        </span>
      </div>
      <UserSection />
    </nav>
  )
}

export default Nav
