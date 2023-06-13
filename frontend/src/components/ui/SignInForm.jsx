import { useState, useContext } from "react"
import { TokenContext } from "../../contexts/AuthContext"
import SignUpForm from "./SignUpForm"

const SignInForm = () => {
  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")
  const [usernameError, setUsernameError] = useState("")
  const tokenContext = useContext(TokenContext)
  const [showSignUpForm, setShowSignUpForm] = useState(false)

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
      const url = "http://localhost:8080/login"
      const response = await fetch(url, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(loginData),
      })

      if (response.ok) {
        console.log("Response Headers:", response.headers)
        const token = response.headers.get("Authorization")
        console.log("Token:", token)
        tokenContext.setToken(token)
      } else {
        const errorText = await response.text()
        console.log("Login failed:", errorText)
      }
    } catch (error) {
      console.log("An error occurred:", error)
    }
  }

  const toggleSignUpForm = () => {
    setShowSignUpForm(!showSignUpForm)
  }

  return (
    <div className="flex flex-col items-center justify-center min-h-screen">
      <div className="w-full max-w-md">
        {showSignUpForm ? (
          <SignUpForm toggleSignUpForm={toggleSignUpForm} />
        ) : (
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
              <button
                className="text-blue-500 hover:text-blue-700 font-bold"
                type="button"
                onClick={toggleSignUpForm}
              >
                Sign Up
              </button>
            </div>
          </form>
        )}
      </div>
    </div>
  )
}

export default SignInForm
