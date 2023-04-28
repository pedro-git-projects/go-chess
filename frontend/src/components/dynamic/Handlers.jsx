import sendMessage from "../../hooks/sendMessage"

export const handleJoinRoom = async (roomCode, ws, navigate, setResponse, setError) => {
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

export const handleCreateRoom = async (ws, navigate, setResponse) => {
  try {
    const message = JSON.stringify({message: "create"})
    const resp = await sendMessage(ws, message) 
    setResponse(resp)
    console.log("clientID: ", resp.client_id)
    navigate(`/room/${resp.room_id}`, { state: { roomID: resp.room_id, clientID: resp.client_id }})
  } catch(err) {
    console.log("Error: ", err)
  }
}
