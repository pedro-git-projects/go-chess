import { useLocation, useNavigate } from "react-router-dom"
import ChessBoard from "../game/ChessBoard"
import { Layout } from "../ui/Layout"
import { useState,useEffect } from "react"
import ConnectToWS from "../websockets/ConnectToWS"

const Room = () => {
  const location = useLocation()
  const navigate = useNavigate()
  const roomID = location.state?.roomID || ""
  useEffect(() => {
    if (!roomID) {
      navigate("/play")
    }
  }, [roomID, navigate])

  const [ws, setWs] = useState(null);
  useEffect(() => {
    async function connectToWebSocket() {
      const websocket = await ConnectToWS("ws://localhost:8080/board");
      setWs(websocket);
    }
    connectToWebSocket();
  }, []);

  if (!roomID || !ws) {
    return null
  }
  return (
    <Layout>
      <div className="w-full h-full flex align-middle items-center justify-center">
        <div className="mx-auto">
          <ChessBoard roomID={roomID} ws={ws}></ChessBoard>
        </div>
      </div>
    </Layout>
  )
}

export default Room

