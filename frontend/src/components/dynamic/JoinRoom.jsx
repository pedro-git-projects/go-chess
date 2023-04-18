import { useState } from "react"
import { useNavigate } from "react-router-dom"
import ConnectToWS from "../websockets/ConnectToWS"
import SendMessage from "../websockets/SendMessage"

const JoinRoom = () => {
  const navigate = useNavigate()
  const [roomCode, setRoomCode] = useState("")
  const [error, setError] = useState("")
  const [response, setResponse] = useState(null) 

  const handleJoinRoom = async (e) => {
    e.preventDefault()
    try {
      const ws = await ConnectToWS("ws://localhost:8080/join-room")
      console.log("WebSocket connection established.")
      const message = JSON.stringify({message: "join", room_id: roomCode})    
      const resp = await SendMessage(ws, message)
      console.log("response received:", resp)
      setResponse(resp)
      if (resp.room_id === "") {
        setError(resp.error)
      } else {
        console.log("clientID: ", resp.client_id)
        navigate(`/room/${resp.room_id}`, { state: {roomID: resp.room_id}})
      }
      ws.close()
    } catch(err) {
      console.log("Error: ", err)
    }
  }
  const handleRoomCodeChange = (e) => {
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
