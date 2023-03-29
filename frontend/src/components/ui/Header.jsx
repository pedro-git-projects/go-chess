import logo from "../../assets/white_king.svg"

const Header = () => {
  const menuItem = "relative flex h-full items-center p-4 cursor-pointer font-bold text-white hover:bg-white/10 transition-colors ease-in-out"
  const goGradient = "bg-gradient-to-r from-[#00ADD8] to-[#00A29C]"
  const logoText = "font-bold text-2xl text-white drop-shadow-[0_1.2px_1.2px_rgba(0,0,0,0.8)]"

  return (
    <nav className={`flex items-center top-0 ${goGradient}`}>
      <div className="flex items-center p-2 gap-2">
        <img src={logo} alt="white King" width={50}/>
        <div className={`${logoText}`}>
          Go Chess
        </div>
      </div>

      {/* mobile nav */}
      <div className="block md:hidden ml-auto pr-4 my-auto cursor-pointer">
        <div id="">
          <div className="bg-white rounded-full w-8 h-1"></div>
          <div className="bg-white rounded-full w-8 mt-1 h-1"></div>
          <div className="bg-white rounded-full w-8 mt-1 h-1"></div>
        </div>
      </div>

      {/* desktop nav */}
      <div className="hidden md:flex flex-1 items-center justify-end">
        <div className={`${menuItem}`}>Home</div>
        <div className={`${menuItem}`}>Play</div>
        <div className={`${menuItem}`}>Learn</div>
        <div className={`${menuItem}`}>Sign in</div>
      </div>
    </nav>
  ) 
}


export default Header
