import { useState, useContext } from "react"
import { TokenContext } from "../../contexts/AuthContext"
import { useNavigate } from "react-router-dom"
import { VscEye, VscEyeClosed } from "react-icons/vsc"

const SignUpForm = ({ toggleSignUpForm }) => {
  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")
  const [usernameError, setUsernameError] = useState("")
  const [error, setError] = useState("")
  const [showPassword, setShowPassword] = useState(false)
  const [showModal, setShowModal] = useState(false)
  const tokenContext = useContext(TokenContext)
  const navigate = useNavigate()

  const handleUsernameChange = (e) => {
    setUsername(e.target.value)
    setUsernameError("")
    setError("")
  }

  const handlePasswordChange = (e) => {
    setPassword(e.target.value)
    setError("")
  }

  const togglePasswordVisibility = () => {
    setShowPassword((prevState) => !prevState)
  }

  const handleSubmit = async (e) => {
    e.preventDefault()

    if (!username) {
      setUsernameError("Please enter a username")
      return
    }

    const registrationData = {
      username,
      password,
    }

    try {
      const registrationUrl = "http://localhost:8080/register"
      const loginUrl = "http://localhost:8080/login"
      const registrationResponse = await fetch(registrationUrl, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(registrationData),
      })

      if (registrationResponse.ok) {
        const loginData = {
          username,
          password,
        }
        const loginResponse = await fetch(loginUrl, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(loginData),
        })

        if (loginResponse.ok) {
          console.log("Login successful")
          const token = loginResponse.headers.get("Authorization")
          tokenContext.setToken(token)
          setShowModal(true)
          setTimeout(() => {
            setShowModal(false) // Close the modal after a delay
            navigate("/") // Navigate after the modal is closed
          }, 5000) // Set the delay to 3000 milliseconds (5 seconds)
        } else {
          const errorText = await loginResponse.text()
          setError(errorText)
        }
      } else {
        const errorText = await registrationResponse.text()
        setError(errorText)
      }
    } catch (error) {
      console.log("An error occurred:", error)
    }
  }

  const closeModal = () => {
    setShowModal(false)
    navigate("/") // Navigate after the modal is closed
  }

  return (
    <div className="flex flex-col items-center justify-center min-h-screen">
      <div className="w-full max-w-md">
        <form
          onSubmit={handleSubmit}
			className="bg-slate-200 dark:bg-gray-800 shadow-md rounded px-8 pt-6 pb-8 mb-4"
        >
          {error && <p className="text-red-500 mb-4">{error}</p>}

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
                className="absolute top-2 right-2 cursor-pointer"
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
              Sign Up
            </button>
            <button
              className="text-[#00ADD8] hover:text-[#00A29C] font-bold"
              onClick={toggleSignUpForm}
            >
              Back to Sign In
            </button>
          </div>
        </form>
      </div>
      {/* Modal */}
      {showModal && (
        <div className="fixed inset-0 flex items-center justify-center z-50 bg-black bg-opacity-50">
          <div className="bg-slate-200 rounded-lg p-6 shadow-lg">
            <h2 className="text-xl mb-4">Successful Registration</h2>
            <p className="text-gray-800">
              Congratulations! Your registration was successful.
            </p>
            <button
              className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded mt-4"
              onClick={closeModal}
            >
              OK
            </button>
          </div>
        </div>
      )}
    </div>
  )
}

export default SignUpForm
