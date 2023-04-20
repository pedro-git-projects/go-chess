import { useNavigate } from 'react-router-dom'
import { useState } from 'react'
import connectToWS from "../../hooks/connectToWS"
import sendMessage from "../../hooks/sendMessage"

const NewRoom = () => {
  const navigate = useNavigate()
  const [response, setResponse] = useState("")
  const handleClick = async () => {
    try {
      const ws = await connectToWS("ws://localhost:8080/create-room")
      const message = JSON.stringify({message: "create"})
      const resp = await sendMessage(ws, message)
      setResponse(resp)
      ws.close()
      console.log("clientID: ", resp.client_id)
      navigate(`/room/${resp.room_id}`, { state: {roomID: resp.room_id}})    
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
