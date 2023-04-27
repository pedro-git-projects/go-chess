import { useNavigate } from "react-router-dom"

const Hero = () => {
  const navigate = useNavigate()
  const handleClickPlay = () => navigate("/play")
  const handleClickLearn = () => navigate("/learn")
  return (
    <section className="text-white">
      <div
        className="mx-auto max-w-screen-xl px-4 py-32 lg:flex lg:h-screen lg:items-center"
      >
        <div className="mx-auto max-w-3xl text-center">
          <h1
            className="bg-gradient-to-r from-green-300 via-blue-500 to-purple-600 bg-clip-text text-3xl font-extrabold text-transparent sm:text-5xl"
          >
            Play Chess Online. 
            <span className="sm:block"> Have fun with your friends. </span>
          </h1>

          <p className="mx-auto mt-4 max-w-xl sm:text-xl/relaxed">
            <b>Go Chess</b> is a free online chess game powered by <b>Go</b> and <b>React</b>. 
          </p>

          <div className="mt-8 flex flex-wrap justify-center gap-4">
            <button
              className="block w-full rounded border border-blue-600 bg-blue-600 px-12 py-3 text-sm font-medium text-white hover:bg-transparent hover:text-white focus:outline-none focus:ring active:text-opacity-75 sm:w-auto"
              onClick={handleClickPlay} 
            >
              Go Play 
            </button>

            <button
              className="block w-full rounded border border-blue-600 px-12 py-3 text-sm font-medium text-white hover:bg-blue-600 focus:outline-none focus:ring active:bg-blue-500 sm:w-auto"
              onClick={handleClickLearn}
            >
              Go Learn
            </button>
          </div>
        </div>
      </div>
    </section>
  )
}

export default Hero
