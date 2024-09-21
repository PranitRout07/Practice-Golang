import {Outlet} from 'react-router-dom';
function Layout(){
    return (
        <div className='w-full min-h-[100vh] flex justify-center items-center bg-slate-300'>
            <Outlet/>
      </div>
    )
}
export default Layout;