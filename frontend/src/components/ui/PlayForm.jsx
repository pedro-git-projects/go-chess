import React from "react"
import { useNavigate } from "react-router-dom"
import { useState } from "react"
import { useWebSocket } from "../../contexts/WebSocketContext"
import { handleJoinRoom, handleCreateRoom } from "../dynamic/Handlers"

const PlayForm = () => {
  const navigate = useNavigate()
  const [roomCode, setRoomCode] = useState("")
  const [error, setError] = useState("")
  const [response, setResponse] = useState(null)
  const ws = useWebSocket()
  const handleJoinRoomSubmit = async (e) => {
    e.preventDefault()
    await handleJoinRoom(roomCode, ws, navigate, setResponse, setError)
  }
  const handleRoomCodeChange = (e) => {
    console.log(e.target.value)
    setRoomCode(e.target.value)
  }
  const handleNewRoomClick = async () => {
    await handleCreateRoom(ws, navigate, setResponse)
  }
  return (
    <div className="flex items-center justify-center h-screen">
      <div
        className="bg-[#DBD9D6] dark:bg-[#555759] rounded-lg shadow-md p-6 flex flex-col items-center 
        sm:w-1/2 md:w-2/3 lg:w-1/2 xl:w-1/3"
      >
        <button
          onClick={handleNewRoomClick}
          className="block w-full rounded border w-full sm:w-auto bg-[#CE3262] border-[#CE3262] hover:bg-[#c9426d] border-[#c9426d] text-white font-bold py-2 px-4 rounded mb-4"
        >
          Create Room
        </button>
        <form onSubmit={handleJoinRoomSubmit} className="w-full space-y-4">
          <div className="flex items-center">
            <input
              placeholder="Enter room code:"
              type="text"
              id="roomCode"
              value={roomCode}
              onChange={handleRoomCodeChange}
              className={`border ${
                error ? "bg-red-200" : "border-gray-300"
              } px-3 py-2 rounded-md w-full text-sm focus:outline-none focus:ring-2 focus:ring-[#00A29C]`}
            />
            <button
              type="submit"
              className="block w-full rounded border border-[#00A29C] bg-[#00A29C] hover:bg-[#3ba8a4] text-white font-bold py-1 px-4 rounded ml-2"
            >
              Join
            </button>
          </div>
        </form>
        {error && <p className="text-red-600">{error}</p>}
      </div>
    </div>
  )
}

export default PlayForm
