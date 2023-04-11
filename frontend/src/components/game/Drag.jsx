import '../../styles/drag.css'
import blackKing from "../../assets/black_king.svg"

export const GameBoard = () => (
  // TODO: center piece and fill container
  <div className="board-container">
    <div id="gameboard"> 
      <div className="square beige">
        <img src={blackKing} alt="Black king piece" id="king" className="chess-piece" draggable/>
      </div>
      <div className="square brown"></div>
      <div className="square beige"></div>
    </div>
    <p id="info">X</p>
  </div>
)

const drag = () => {
  const king = document.querySelector(".chess-piece")
  const getSquares = document.querySelectorAll(".square")
  king.addEventListener('dragStart', dragStart)
}


const dragStart = (e) => {
  console.log(e.target)
}
