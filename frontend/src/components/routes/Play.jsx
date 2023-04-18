import { Layout } from "../ui/Layout"
import NewRoom from "../dynamic/NewRoom"
import JoinRoom from "../dynamic/JoinRoom"

const Play = () => ( 
  <Layout>    
    <h1 className="text-3xl font-bold underline">Start a new chess game</h1>
    <NewRoom/>
    <JoinRoom/>
  </Layout>
)

export default Play 
