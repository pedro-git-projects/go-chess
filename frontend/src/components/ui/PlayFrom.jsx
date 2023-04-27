import React from 'react'
import { useNavigate } from 'react-router-dom'
import { useState } from 'react'
import { useWebSocket } from '../../contexts/WebSocketContext'
import { handleJoinRoom, handleCreateRoom } from '../dynamic/Handlers'

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
      <div className="bg-white rounded-lg shadow-md p-6 flex flex-col items-center">
        <button
          onClick={handleNewRoomClick}
          className="w-full sm:w-auto bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded mb-4"
        >
          Create Room
        </button>
        <form onSubmit={handleJoinRoomSubmit} className="w-full space-y-4">
          <label htmlFor="roomCode" className="block text-gray-700 font-medium">
            Enter room code:
          </label>
          <div className="flex items-center">
            <input
              type="text"
              id="roomCode"
              value={roomCode}
              onChange={handleRoomCodeChange}
              className="border border-gray-300 px-3 py-2 rounded-md w-full text-sm focus:outline-none focus:ring-2 focus:ring-blue-600"
            />
            <button
              type="submit"
              className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded ml-2"
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
