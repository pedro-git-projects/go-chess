import { useState, useEffect } from "react";

const DarkToggle = () => {
  const btn = "bg-[#CE3262] hover:bg-[#BF2E5A] text-white font-bold py-2 px-4 rounded"
  const [theme, setTheme] = useState("light")
  useEffect(() => {
    theme === "dark" 
      ? document.documentElement.classList.add("dark")
      : document.documentElement.classList.remove("dark")
  }, [theme])
  const handleToggle = () => setTheme(theme === "dark" ? "light" : "dark")
  return (
    <button 
      className={`${btn}`}
      onClick={() => handleToggle()}>
      &#9728;&#65039;
    </button>
  )

}


export default DarkToggle
