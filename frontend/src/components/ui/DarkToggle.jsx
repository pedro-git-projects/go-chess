import { useState, useEffect } from "react";

const DarkToggle = () => {
  const btn = "bg-[#CE3262] hover:bg-[#BF2E5A] text-white font-bold py-2 px-4 rounded"

  const [theme, setTheme] = useState("light")
  useEffect(() => {
    theme === "dark" 
      ? document.documentElement.classList.add("dark")
      : document.documentElement.classList.remove("dark")
  }, [theme])
  const handleToggle = () => {
    setTheme(theme === "dark" ? "light" : "dark")
    console.log(theme)
  }

  return (
    <div
      className={`${btn}`}
      onClick={() => handleToggle()}>
      { 
        theme === "dark"
          ? <html>&#127774;</html> 
          : <html>&#127769;</html> 
      } 
    </div>
  )
}


export default DarkToggle
