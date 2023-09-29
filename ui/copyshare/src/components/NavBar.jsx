import React from 'react'

const NavBar = ({ toggleModal }) => {
    return (
        <div className='top-0 absolute bg-[#053B50] w-full h-12 flex items-center px-10 justify-between'>
            <div><h1 className='text-xl cursor-pointer'>CopyShare</h1></div>
            <div className='flex space-x-5'>
                <h1 onClick={() => { window.location.href = import.meta.env.VITE_BASE_CLIENT_URL + "/snippets" }} className='text-lg cursor-pointer hover:scale-105 duration-150 ease-linear transition-all'>
                    My Snippets
                </h1>
                <h1 onClick={toggleModal} className='text-lg cursor-pointer hover:scale-105 duration-150 ease-linear transition-all'>
                    Login
                </h1>
            </div>
        </div>
    )
}

export default NavBar