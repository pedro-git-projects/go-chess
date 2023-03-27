import '../../styles/drag.css'
import blackKing from "../../assets/black_king.svg"

export const GameBoard = () => (
  // TODO: center piece and fill container
  <div class="board-container">
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


