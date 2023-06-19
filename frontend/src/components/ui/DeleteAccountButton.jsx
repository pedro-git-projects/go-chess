import { useContext, useState } from "react"
import { useNavigate } from "react-router-dom"
import { TokenContext } from "../../contexts/AuthContext"

const DeleteAccount = () => {
  const [password, setPassword] = useState("")
  const [modalOpen, setModalOpen] = useState(false)
  const [errorMessage, setErrorMessage] = useState("")
  const { token, resetToken } = useContext(TokenContext)
  const navigate = useNavigate()

  const handlePasswordChange = (e) => {
    setPassword(e.target.value)
  }

  const handleDeleteUser = () => {
    setErrorMessage("")
    fetch("http://localhost:8080/delete", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: token,
      },
      body: JSON.stringify({ password }),
    })
      .then((response) => {
        if (response.ok) {
          resetToken()
          navigate("/")
        } else {
          setErrorMessage("Failed to delete user")
        }
      })
      .catch((error) => {
        console.error("Error deleting user:", error)
        setErrorMessage("An error occurred")
      })
  }

  const openModal = () => {
    setModalOpen(true)
  }

  const closeModal = () => {
    setModalOpen(false)
    setPassword("")
    setErrorMessage("")
  }

  return (
    <div style={{ display: "flex", justifyContent: "center" }}>
      <button
        className="block w-full rounded border w-full sm:w-auto bg-[#CE3262] border-[#CE3262] hover:bg-[#c9426d] border-[#c9426d] text-white font-bold py-2 px-4 rounded mb-4"
        onClick={openModal}
      >
        Delete Account
      </button>

      {/* Delete confirmation modal */}
      {modalOpen && (
        <div className="fixed inset-0 flex items-center justify-center z-50">
			<div className="bg-white dark:bg-gray-800 w-96 p-6 rounded-lg">
				<h2 className="text-lg font-bold mb-4 dark:text-white">Confirm Account Deletion</h2>
            <div className="mb-4">
				<label htmlFor="password" className="block font-medium mb-2 dark:text-white">
                Password
              </label>
              <input
                type="password"
                id="password"
                value={password}
                onChange={handlePasswordChange}
                className="border border-gray-300 px-3 py-2 rounded"
              />
            </div>
            {errorMessage && (
              <p className="text-red-500 mb-4">{errorMessage}</p>
            )}
            <div className="flex justify-end">
              <button
                className="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded mr-2"
                onClick={handleDeleteUser}
              >
                Delete
              </button>
              <button
                className="bg-gray-300 hover:bg-gray-400 text-gray-800 font-bold py-2 px-4 rounded"
                onClick={closeModal}
              >
                Cancel
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  )
}

export default DeleteAccount
