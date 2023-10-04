import React from 'react'

const NavBar = ({ toggleModal, username, setUsername }) => {


    const onClickHandler = async (e) => {
        e.preventDefault()
        try {
            const res = await fetch(`${import.meta.env.VITE_BASE_SERVER_URL}/auth/logout`, {
                "method": "GET",
                "headers": {
                    "Accept": "application/json",
                },
            })
            if (res && res.status !== 200) {
                console.error("Internal server error")
                return
            }
            setUsername("")
        } catch (error) {
            console.error(error)

        }
    }

    return (
        <div className='top-0 absolute bg-[#053B50] w-full h-12 flex items-center px-10 justify-between'>
            <div><h1 className='text-xl cursor-pointer'>CopyShare</h1></div>
            <div className='flex space-x-5'>

                {username && username.length > 0 ?
                    <div className='flex space-x-5'>
                        <p className='text-lg'>{username}</p>

                        <h1 onClick={() => { window.location.href = import.meta.env.VITE_BASE_CLIENT_URL + "/snippets" }} className='text-lg cursor-pointer hover:scale-105 duration-150 ease-linear transition-all'>
                            My Snippets
                        </h1>
                        <button onClick={onClickHandler} className='text-lg cursor-pointer hover:scale-105 duration-150 ease-linear transition-all'>Logout</button>
                    </div> :
                    <h1 onClick={toggleModal} className='text-lg cursor-pointer hover:scale-105 duration-150 ease-linear transition-all'>
                        Login
                    </h1>
                }

            </div>
        </div>
    )
}

export default NavBar