import { useState } from "react"

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
        <a href="/nada"> <img src={logo} alt="white King" width={50}/></a>
        <div className={`${logoText}`}>
          <a href="">Go Chess</a>
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
            <a href="">Home</a> 
          </div>
          <div className={`${mobileItem}`}>
            <a href="">Play</a> 
          </div>
          <div className={`${mobileItem}`}>
            <a href="">Learn</a> 
          </div>
          <div className={`${mobileItem}`}>
            <a href="">Sign in</a> 
          </div>
        </div>
      </button>

      {/* desktop nav */}
      <div className="hidden md:flex flex-1 items-center justify-end">
        <div className={`${menuItem}`}>
          <a href="">Home</a> 
        </div>
        <div className={`${menuItem}`}>
          <a href="">Play</a> 
        </div>
        <div className={`${menuItem}`}>
          <a href="">Learn</a> 
        </div>
        <div className={`${menuItem}`}>
          <a href="">Sign in</a> 
        </div>
      </div>
    </nav>
  ) 
}

export default Header
