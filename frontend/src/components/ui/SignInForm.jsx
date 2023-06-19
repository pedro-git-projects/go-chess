import { useState, useContext } from "react"
import { useNavigate } from "react-router-dom"
import { TokenContext } from "../../contexts/AuthContext"
import { VscEye, VscEyeClosed } from "react-icons/vsc"
import SignUpForm from "./SignUpForm"

const SignInForm = () => {
  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")
  const [usernameError, setUsernameError] = useState("")
  const [showPassword, setShowPassword] = useState(false)
  const [loginError, setLoginError] = useState("")
  const tokenContext = useContext(TokenContext)
  const navigate = useNavigate()
  const [showSignUpForm, setShowSignUpForm] = useState(false)

  const handleUsernameChange = (e) => {
    setUsername(e.target.value)
    setUsernameError("")
    setLoginError("")
  }

  const handlePasswordChange = (e) => {
    setPassword(e.target.value)
    setLoginError("")
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
        navigate("/")
      } else {
        setLoginError("Login failed. Please check your credentials.")
      }
    } catch (error) {
      setLoginError("An error occurred. Please try again.")
    }
  }

  const toggleSignUpForm = () => {
    setShowSignUpForm(!showSignUpForm)
  }

  const togglePasswordVisibility = () => {
    setShowPassword(!showPassword)
  }

  return (
    <div className="flex flex-col items-center justify-center min-h-screen">
      <div className="w-full max-w-md">
        {showSignUpForm ? (
          <SignUpForm toggleSignUpForm={toggleSignUpForm} />
        ) : (
          <form
            onSubmit={handleSubmit}
			  className="bg-slate-200 dark:bg-gray-800 shadow-md rounded px-8 pt-6 pb-8 mb-4"
          >
            {loginError && (
              <p className="text-red-500 text-xs italic mb-4">{loginError}</p>
            )}
            <div className="mb-8">
              <label
				  className="block text-gray-700 dark:text-white font-bold mb-2"
                htmlFor="username"
              >
                Username
              </label>
              <input
                className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                id="username"
                type="text"
                value={username}
                onChange={handleUsernameChange}
              />
              {usernameError && (
                <p className="text-red-500 text-xs italic">{usernameError}</p>
              )}
            </div>
            <div className="mb-8">
              <label
				  className="block text-gray-700 dark:text-white font-bold mb-2"
                htmlFor="password"
              >
                Password
              </label>
              <div className="relative">
                <input
                  className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                  id="password"
                  type={showPassword ? "text" : "password"}
                  value={password}
                  onChange={handlePasswordChange}
                />
                <div
                  className="absolute inset-y-0 right-0 flex items-center pr-2 cursor-pointer"
                  onClick={togglePasswordVisibility}
                >
                  {showPassword ? <VscEye /> : <VscEyeClosed />}
                </div>
              </div>
            </div>
            <div className="flex items-center justify-between">
              <button
                className="rounded bg-[#00ADD8] hover:bg-[#5dc9e2] text-white font-bold py-1 px-4 rounded ml-2"
                type="submit"
              >
                Sign In
              </button>
              <button
                className="text-[#00ADD8] hover:text-[#00A29C] font-bold"
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
