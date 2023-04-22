import React, { useState, useEffect } from "react"
import whitePawn from "../../assets/white_pawn.svg"
import blackPawn from "../../assets/black_pawn.svg"
import whiteKnight from "../../assets/white_horse.svg"
import blackKnight from "../../assets/black_horse.svg"
import whiteBishop from "../../assets/white_bishop.svg"
import blackBishop from "../../assets/black_bishop.svg"
import whiteRook from "../../assets/white_rook.svg"
import blackRook from "../../assets/black_rook.svg"
import whiteQueen from "../../assets/white_queen.svg"
import blackQueen from "../../assets/black_queen.svg"
import whiteKing from "../../assets/white_king.svg"
import blackKing from "../../assets/black_king.svg"
import sendMessage from "../../hooks/sendMessage"
import { useWebSocket } from "../../contexts/WebSocketContext"

const getPieceSymbol = (piece) => {
  const [color, type] = piece.split(" ")
  const svg = {
    "white-pawn": whitePawn,
    "black-pawn": blackPawn,
    "white-rook": whiteRook,
    "black-rook": blackRook,
    "white-knight": whiteKnight,
    "black-knight": blackKnight,
    "white-bishop": whiteBishop,
    "black-bishop": blackBishop,
    "white-queen": whiteQueen,
    "black-queen": blackQueen,
    "white-king": whiteKing,
    "black-king": blackKing,
  }[`${color}-${type}`]
  return <img src={svg} alt={type}/>
}

const renderSquare = ({roomID}, colIndex, rowIndex, boardState, setBoardState) => {
  const coordinate = `${String.fromCharCode(97 + colIndex)}${8 - rowIndex}`
  const square =  boardState.find((square) => square.coordinate === coordinate)
  const isEvenSquare = (colIndex + rowIndex) % 2 === 0
  const backgroundColor = isEvenSquare ? "bg-gray-400" : "bg-white"
  const ws = useWebSocket()
  const handleClick = async () => {
    console.log(coordinate)
    const msg = JSON.stringify({message:"calc",coordinate: coordinate, room_id:roomID})
    const resp = await sendMessage(ws, msg)
    console.log("response received:", resp)      
    setBoardState(boardState.map((square) => {
      if (JSON.parse(resp.legal_movements).some((d) => d.coordinate === square.coordinate)) {          
        return { ...square, highlighted: true }
      }
      return { ...square, highlighted: false }
    }))
  }

  return (
    <div
      key={`${colIndex}${rowIndex}`}
      className={`w-16 h-16 flex items-center justify-center ${backgroundColor} ${square?.highlighted ? 'bg-yellow-500' : ''}`}
      onClick={handleClick}
    >
      {square && square.piece !== "empty" && getPieceSymbol(square.piece)}
    </div>
  )
}

const renderRow = ({roomID}, rowIndex, boardState, setBoardState) => (
  <div key={`row${rowIndex}`} className="flex flex-row-reverse">
    { Array.from(Array(8).keys()).map((colIndex) => renderSquare({roomID}, colIndex, rowIndex, boardState, setBoardState)) }
  </div>
)

const renderBoard = ({roomID}, boardState, setBoardState) => (
  <div>
    { Array.from(Array(8).keys()).map((rowIndex) => renderRow({roomID}, rowIndex, boardState, setBoardState)) }
  </div>
)

const ChessBoard = ({roomID, clientID}) => {
  const [boardState, setBoardState] = useState([])
  const ws = useWebSocket()
  useEffect(() => {
    const fetchBoardState = async () => {
      const msg = JSON.stringify({message:"render", room_id:roomID, clientID:clientID})
      const resp = await sendMessage(ws, msg)
      console.log("response received:", resp)
      setBoardState(JSON.parse(resp.state))
    }
    fetchBoardState()
  }, [])

  // Listen to WebSocket messages and update boardState
  useEffect(() => {
    const handleMessage = (event) => {
      const data = JSON.parse(event.data)
      if (data.legal_movements) {
        const legalMovements = JSON.parse(data.legal_movements)
        const newBoardState = boardState.map((square) => {
          if (legalMovements.some((d) => d.coordinate === square.coordinate)) {
            return { ...square, highlighted: true }
          } else {
            return { ...square, highlighted: false }
          }
        })
        setBoardState(newBoardState)
      }
    }
    ws.addEventListener('message', handleMessage)
    return () => {
      ws.removeEventListener('message', handleMessage)
    }
  }, [boardState, ws])
  return renderBoard({roomID}, boardState, setBoardState)
}

export default ChessBoard
