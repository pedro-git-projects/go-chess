import { useRouteError } from "react-router-dom"

const Err = () => {
  const error = useRouteError()
  console.error(error)
  return (
    <div className="flex h-screen" id="error-page">
      <div className="m-auto">
        <h1 className="text-3xl mb-4 font-bold text-center">Oops!</h1>
        <p className="text-2xl mb-4 text-center">
          Sorry, an unexpected error has occurred.
        </p>
        <p className="text-center text-xl">
          <i>{error.statusText || error.message}</i>
        </p>
      </div>
    </div>
  )
}

export default Err
