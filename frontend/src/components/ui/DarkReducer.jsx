import {  useReducer } from "react";

const reducer = (state, action) => {
  if (action.type === "toggle") {
    state.isDark === true 
      ? document.documentElement.classList.add("dark")
      : document.documentElement.classList.remove("dark")
    return {
      isDark: !state.isDark
    }
  }
  throw Error('Unknown action.')
}

const DarkReducer = () => {
  const btn = "bg-[#CE3262] hover:bg-[#BF2E5A] text-white font-bold py-2 px-4 rounded cursor-pointer"
  const [theme, dispatch] = useReducer(reducer, { isDark: false }) 
  return (
    <div
      className={`${btn}`}
      onClick={() => dispatch({type: "toggle"})}>
      { 
          document.documentElement.classList.contains("dark")
          ? <html>&#127774;</html> 
          : <html>&#127769;</html> 
      } 
    </div>
  )
}

export default DarkReducer 
