import Header from "./Header"
import Footer from "./Footer"

export const Layout = ({children}) => (
<div className="min-h-screen bg-color overflow-hidden">
        <Header/>
        <main>{children}</main>
        <Footer/>
    </div>
)
