import { useState, useContext } from "react"
import { TokenContext } from "../../contexts/AuthContext"
import { useNavigate } from "react-router-dom"

const SignUpForm = ({ toggleSignUpForm }) => {
  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")
  const [usernameError, setUsernameError] = useState("")
  const [error, setError] = useState("")
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
        // Make a login request to obtain the token
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
          className="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4"
        >
          {error && <p className="text-red-500 mb-4">{error}</p>}

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
              Sign Up
            </button>
            <button
              className="text-blue-500 hover:text-blue-700 font-bold"
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
          <div className="bg-white rounded-lg p-6 shadow-lg">
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
