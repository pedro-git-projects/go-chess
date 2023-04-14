import {useParams} from "react-router-dom"
import ChessBoard from "../game/ChessBoard"
import {Layout} from "../ui/Layout"

const Room = () => {
  const { clientID } = useParams()
  return (
    <Layout>
      <div className="w-full h-full flex align-middle items-center justify-center">
        <div className="mx-auto">
          <ChessBoard clientID={clientID}></ChessBoard>
        </div>
      </div>
    </Layout>
  )
}
export default Room
