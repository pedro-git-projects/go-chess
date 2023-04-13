import {  useReducer } from "react"

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
  const btn = "text-white font-bold py-2 px-4 rounded cursor-pointer"
  const [theme, dispatch] = useReducer(reducer, { isDark: false }) 
  return (
    <div
      className={`${btn}`}
      onClick={() => dispatch({type: "toggle"})}>
      { 
        document.documentElement.classList.contains("dark")
          ?  "Go back" 
          :  "Go Dark" 
      } 
    </div>
  )
}

export default DarkReducer 
