import ChessBoard from "../game/ChessBoard"
import { Layout } from "../ui/Layout"

const Root = () => ( 
  <Layout>    
    <h1 className="text-3xl font-bold underline">Tailwind</h1>
    <div className="w-full h-full flex items-center justify-center">
      <div className="mx-auto">
        <ChessBoard/>
      </div>
    </div>
  </Layout>
)

export default Root
