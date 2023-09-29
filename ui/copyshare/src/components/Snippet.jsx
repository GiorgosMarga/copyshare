import React, { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom';

const Snippet = () => {
    const { id } = useParams();
    const [response, setResponse] = useState({
        title: "",
        content: "",
        createdAt: ""
    })
    useEffect(() => {
        async function fetchData() {
            try {
                // Perform asynchronous operations here (e.g., fetch data)
                const response = await fetch(`${import.meta.env.VITE_BASE_SERVER_URL}/snippet/${id}`, {

                    credentials: 'include'

                });
                if (response.status === 404) {
                    console.log(404)
                    setResponse(prevState => Object.assign({}, prevState, { title: "Error snippet does not exist." }))
                    return
                }
                const data = await response.json();
                setResponse(data)
            } catch (error) {
                // Handle errors
                console.error('Error fetching data:', error);
            }
        }

        // Call the async function
        fetchData();
    }, []);


    return (
        <div className=' w-[70%] h-[70%] rounded-xl shadow-lg flex bg-[#053B50] flex-col'>
            <p className='text-white text-xl font-semibold self-center p-5'>{`Snippet: ${id}`}</p>
            <div className='flex flex-col py-2 px-10  h-full space-y-5 rounded-xl bg-[#053B50]'>
                <p className=" bg-[#176B87]/10 px-5 py-[5px] rounded-xl" >{response.title}</p>
                <p className=' bg-[#176B87]/10 px-5 p-2 min-w-[70%] min-h-[70%]' >{response.content}</p>
                <p className='text-sm font-bold' >{`Created at: ${response.createdAt}`}</p>
            </div>
        </div >
    )
}

export default Snippet