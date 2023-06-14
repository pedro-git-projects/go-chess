import { useContext, useEffect } from "react"
import { TokenContext } from "../../contexts/AuthContext"
import { useNavigate } from "react-router-dom"
import { Layout } from "../ui/Layout"
import DeleteAccount from "../ui/DeleteAccountButton"
import ChangePasswordForm from "../ui/ChangePasswordForm"

const Account = () => {
  const navigate = useNavigate()
  const { token } = useContext(TokenContext)

  useEffect(() => {
    const userIsLoggedIn = token !== null
    if (!userIsLoggedIn) {
      navigate("/signin")
    }
  }, [token, navigate])

  if (token === null) {
    return null
  }

  return (
    <Layout>
      <h1 className="text-3xl font-bold underline text-center p-5"></h1>
      <div>
        <ChangePasswordForm />
      </div>
      <hr className="my-12 h-px border-t-0 bg-transparent bg-gradient-to-r from-transparent via-neutral-500 to-transparent opacity-25 dark:opacity-100" />
      <div className="mt-4">
        <DeleteAccount />
      </div>
    </Layout>
  )
}

export default Account
