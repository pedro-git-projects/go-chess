import { useState } from "react"
import {Link} from "react-router-dom"

import logo from "../../assets/white_king.svg"
const Header = () => {
  const menuItem = "relative flex h-full items-center p-4 cursor-pointer font-bold text-white hover:bg-white/10 transition-colors ease-in-out"
  const mobileItem= "relative flex h-full items-center p-4 cursor-pointer justify-center font-bold text-white hover:bg-white/10 transition-colors ease-in-out"
  const goGradient = "bg-gradient-to-r from-[#00ADD8] to-[#00A29C]"
  const logoText = "font-bold text-2xl text-white drop-shadow-[0_1.2px_1.2px_rgba(0,0,0,0.8)]"
  const line = `h-1 w-8 my-1 rounded-full bg-white transition ease transform duration-300`
  const [isOpen, setIsOpen] = useState(false)
  return (
    <nav className={`flex items-center top-0 ${goGradient}`}>
      <div className="flex items-center p-2 gap-2">
        <Link to={`/`}> <img src={logo} alt="white King" width={50}/> </Link>
        <div className={`${logoText}`}>
          <Link to={`/`}> Go Chess </Link> 
        </div>
      </div>

      {/* mobile nav */}
      {/* button && button animation */}
      <button className="md:hidden flex flex-col ml-auto pr-4 my-auto cursor-pointer"
        onClick={() => setIsOpen(!isOpen)}>
        <div
          className={`${line} ${
                isOpen
                  ? "rotate-45 translate-y-3 opacity-100"
                  : "opacity-100"
            }`}
        />
        <div className={`${line} ${isOpen ? "opacity-0" : "opacity-100"}`} />
        <div
          className={`${line} ${
                isOpen
                  ? "-rotate-45 -translate-y-3 opacity-100"
                  : "opacity-100"
            }`}
        />      
        {/* submenu */}
        <div className={`absolute top-16 ${goGradient} w-full left-0 ${
            isOpen 
              ? "block" 
              : "hidden"
        }`}>
          <div className={`${mobileItem}`}>
            <Link to={`/`}>Home</Link>  
          </div>
          <div className={`${mobileItem}`}>
            <Link to={`/play`}>Play</Link>  
          </div>
          <div className={`${mobileItem}`}>
            <Link to={`/learn`}>Learn</Link>  
          </div>
          <div className={`${mobileItem}`}>
            <Link to={`/signin`}>Sign In</Link>  
          </div>
        </div>
      </button>

      {/* desktop nav */}
      <div className="hidden md:flex flex-1 items-center justify-end">
        <div className={`${menuItem}`}>
          <Link to={`/`}>Home</Link>  
        </div>
        <div className={`${menuItem}`}>
          <Link to={`/play`}>Play</Link>  
        </div>
        <div className={`${menuItem}`}>
          <Link to={`/learn`}>Learn</Link>  
        </div>
        <div className={`${menuItem}`}>
          <Link to={`/signin`}>Sign In</Link>  
        </div>
      </div>
    </nav>
  ) 
}

export default Header
