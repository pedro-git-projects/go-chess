import { useContext } from "react"
import { Link, useLocation, useNavigate } from "react-router-dom"
import { TokenContext } from "../../contexts/AuthContext"

const SignBtn = ({ className }) => {
  const { token, resetToken } = useContext(TokenContext)
  const location = useLocation()
  const navigate = useNavigate()

  const handleSignOut = async () => {
    try {
      await fetch("http://localhost:8080/signout", {
        method: "POST",
        headers: {
          Authorization: `${token}`
        }
      })
      resetToken()
      navigate("/")
    } catch (error) {
      console.error("Sign out failed:", error)
    }
  }

  return (
    <>
      {token ? (
        <Link
          to={location.pathname}
          className={`${className}`}
          onClick={handleSignOut}
        >
          Sign out
        </Link>
      ) : (
        <Link
          to="/signin"
          className={`${className}`}
        >
          Sign in
        </Link>
      )}
    </>
  )
}

export default SignBtn
