import { GameBoard } from "../game/Drag"
import { GetBoard } from "../game/GetBoard"
import { Layout } from "../ui/Layout"

const Root = () => ( 
  <Layout>    
    <h1 className="text-3xl font-bold underline">Tailwind</h1>
    <GetBoard/> 
    <GameBoard/>
  </Layout>
)

export default Root
