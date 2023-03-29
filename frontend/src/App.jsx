import { GameBoard } from "./components/game/Drag"
import { GetBoard } from "./components/game/GetBoard"
import {Layout} from "./components/ui/Layout"

const App = () =>  (
  <Layout>    
    <h1 className="text-3xl font-bold underline">Tailwind</h1>
    <GetBoard/> 
    <GameBoard/>
  </Layout>
)

export default App
