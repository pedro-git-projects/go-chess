import { useNavigate } from "react-router-dom"

const MustBeSignedIn = () => {
	const navigate = useNavigate()
	const handleClick = () => navigate("/signin")
return (
    <div className="flex flex-col items-center justify-center h-screen">
      <h1 className="text-3xl font-bold mb-4">You must be signed in to play</h1>
      <div onClick={handleClick} href="/signin" className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
        Sign In
      </div>
    </div>
  )
}

export default MustBeSignedIn
