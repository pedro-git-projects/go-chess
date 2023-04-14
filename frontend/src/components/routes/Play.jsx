import { Layout } from "../ui/Layout"
import NewRoom from "../utils/NewRoom"

const Play = () => ( 
  <Layout>    
    <h1 className="text-3xl font-bold underline">Start a new chess game</h1>
    <NewRoom/>
  </Layout>
)

export default Play 
