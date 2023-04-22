import { createContext, useContext, useEffect, useState } from "react"
import connectToWS from "../hooks/connectToWS"

const WebSocketContext = createContext()

export const useWebSocket = () => useContext(WebSocketContext) 

export const WebsocketProvider = ({ children }) => {
  const [ws, setWS] = useState(null)
  useEffect(() => {
    const connectToWebSocket = async () =>  {
      const websocket = await connectToWS("ws://localhost:8080/game")
      console.log("connection opened")
      setWS(websocket)
    }
    connectToWebSocket()
  },[])
  return (
    <WebSocketContext.Provider value={ws}>
      {children}
    </WebSocketContext.Provider>
  )
} 

