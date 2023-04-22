import { useLocation, useNavigate } from "react-router-dom"
import ChessBoard from "../game/ChessBoard"
import { Layout } from "../ui/Layout"
import { useEffect } from "react"
import { useWebSocket } from '../../contexts/WebSocketContext'

const Room = () => {
  const location = useLocation()
  const navigate = useNavigate()
  const roomID = location.state?.roomID || ""
  const clientID = location.state?.clientID || ""
  const ws = useWebSocket()
  useEffect(() => {
    if (!roomID) {
      navigate("/play")
    }
  }, [roomID, navigate])
  if (!roomID || !ws) {
    return null
  }
  return (
    <Layout>
      <div className="w-full h-full flex align-middle items-center justify-center">
        <div className="mx-auto">
          <ChessBoard roomID={roomID} clientID={clientID}></ChessBoard>
        </div>
      </div>
    </Layout>
  )
}

export default Room

