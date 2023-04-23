import { useNavigate } from 'react-router-dom'
import { useState } from 'react'
import { useWebSocket } from '../../contexts/WebSocketContext'
import sendMessage from "../../hooks/sendMessage"

const NewRoom = () => {
  const navigate = useNavigate()
  const [response, setResponse] = useState("")
  const ws = useWebSocket() 

  const handleClick = async () => {
    try {
      const message = JSON.stringify({message: "create"})
      const resp = await sendMessage(ws, message) 
      setResponse(resp)
      console.log("clientID: ", resp.client_id)
      console.log("turn: ", resp.turn)
      navigate(`/room/${resp.room_id}`, { state: { roomID: resp.room_id, clientID: resp.client_id, turn:  resp.turn}})
    } catch(err) {
      console.log("Error: ", err)
    }
  }
  return (
    <>
      <button type="button" onClick={handleClick}>Create room</button>
    </>
  )
}

export default NewRoom
