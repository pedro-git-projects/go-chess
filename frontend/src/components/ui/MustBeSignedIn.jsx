const MustBeSignedIn = () => {
return (
    <div className="flex flex-col items-center justify-center h-screen">
      <h1 className="text-3xl font-bold mb-4">You must be signed in to play</h1>
      <a href="/signin" className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
        Sign In
      </a>
    </div>
  )
}

export default MustBeSignedIn
