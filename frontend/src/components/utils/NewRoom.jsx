import { useNavigate } from 'react-router-dom'
import { useState } from 'react'
import SendMessage from '../websockets/SendMessage'
import ConnectToWS from '../websockets/ConnectToWS'

const NewRoom = () => {
  const navigate = useNavigate()
  const [message, setMessage] = useState("")
  const [response, setResponse] = useState("")
  const handleClick = async () => {
    try {
      const ws = await ConnectToWS("ws://localhost:8080/create-room")
      const message = JSON.stringify({message: "create"})
      const resp = await SendMessage(ws, message)
      setResponse(resp)
      setMessage("")
      ws.close()
    } catch(err) {
      console.log("Error: ", err)
    }
    console.log(resp)

  }
  return (
    <>
      <button type="button" onClick={handleClick}>Create room</button>
    </>
  )
}

export default NewRoom
