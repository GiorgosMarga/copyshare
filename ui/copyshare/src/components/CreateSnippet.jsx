import React, { useState } from 'react'
import FormInput from './FormInput'
import BounceLoader from 'react-spinners/BounceLoader'
const CreateSnippet = () => {
    const [title, setTitle] = useState("")
    const [content, setContent] = useState("")
    const [loading, setLoading] = useState(false)
    const [errors, setErrors] = useState({
        title: "",
        content: ""
    })

    const resetForm = () => {
        setTitle("")
        setContent("")
    }

    const resetErrors = () => {
        setErrors({
            title: "",
            content: ""
        })
    }

    const onChangeTitleHandler = (e) => {
        setTitle(e.target.value)
    }
    const onChangeContentHandler = (e) => {
        setContent(e.target.value)

    }

    const onSubmitHandler = async (e) => {

        e.preventDefault()
        resetErrors()

        setLoading(true)
        const res = await fetch(`${import.meta.env.VITE_BASE_SERVER_URL}/snippet`, {
            method: "POST",
            credentials: 'include',
            headers: {
                "Accept": "application/json",
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ title, content })
        })
        const body = await res.json()
        setLoading(false)

        if (res && res.status !== 201) {
            console.log(body)
            setErrors(body)
            return
        }
        resetForm()
        window.location.href = import.meta.env.VITE_BASE_CLIENT_URL + "/" + body;
    }

    return (
        <div className=' w-[70%] h-[70%] rounded-xl shadow-lg flex bg-[#053B50] flex-col'>
            <p className='text-white text-xl font-semibold self-center p-5'>Create Snippet</p>
            <form className='flex flex-col py-2 px-10  h-full space-y-5 rounded-xl bg-[#053B50]'>
                <FormInput className="bg-[#176B87]/10" value={title} onChangeHandler={onChangeTitleHandler} placeholder={"Title"} error={errors.title} />
                <textarea className=' bg-[#176B87]/10 outline-none placeholder:text-gray-400/50 px-5 p-2 min-w-[70%] min-h-[70%]' placeholder='Content' value={content} onChange={onChangeContentHandler} />
                {loading === false ? <button onClick={onSubmitHandler} className='cursor-pointer hover:scale-105  w-fit  self-center transition-all duration-100 ease-linear'>Create</button> : <div className=' flex justify-center'>
                    <BounceLoader
                        color={"#176B87"}
                        loading={loading}
                        size={28}
                        aria-label="Loading Spinner"
                        data-testid="loader"

                    /></div>}

            </form>
        </div >
    )
}

export default CreateSnippet