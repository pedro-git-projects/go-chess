import { useState, useContext } from "react"
import { Link, useLocation, useNavigate } from "react-router-dom"
import { TokenContext } from "../../contexts/AuthContext"

const SignBtn = ({ className }) => {
  const { token, resetToken } = useContext(TokenContext)
  const location = useLocation()
  const navigate = useNavigate()
  const [showConfirmation, setShowConfirmation] = useState(false)

  const handleSignOut = async () => {
    try {
      await fetch("http://localhost:8080/signout", {
        method: "POST",
        headers: {
          Authorization: `${token}`,
        },
      })
      resetToken()
      navigate("/")
    } catch (error) {
      console.error("Sign out failed:", error)
    }
  }

  const handleClick = () => {
    setShowConfirmation(true)
  }

  const handleCancel = () => {
    setShowConfirmation(false)
  }

  const handleConfirmSignOut = () => {
    handleSignOut()
    setShowConfirmation(false)
  }

  return (
    <>
      {token ? (
        <>
          <Link
            to={location.pathname}
            className={`${className}`}
            onClick={handleClick}
          >
            Sign out
          </Link>
          {showConfirmation && (
            <div className="fixed inset-0 flex items-center justify-center z-50">
				<div className="bg-white dark:bg-gray-800 w-64 p-4 rounded-lg shadow-lg">
					<p className="mb-4 dark:text-white">Are you sure you want to sign out?</p>
                <div className="flex justify-end">
                  <button
					  className="mr-2 px-4 py-2 text-sm font-medium text-gray-600 dark:text-white rounded-lg border border-gray-300 hover:bg-gray-100 dark:hover:bg-gray-900"
                    onClick={handleCancel}
                  >
                    Cancel
                  </button>
                  <button
					  className="px-4 py-2 text-sm font-medium text-red-600 rounded-lg border border-red-500 hover:bg-red-100 dark:hover:bg-gray-900"
                    onClick={handleConfirmSignOut}
                  >
                    Sign out
                  </button>
                </div>
              </div>
            </div>
          )}
        </>
      ) : (
        <Link to="/signin" className={`${className}`}>
          Sign in
        </Link>
      )}
    </>
  )
}

export default SignBtn
