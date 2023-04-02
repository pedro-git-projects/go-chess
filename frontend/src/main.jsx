import React from 'react'
import ReactDOM from 'react-dom/client'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import Err from './components/routes/Err'
import Learn from './components/routes/Learn'
import Play from './components/routes/Play'
import Root from './components/routes/Root'
import SignIn from './components/routes/Signin'
import SignUp from './components/routes/Signup'
import "./styles/layout.css"

const router = createBrowserRouter([
  {
    path: "/",
    element: <Root/>, 
    errorElement: <Err/>,
  },
  {
    path: "/play",
    element: <Play/>, 
  },
  {
    path: "/learn",
    element: <Learn/>, 
  },
  {
    path: "/signin",
    element: <SignIn/>, 
  },
  {
    path: "/signup",
    element: <SignUp/>, 
  },
])

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <RouterProvider router={router}/>
  </React.StrictMode>,
)
