import { createContext, useState } from "react"

export const TokenContext = createContext(null)

export const TokenProvider = ({ children }) => {
  const [token, setToken] = useState(null)

  const resetToken = () => {
    setToken(null)
  }

  return (
    <TokenContext.Provider value={{ token, setToken, resetToken }}>
      {children}
    </TokenContext.Provider>
  )
}
