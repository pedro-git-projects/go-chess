import { useState } from "react";

const ChessBoardRequest = () => {
  const [pieces, setPieces] = useState(Array(64).fill(null));

  const handleSquareClick = (index) => {
    const piece = pieces[index];
    if (piece) {
      const rank = piece.split(':')[1].split(' ')[1];
      const file = String.fromCharCode(97 + (index % 8));
      const rankIndex = 8 - Math.floor(index / 8);
      const query = `?rank=${rank}&file=${file}&rankIndex=${rankIndex}`;
      fetch(`https://example.com/api/piece${query}`)
        .then((response) => response.json())
        .then((data) => {
          const { newIndex } = data;
          const newPieces = [...pieces];
          newPieces[index] = null;
          newPieces[newIndex] = piece;
          setPieces(newPieces);
        })
        .catch((error) => console.error(error));
    }
  };

  const getPieceColor = (piece) => {
    const color = piece.includes(':') ? piece.split(':')[1].split(' ')[0] : '';
    return color === 'white' ? 'text-white' : 'text-black';
  };

  const getPieceSymbol = (piece) => {
    const type = piece.includes(':') ? piece.split(':')[1].split(' ')[1] : '';
    switch (type) {
      case 'pawn':
        return '♙';
      case 'rook':
        return '♖';
      case 'knight':
        return '♘';
      case 'bishop':
        return '♗';
      case 'queen':
        return '♕';
      case 'king':
        return '♔';
      default:
        return '';
    }
  };

  const renderSquare = (i) => {
    const piece = pieces[i];
    const colorClass = (i + Math.floor(i / 8)) % 2 === 0 ? 'bg-gray-300' : 'bg-gray-500';
    const pieceColorClass = getPieceColor(piece);
    const pieceSymbol = getPieceSymbol(piece);
    return (
      <div
        key={i}
        className={`w-12 h-12 flex justify-center items-center ${colorClass} ${pieceColorClass}`}
        onClick={() => handleSquareClick(i)}
      >
        {pieceSymbol}
      </div>
    );
  };

  const renderRow = (start) => {
    const row = [];
    for (let i = start; i < start + 8; i++) {
      row.push(renderSquare(i));
    }
    return row;
  };

  const renderBoard = () => {
    const board = [];
    for (let i = 0; i < 8; i++) {
      board.push(<div key={i} className="flex">{renderRow(i * 8)}</div>);
    }
    return board;
  };

  return (
    <div className="w-max">
      {renderBoard()}
    </div>
  );
};

export default ChessBoardRequest;
