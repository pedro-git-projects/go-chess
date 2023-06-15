import { useContext } from "react"
import { TokenContext } from "../../contexts/AuthContext"
import { Layout } from "../ui/Layout"
import PlayForm from "../ui/PlayForm"
import MustBeSignedIn from "../ui/MustBeSignedIn"

const Play = () => { 
	const { token } = useContext(TokenContext)
	return (
		<Layout>
			{token ? (
				<PlayForm />
			) : (
				<MustBeSignedIn />
			)}
		</Layout>
	)
}

export default Play
