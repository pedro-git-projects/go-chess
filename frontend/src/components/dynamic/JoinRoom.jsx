
import { useState } from "react"
import { useNavigate } from "react-router-dom"
import {useWebSocket} from "../../contexts/WebSocketContext"
import connectToWS from "../../hooks/connectToWS"
import sendMessage from "../../hooks/sendMessage"

const JoinRoom = () => {
  const navigate = useNavigate()
  const [roomCode, setRoomCode] = useState("")
  const [error, setError] = useState("")
  const [response, setResponse] = useState(null) 
  const ws = useWebSocket()

  const handleJoinRoom = async (e) => {
    e.preventDefault()
    try {
      const message = JSON.stringify({message: "join", room_id: roomCode})    
      const resp = await sendMessage(ws, message)
      console.log("response received:", resp)
      setResponse(resp)
      if (resp.room_id === "") {
        setError(resp.error)
      } else {
        console.log("clientID: ", resp.client_id)
        navigate(`/room/${resp.room_id}`, { state: { roomID: resp.room_id, clientID: resp.client_id }})
      }
    } catch(err) {
      console.log("Error: ", err)
    }
  }
  const handleRoomCodeChange = (e) => {
    console.log(e.target.value)
    setRoomCode(e.target.value)
  }
  return (
    <div>
      <form onSubmit={handleJoinRoom}>
        <label htmlFor="roomCode">Enter room code:</label>
        <input type="text" id="roomCode" value={roomCode} onChange={handleRoomCodeChange} />
        <button type="submit">Join room</button>
      </form>
      {error && <p>{error}</p>}
    </div>
  )
}

export default JoinRoom

