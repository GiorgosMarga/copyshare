import { useEffect, useState } from 'react'
import NavBar from './components/NavBar'
import AuthForm from './components/AuthForm';
import CreateSnippet from './components/CreateSnippet';

import { Routes, Route } from "react-router-dom";
import Snippet from './components/Snippet';
import SnippetList from './components/SnippetList';
// Make sure to bind modal to your appElement (https://reactcommunity.org/react-modal/accessibility/)


function App() {
  const [modalIsOpen, setIsOpen] = useState(false);
  const [username, setUser] = useState("");

  useEffect(() => {
    async function whoAmI() {
      try {
        // Perform asynchronous operations here (e.g., fetch data)
        const response = await fetch(`${import.meta.env.VITE_BASE_SERVER_URL}/auth/whoAmI`, {
          credentials: 'include'
        });
        const data = await response.json();
        setUser(data.username)
      } catch (error) {
        // Handle errors
        console.error('Error fetching data:', error);
      }
    }
    whoAmI();
  }, [])


  const toggleModal = () => {
    setIsOpen(prevState => !prevState);
  }
  return <div className='relative bg-[#176B87] w-screen h-screen flex flex-col items-center justify-center text-[#EEEEEE]'>
    <NavBar toggleModal={toggleModal} username={username} />
    <AuthForm isOpen={modalIsOpen} toggleModal={toggleModal} setUser={setUser} />
    <Routes>
      <Route element={<CreateSnippet />} exact path="" />
      <Route path='/:id' exact element={<Snippet />} />
      <Route path='/snippets' exact element={<SnippetList />} />
    </Routes>
  </div>
}

export default App
