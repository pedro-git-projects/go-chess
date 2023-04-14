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

const renderSquare = ({clientID}, colIndex, rowIndex, boardState, setBoardState) => {
  const coordinate = `${String.fromCharCode(97 + colIndex)}${8 - rowIndex}`
  const square = boardState.find((square) => square.coordinate === coordinate)
  const isEvenSquare = (colIndex + rowIndex) % 2 === 0
  const backgroundColor = isEvenSquare ? "bg-gray-400" : "bg-white"

  const handleClick = async () => {
    console.log(coordinate)
    const response = await fetch(`http://localhost:8080/calc/${clientID}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ coordinate })
    })

    if (response.ok) {
      const data = await response.json()
      setBoardState(boardState.map((square) => {
        if (data.some((d) => d.coordinate === square.coordinate)) {
          return { ...square, highlighted: true }
        }
        return { ...square, highlighted: false }
      }))
    }
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

const renderRow = ({clientID}, rowIndex, boardState, setBoardState) => (
  <div key={`row${rowIndex}`} className="flex flex-row-reverse">
    { Array.from(Array(8).keys()).map((colIndex) => renderSquare({clientID}, colIndex, rowIndex, boardState, setBoardState)) }
  </div>
)

const renderBoard = ({clientID}, boardState, setBoardState) => (
  <div>
    { Array.from(Array(8).keys()).map((rowIndex) => renderRow({clientID},rowIndex, boardState, setBoardState)) }
  </div>
)

const ChessBoard = ({clientID}) => {
  const [boardState, setBoardState] = useState([])
  useEffect(() => {
    async function fetchBoardState() {
      const response = await fetch(`http://localhost:8080/board/${clientID}`)
      const data = await response.json()
      setBoardState(data)
    }
    fetchBoardState()
  }, [])

  return renderBoard({clientID}, boardState, setBoardState)
}

export default ChessBoard
