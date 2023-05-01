import Header from "./Header"
import Footer from "./Footer"

export const Layout = ({ children }) => (
  <div>
    <Header />
    <main className="min-h-screen bg-color overflow-hidden">{children}</main>
    <Footer />
  </div>
)
