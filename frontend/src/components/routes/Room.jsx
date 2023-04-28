import { useLocation, useNavigate } from "react-router-dom"
import ChessBoard from "../game/ChessBoard"
import { Layout } from "../ui/Layout"
import { useState, useEffect } from "react"
import { useWebSocket } from '../../contexts/WebSocketContext'

const Room = () => {
  const location = useLocation()
  const navigate = useNavigate()
  const roomID = location.state?.roomID || ""
  const clientID = location.state?.clientID || ""
  console.log('location.state?.turn:', location.state?.turn)
  const [turn, setTurn] = useState(location.state?.turn || "white")  
  const ws = useWebSocket()
  useEffect(() => {
    if (!roomID) {
      navigate("/play")
    }
  }, [roomID, navigate])
  const handleTurnUpdate = (newTurn) => {
    setTurn(newTurn)
  }
  if (!roomID || !ws) {
    return null
  }
  return (
    <Layout>
      {turn === "white" ? (<h2 className="text-gray-500 dark:text-[#DDE6EB] text-3xl font-bold text-center py-3">{`${turn}'s turn`}</h2>) : (<h2 className="text-black dark:text-[#93AFC0] text-3xl font-bold text-center py-3">{`${turn}'s turn`}</h2>) 
 }
      <div className="w-full h-full flex align-middle items-center justify-center">
        <div className="mx-auto">
          <ChessBoard roomID={roomID} clientID={clientID} turn={turn} onTurnUpdate={handleTurnUpdate}></ChessBoard>
        </div>
      </div>
    </Layout>
  )
}

export default Room

