import { useState, useContext } from "react"
import { TokenContext } from "../../contexts/AuthContext"

const ChangePasswordForm = () => {
  const [oldPassword, setOldPassword] = useState("")
  const [newPassword, setNewPassword] = useState("")
  const [errorMessage, setErrorMessage] = useState("")
  const [successMessage, setSuccessMessage] = useState("") // Added successMessage state
  const { token } = useContext(TokenContext)

  const handleChangePassword = async () => {
    setErrorMessage("")
    setSuccessMessage("") // Reset success message

    try {
      const response = await fetch("http://localhost:8080/change-password", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: token,
        },
        body: JSON.stringify({
          old_password: oldPassword,
          new_password: newPassword,
        }),
      })
      console.log("Response status:", response.status)
      console.log("Response data:", response.data)

      if (!response.ok) {
        const errorData = await response.json()
        throw new Error(errorData.error)
      }

      // Password change successful
      setSuccessMessage("Password changed successfully!") // Set success message
    } catch (error) {
      setErrorMessage("Couldn't change password")
    }
  }

  return (
    <div className="max-w-md mx-auto">
      <h2 className="text-2xl font-bold mb-4">Change Password</h2>
      {errorMessage && <div className="text-red-500 mb-4">{errorMessage}</div>}
      {successMessage && (
        <div className="text-green-500 mb-4">{successMessage}</div>
      )}{" "}
      {/* Render success message */}
      <div className="mb-4">
        <label htmlFor="oldPassword" className="block font-medium mb-1">
          Old Password
        </label>
        <input
          type="password"
          id="oldPassword"
          className="w-full border rounded py-2 px-3"
          value={oldPassword}
          onChange={(event) => setOldPassword(event.target.value)}
        />
      </div>
      <div className="mb-4">
        <label htmlFor="newPassword" className="block font-medium mb-1">
          New Password
        </label>
        <input
          type="password"
          id="newPassword"
          className="w-full border rounded py-2 px-3"
          value={newPassword}
          onChange={(event) => setNewPassword(event.target.value)}
        />
      </div>
      <button
        type="button"
        className="bg-blue-500 text-white py-2 px-4 rounded hover:bg-blue-600"
        onClick={handleChangePassword}
      >
        Change Password
      </button>
    </div>
  )
}

export default ChangePasswordForm
