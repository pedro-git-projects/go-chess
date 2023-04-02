import React from 'react'
import ReactDOM from 'react-dom/client'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import Err from './components/routes/Err'
import Root from './components/routes/Root'
import "./styles/layout.css"

const router = createBrowserRouter([
  {
    path: "/",
    element: <Root/>, 
    errorElement: <Err/>,
  },
])

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <RouterProvider router={router}/>
  </React.StrictMode>,
)
