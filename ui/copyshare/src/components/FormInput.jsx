import React from 'react'

const FormInput = ({ error, className, value, placeholder, onChangeHandler, type }) => {

    return (
        <div className='flex flex-col  space-y-[3px]'>
            {error ? <p className='pl-1 text-red-700 text-xs font-bold'>{error}</p> : null}
            <input className={` ${error && error.length !== 0 ? "border-red-700 border-2" : ""} outline-none bg-[#053B50] px-5 py-[5px] rounded-xl  placeholder:text-gray-400/50 ${className}`} onChange={onChangeHandler} value={value} type={type} placeholder={placeholder} />
        </div>
    )
}

export default FormInput