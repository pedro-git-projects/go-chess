import sendMessage from "../../hooks/sendMessage"

export const handleJoinRoom = async (
  roomCode,
  ws,
  navigate,
  setResponse,
  setError,
) => {
  try {
    const message = JSON.stringify({ message: "join", room_id: roomCode })
    const resp = await sendMessage(ws, message)
    console.log("response received:", resp)
    setResponse(resp)
    console.log("Join room response: ", resp)
    if (resp.room_id === "") {
      setError(resp.error)
    } else {
      navigate(`/room/${resp.room_id}`, {
        state: { 
          roomID: resp.room_id, 
          clientID: resp.client_id, 
          clientColor: resp.client_color,
        },
      })
    }
  } catch (err) {
    console.log("Error: ", err)
  }
}

export const handleCreateRoom = async (ws, navigate, setResponse) => {
  try {
    const message = JSON.stringify({ message: "create" })
    const resp = await sendMessage(ws, message)
    setResponse(resp)
    navigate(`/room/${resp.room_id}`, {
      state: { roomID: resp.room_id, clientID: resp.client_id, clientColor: resp.client_color },
    })
  } catch (err) {
    console.log("Error: ", err)
  }
}

export const handleUpdateRoom = (ws, onUpdateRoom) => {
  ws.addEventListener("message", (e) => {
    const data = JSON.parse(e.data)
    console.log("response received: ", data)
    if(data.type == "room-update") {
      onUpdateRoom(data.number_of_clients_in_room);
    }
  })
}
