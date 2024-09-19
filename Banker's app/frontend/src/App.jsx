import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'

function App() {
  const [count, setCount] = useState(0)

  return (
    <div className='w-full min-h-screen'>

    
    <div className='w-full min-h-screen flex justify-center items-center'>
      
        <div className='w-1/4  border-2  border-black h-[30%]'>
          <div className='flex-col space-y-1 text-2xl p-3'>
            <div className='flex space-x-1'>
              <div>
              <label>Firstname</label>
              </div>
            
              <div>
              <input className='bg-slate-100'/>
              </div>
            
            </div>
            <div>
            <label>Lastname</label>
            <input/>
            </div>
           <div>
           <label>Email</label>
           <input/>
           </div>
            
            <div>
            <label>Password</label>
            <input/>
            </div>
          </div>
        </div>
   



    </div>
    </div>
  )
}

export default App
