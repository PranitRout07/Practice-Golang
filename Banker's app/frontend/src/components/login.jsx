import { useState } from "react";
import { Link } from "react-router-dom";
import axios from 'axios'

function Login(){
  const [resp,setResp] = useState({})
  const [data,setData] = useState({
    email:'',
    password:''
  })


  function loginRequest(e){
    e.preventDefault();
    // console.log(data)
    axios.post("/api/login",data).then((response)=>{
      console.log(response.data)
      setResp(response)
    })

    console.log(resp)
    

}
    return(
        <div className='flex-col items-center text-center space-y-4'>
        <h1 className='text-2xl text-teal-600 font-semibold'>Login</h1>
        <div className='bg-white p-16 flex-col space-y-4 shadow-2xl' >
            <div>
              <input className='p-2 text-xl outline-none border-b-2 ' placeholder='email' required type='text' value={data.email} onChange={(e)=>{setData((prev)=>{
  return {...prev,email:e.target.value}
})}} />
            </div>
            <div>
              <input className='p-2 text-xl outline-none border-b-2 ' placeholder='password' required type='text' value={data.password} onChange={(e)=>{setData((prev)=>{
  return {...prev,password:e.target.value}
})}}/>
            </div>
            <div>
              <button onClick={loginRequest} className='p-2 w-full text-center bg-teal-400 hover:scale-105 hover:bg-teal-300'> 
                Login
              </button>
            </div>
            <div className="text-center text-red-600">
            <span>Error</span>
            </div>
            <div className="text-center text-[14px]">
            <span>Don't have an account?<Link to="/">Register</Link></span>
            </div>
        </div>
      </div>
    )
}

export default Login;
