import React, { useEffect, useState } from 'react'
import SnippetQuickView from './SnippetQuickView'

const SnippetList = () => {
    const [response, setResponse] = useState([])
    useEffect(() => {
        async function fetchData() {
            try {
                // Perform asynchronous operations here (e.g., fetch data)
                const response = await fetch(`${import.meta.env.VITE_BASE_SERVER_URL}/snippet`, {
                    credentials: 'include'
                });
                const data = await response.json();
                setResponse(data)
            } catch (error) {
                // Handle errors
                console.error('Error fetching data:', error);
            }
        }
        fetchData();
    }, []);
    return (
        <div className=' w-[70%] h-[70%] rounded-xl pb-2 px-5 md:px-10 shadow-lg flex bg-[#053B50] flex-col overflow-scroll space-y-3 scrollbar-hide'>
            <p className='text-white text-xl font-semibold self-center p-5'>Snippets</p>
            {response.map((snippet) => <SnippetQuickView title={snippet.title} id={snippet.id} createdAt={snippet.createdAt} />)}
        </div >
    )
}

export default SnippetList