import Header from "./Header"
import Footer from "./Footer"

export const Layout = ({children}) => (
    <div className="h-screen bg-color">
        <Header/>
        <main>{children}</main>
        <Footer/>
    </div>
)
