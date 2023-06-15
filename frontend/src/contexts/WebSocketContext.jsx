import { createContext, useContext, useEffect, useState } from "react"
import connectToWS from "../hooks/connectToWS"

const WebSocketContext = createContext()

export const useWebSocket = () => useContext(WebSocketContext)

export const WebsocketProvider = ({ children }) => {
  const [ws, setWS] = useState(null)
  useEffect(() => {
    const storedUrl = localStorage.getItem("websocketUrl")
    const connectToWebSocket = async () => {
      const websocket = await connectToWS("ws://localhost:8080/game")
      console.log("connection opened")
      setWS(websocket)
      localStorage.setItem("websocketUrl", websocket.url) // Store WebSocket URL in localStorage
    }

    if (storedUrl) {
      connectToWS(storedUrl)
        .then((websocket) => {
          console.log("reconnected to WebSocket")
          setWS(websocket)
        })
        .catch((error) => {
          console.error("failed to reconnect to WebSocket:", error)
        })
    } else {
      connectToWebSocket()
    }
  }, [])
  return (
    <WebSocketContext.Provider value={ws}>{children}</WebSocketContext.Provider>
  )
}
