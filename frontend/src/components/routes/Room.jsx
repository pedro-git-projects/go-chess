import {useParams} from "react-router-dom"
import {Layout} from "../ui/Layout"

// TODO use clientID to create game with 
// the new multiple games funcitonality
const Room = () => {
  const { clientID } = useParams()
  return (
    <Layout>
      <h1>{clientID}</h1>
    </Layout>
  )
}
export default Room
