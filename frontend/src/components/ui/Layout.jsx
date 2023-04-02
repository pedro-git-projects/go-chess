import Header from "./Header"
import Footer from "./Footer"

export const Layout = ({children}) => (
    <div className="h-screen dark:bg-slate-800">
        <Header/>
        <main>{children}</main>
        <Footer/>
    </div>
)
