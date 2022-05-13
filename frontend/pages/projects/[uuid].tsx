import { useRouter } from 'next/router'
import useSWR from 'swr'
import Nav from '../../components/Nav'
import { fetcher } from '../../utils'

const ProjectPage = () => {
  const { uuid } = useRouter().query
  const { data: project } = useSWR<{ name: string; git: string }>(
    `/api/projects/${uuid}`,
    fetcher,
  )

  return (
    <>
      <Nav />
      <main>
        <div className="max-w-4xl mx-auto mt-8 px-4 flex justify-between items-center">
          <div>
            <h1 className="text-3xl font-bold">{project?.name}</h1>
            <p className="text-sm mt-4 text-gray-600">
              Deploys from {project?.git}
            </p>
          </div>
          <div className='flex gap-2'>

            <button className="rounded-lg px-6 py-2 text-sm border">Settings</button>
            <button className="rounded-lg px-6 py-2 text-sm border bg-black text-white border-white">Visit</button>
          </div>
        </div>
        <hr className="mt-4" />
      </main>
    </>
  )
}
export default ProjectPage
