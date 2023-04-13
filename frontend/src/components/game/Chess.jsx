import React, { useState } from 'react';

const ChessBoard = ()  => {
  const [squares, setSquares] = useState([
    'a8:black rook',
    'b8:black knight',
    'c8:black bishop',
    'd8:black queen',
    'e8:black king',
    'f8:black bishop',
    'g8:black knight',
    'h8:black rook',
    'a7:black pawn',
    'b7:black pawn',
    'c7:black pawn',
    'd7:black pawn',
    'e7:black pawn',
    'f7:black pawn',
    'g7:black pawn',
    'h7:black pawn',
    '', '', '', '', '', '', '', '',
    '', '', '', '', '', '', '', '',
    '', '', '', '', '', '', '', '',
    '', '', '', '', '', '', '', '',
    'a2:white pawn',
    'b2:white pawn',
    'c2:white pawn',
    'd2:white pawn',
    'e2:white pawn',
    'f2:white pawn',
    'g2:white pawn',
    'h2:white pawn',
    'a1:white rook',
    'b1:white knight',
    'c1:white bishop',
    'd1:white queen',
    'e1:white king',
    'f1:white bishop',
    'g1:white knight',
    'h1:white rook',
  ]);

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
  const renderSquare = (file, rank, index) => {
    const squareColor = (file.charCodeAt(0) + rank) % 2 === 0 ? 'bg-gray-500' : 'bg-white';
    const piece = squares[index] || '';
    const pieceColor = getPieceColor(piece);
    const pieceSymbol = getPieceSymbol(piece);

    return (
      <div key={file + rank} className={`w-12 h-12 flex items-center justify-center ${squareColor} ${pieceColor}`}>
        {pieceSymbol}
      </div>
    );
  };


  const renderRank = (rank) => {
    const files = ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'];
    return (
      <div key={`rank-${rank}`} className="flex">
        {files.map((file) => renderSquare(file, rank, (8 - rank) * 8 + files.indexOf(file)))}
      </div>
    );
  };

  const ranks = [8, 7, 6, 5, 4, 3, 2, 1];
  return (
    <div className="flex flex-col items-center justify-center min-h-screen">
      <div className="border border-gray-900 rounded-md overflow-hidden">
        {ranks.map((rank) => renderRank(rank))}
      </div>
    </div>
  );
}

export default ChessBoard
