import React from "react"
import ReactDOM from "react-dom/client"
import { createBrowserRouter, RouterProvider } from "react-router-dom"
import { WebsocketProvider } from "./contexts/WebSocketContext"
import { TokenProvider } from "./contexts/AuthContext"
import Err from "./components/routes/Err"
import Learn from "./components/routes/Learn"
import Play from "./components/routes/Play"
import Post from "./components/routes/Post"
import Room from "./components/routes/Room"
import Root from "./components/routes/Root"
import SignIn from "./components/routes/Signin"
import SignUp from "./components/routes/Signup"
import "./styles/layout.css"

const router = createBrowserRouter([
  {
    path: "/",
    element: <Root />,
    errorElement: <Err />,
  },
  {
    path: "/play",
    element: <Play />,
  },
  {
    path: "/learn",
    element: <Learn />,
  },
  {
    path: "/signin",
    element: <SignIn />,
  },
  {
    path: "/signup",
    element: <SignUp />,
  },
  {
    path: "/room/:clientID",
    element: <Room />,
  },
  {
    path: "posts/:id",
    element: <Post />,
  },
])

ReactDOM.createRoot(document.getElementById("root")).render(
  <React.StrictMode>
    <WebsocketProvider>
      <TokenProvider>
        <RouterProvider router={router} />
      </TokenProvider>
    </WebsocketProvider>
  </React.StrictMode>,
)
