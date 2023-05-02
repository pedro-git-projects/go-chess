import { useState, useRef } from "react"

const DarkRef = () => {
  const btn =
    "bg-[#CE3262] hover:bg-[#BF2E5A] text-white font-bold py-2 px-4 rounded"
  const [isDark, setIsDark] = useState(false)
  const ref = useRef(null)

  const handleClick = () => {
    const nextDark = !isDark
    setIsDark(nextDark)
    if (nextDark) {
      ref.current = true
      console.log(ref.current)
      document.documentElement.classList.add("dark")
    } else {
      console.log(ref.current)
      ref.current = false
      document.documentElement.classList.remove("dark")
    }
  }
  return (
    <div className={`${btn}`} onClick={() => handleClick()}>
      {isDark ? "&#127774;" : "&#127769;"}
    </div>
  )
}

export default DarkRef
