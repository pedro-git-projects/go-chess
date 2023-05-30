import { useLocation, useNavigate } from "react-router-dom"
import ChessBoard from "../game/ChessBoard"
import { Layout } from "../ui/Layout"
import { useState, useEffect } from "react"
import { useWebSocket } from "../../contexts/WebSocketContext"
import whiteKing from "../../assets/white_king.svg"
import blackKing from "../../assets/black_king.svg"
import { handleUpdateRoom } from "../dynamic/Handlers"
import CopyToClipboard from "../ui/CopyToClipboard"

const Room = () => {
  const location = useLocation()
  const navigate = useNavigate()
  const roomID = location.state?.roomID || ""
  const clientID = location.state?.clientID || ""
  const clientColor = location.state?.clientColor || ""
  const [turn, setTurn] = useState(location.state?.turn || "white")
  const [numberOfClientsInRoom, setNumberOfClientsInRoom] = useState(
    location.state?.numberOfClientsInRoom || 1,
  )
  const ws = useWebSocket()

  useEffect(() => {
    if (!roomID) {
      navigate("/play")
    }
    handleUpdateRoom(ws, (numberOfClientsInRoom) => {
      setNumberOfClientsInRoom(numberOfClientsInRoom)
    })

    const handleBeforeUnload = (event) => {
      event.preventDefault()
      event.returnValue = "" // Required for Chrome
      navigate("/play")
    }

    window.addEventListener("beforeunload", handleBeforeUnload)

    return () => {
      window.removeEventListener("beforeunload", handleBeforeUnload)
    }
  }, [roomID, navigate, ws])

  const handleTurnUpdate = (newTurn) => {
    setTurn(newTurn)
  }

  if (!roomID || !ws) {
    return null
  }

  if (numberOfClientsInRoom < 2) {
    return (
      <Layout>
        <div className="min-h-screen flex justify-center items-center bg-gray-100">
          <div className="text-center max-w-md w-full bg-white shadow-md rounded-lg p-8">
            <p className="mb-4">Waiting for another player to join...</p>
            <p className="mb-4">Send them the room code!</p>
            <div className="relative">
              <CopyToClipboard className="w-full" />
            </div>
          </div>
        </div>
      </Layout>
    )
  }
  return (
    <Layout>
      {turn === clientColor ? (
        <h2 className="text-black dark:text-white text-3xl font-bold text-center py-3">
          Your turn
        </h2>
      ) : (
        <h2 className="text-black dark:text-white text-3xl font-bold text-center py-3">
          Opponent's turn
        </h2>
      )}
      <div className="w-full h-full flex align-middle items-center justify-center">
        <div className="mx-auto">
          <ChessBoard
            roomID={roomID}
            clientID={clientID}
            clientColor={clientColor}
            turn={turn}
            onTurnUpdate={handleTurnUpdate}
          ></ChessBoard>
        </div>
      </div>
      <h2 className="text-black dark:text-white text-3xl font-bold text-center py-3">
        <span className="inline-block mx-2">
          {clientColor === "black" ? (
            <img src={blackKing} alt="black king"></img>
          ) : (
            <img src={whiteKing} alt="white king"></img>
          )}
        </span>
        You're playing as {clientColor}
      </h2>
    </Layout>
  )
}

export default Room
