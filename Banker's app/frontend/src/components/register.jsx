import { useState } from "react";
import { Link } from "react-router-dom";
import axios from 'axios'
function Register(){
    const [resp,setResp] = useState({})
    const [data,setData] = useState({
      first_name:'',
      last_name:'',
      email:'',
      password:'',

    })
    function registerRequest(e){
        e.preventDefault();
        // console.log(data)
        axios.post("/api/account",data).then((resp)=>{
          console.log(resp.data)
          setResp(resp.data)
        })
        

    }
    console.log(resp,"resp")

    

    return(
        <div className='flex-col items-center text-center space-y-4'>
        <h1 className='text-2xl text-teal-600 font-semibold'>Register</h1>
        <div className='bg-white p-16 flex-col space-y-4 shadow-2xl' >
            <div>
              <input className='p-2 text-xl outline-none border-b-2 ' value={data.first_name} onChange={(e)=>{setData((prev)=>{
  return {...prev,first_name:e.target.value}
})}} placeholder='firstname' required type='text'/>
            </div>  
            <div>
              <input className='p-2 text-xl outline-none border-b-2 ' value={data.last_name} onChange={(e)=>{setData((prev)=>{
  return {...prev,last_name:e.target.value}
})}} placeholder='lastname' required type='text'/>
            </div>
            <div>
              <input className='p-2 text-xl outline-none border-b-2 ' value={data.email} onChange={(e)=>{setData((prev)=>{
  return {...prev,email:e.target.value}
})}} placeholder='email' required type='text'/>
            </div>
            <div>
              <input className='p-2 text-xl outline-none border-b-2 ' placeholder='password' value={data.password} onChange={(e)=>{setData((prev)=>{
  return {...prev,password:e.target.value}
})}} required type='text'/>
            </div>
            <div>
              <button onClick={registerRequest} className='p-2 w-full text-center bg-teal-400 hover:scale-105 hover:bg-teal-300'> 
                Register
              </button>
            </div>
           {Object.keys(resp).length !== 0?<div className="text-center text-red-600">
            {console.log(resp)}
             <span>{resp}</span> 
            </div>:""}
            <div className="text-center text-[14px]">
            <span>Do have an account?<Link to="/login">Login</Link></span>
            </div>
        </div>
      </div>
    )
}

export default Register;

