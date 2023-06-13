import React, { useState } from "react"

const SignInForm = () => {
  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")
  const [usernameError, setUsernameError] = useState("")

  const handleUsernameChange = (e) => {
    setUsername(e.target.value)
    setUsernameError("")
  }

  const handlePasswordChange = (e) => {
    setPassword(e.target.value)
  }

  const handleSubmit = async (e) => {
    e.preventDefault()

    if (!username) {
      setUsernameError("Please enter a username")
      return
    }

    const loginData = {
      username,
      password,
    }

    try {
      const response = await fetch("http://localhost:8080/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(loginData),
      })

	console.log(response)

      if (response.ok) {
      	console.log("Response Headers:", response.headers)
        const token = response.headers.get("Authorization")
        console.log("Token:", token)
        // Do something with the token (e.g., store it in local storage)
      } else {
        const errorText = await response.text()
        console.log("Login failed:", errorText)
      }
    } catch (error) {
      console.log("An error occurred:", error)
    }
  }

  return (
    <div className="flex flex-col items-center justify-center min-h-screen">
      <div className="w-full max-w-md">
        <form
          onSubmit={handleSubmit}
          className="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4"
        >
          <div className="mb-8">
            <label
              className="block text-gray-700 font-bold mb-2"
              htmlFor="username"
            >
              Username
            </label>
            <input
              className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              id="username"
              type="text"
              placeholder="Username"
              value={username}
              onChange={handleUsernameChange}
            />
            {usernameError && (
              <p className="text-red-500 text-xs italic">{usernameError}</p>
            )}
          </div>
          <div className="mb-8">
            <label
              className="block text-gray-700 font-bold mb-2"
              htmlFor="password"
            >
              Password
            </label>
            <input
              className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              id="password"
              type="password"
              placeholder="********"
              value={password}
              onChange={handlePasswordChange}
            />
          </div>
          <div className="flex items-center justify-between">
            <button
              className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
              type="submit"
            >
              Sign In
            </button>
          </div>
        </form>
      </div>
    </div>
  )
}

export default SignInForm
