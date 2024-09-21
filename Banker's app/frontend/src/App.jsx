import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'


import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";


import Register from './components/register'
import Login from './components/login'
import Layout from './components/layout';

const router = createBrowserRouter([
  {
    path: "/",
    element: <div><Layout/></div>,
    children:[
      {
        path: "login",
        element: <Login />,
      },
      {
        path: "",
        element: <Register />,
      },
    ]
  },
]);

function App() {
  const [count, setCount] = useState(0)

  return (
    <div className='w-full min-h-[100vh] flex justify-center items-center bg-slate-300'>
      <RouterProvider router={router} >
        <Layout/>
      </RouterProvider>
    </div>
  )
}

export default App
