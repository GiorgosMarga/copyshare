import React from 'react'

const SnippetQuickView = ({ title, createdAt, id }) => {

    const onClickHandler = () => {
        window.location.href = import.meta.env.VITE_BASE_CLIENT_URL + "/" + id
    }

    return (
        <div onClick={onClickHandler} className='cursor-pointer hover:scale-105 transition-all duration-150 ease-in-out flex py-5 px-5 h-fit  space-y-5 rounded-xl bg-[#176B87]/10'>
            <div className='flex flex-col space-y-1'>
                <p className="text-lg rounded-xl" >{title}</p>
                <p className='text-xs font-semibold'>{createdAt}</p>
            </div>
        </div>
    )
}

export default SnippetQuickView