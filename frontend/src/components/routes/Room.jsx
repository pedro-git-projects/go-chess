import { useLocation, useNavigate } from "react-router-dom"
import ChessBoard from "../game/ChessBoard"
import { Layout } from "../ui/Layout"
import { useEffect } from "react"

const Room = () => {
  const location = useLocation()
  const navigate = useNavigate()
  const clientID = location.state?.clientID || ""
  useEffect(() => {
    if (!clientID) {
      navigate("/play")
    }
  }, [clientID, navigate])

  if (!clientID) {
    return null
  }
  return (
    <Layout>
      <div className="w-full h-full flex align-middle items-center justify-center">
        <div className="mx-auto">
          <ChessBoard clientID={clientID}></ChessBoard>
        </div>
      </div>
    </Layout>
  )
}

export default Room

