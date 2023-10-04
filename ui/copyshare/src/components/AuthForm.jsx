import React, { useState } from 'react'
import Modal from 'react-modal';
import BounceLoader from "react-spinners/BounceLoader";
import FormInput from './FormInput';
const AuthForm = ({ isOpen, setUser, toggleModal }) => {
    const [isLogin, setIsLogin] = useState(true)
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")
    const [username, setUsername] = useState("")
    const [confirmPassword, setConfirmPassword] = useState("")
    const [loading, setLoading] = useState(false);
    const [errors, setErrors] = useState({
        password: "",
        email: "",
        confirmPassword: "",
        username: "",
        general: ""
    })

    const resetErrors = () => {
        setErrors({
            password: "",
            email: "",
            confirmPassword: "",
            username: "",
            general: ""
        })
    }

    const resetForm = () => {

        setEmail("")
        setPassword("")
        setConfirmPassword("")
        setUsername("")
    }
    const onChangeEmailHandler = (e) => {
        setEmail(e.target.value)
    }

    const onChangePasswordHandler = (e) => {
        setPassword(e.target.value)
    }

    const onChangeConfirmPasswordHandler = (e) => {
        setConfirmPassword(e.target.value)
    }
    const onChangeUsernameHandler = (e) => {
        setUsername(e.target.value)
    }

    const toggleLogin = () => {
        resetErrors()
        setIsLogin(prevState => !prevState)
    }
    const onClickHandler = async (e) => {
        e.preventDefault()
        setLoading(true)
        resetErrors()
        let res
        if (isLogin) {
            try {
                res = await fetch(`${import.meta.env.VITE_BASE_SERVER_URL}/auth/login`, {
                    "method": "POST",
                    "headers": {
                        "Accept": "application/json",
                        "Content-Type": "application/json"
                    },
                    body: JSON.stringify({ "email": email, "password": password })
                })
                const bd = await res.json()
                if (res && res.status !== 200) {
                    setErrors(bd)
                    setLoading(false)
                    return
                }
                setUser(bd.username)
                toggleModal()
                resetForm()
            } catch (err) {
                setLoading(false)
                setErrors({
                    general: "Internal server error. Please try again later."
                })
            }

        } else {
            if (confirmPassword !== password) {
                setErrors(prevState => Object.assign({}, prevState, { confirmPassword: "passwords don't match" }))
                setLoading(false)
                return
            }
            try {
                res = await fetch(`${import.meta.env.VITE_BASE_SERVER_URL}/auth/register`, {
                    "method": "POST",
                    "headers": {
                        "Accept": "application/json",
                        "Content-Type": "application/json"
                    },
                    body: JSON.stringify({ "email": email, "password": password, "username": username })
                })
                const bd = await res.json()
                if (res && res.status !== 201) {
                    setErrors(bd)
                    setLoading(false)
                    return
                }
                console.log(bd)
                setUsername(bd.username)
                toggleModal()
                resetForm()
            } catch (error) {
                setErrors({
                    general: "Internal server error. Please try again later."
                })
                setLoading(false)
            }

        }
        setLoading(false)

    }
    return (

        <Modal
            shouldCloseOnOverlayClick={true}
            onRequestClose={toggleModal}
            isOpen={isOpen}
            contentLabel="Example Modal"
            className="relative text-[#EEEEEE] bg-[#176B87] flex flex-col justify-center items-center h-[70%] w-[70%] rounded-xl shadow-xl p-10 outline-none"
            overlayClassName="bg-gray-500/50 absolute top-0 w-screen h-screen flex justify-center items-center"
        >

            <h1 className='absolute top-20 text-xl flex font-bold'>Login to your <h1 className='px-2 font-extrabold text-2xl mb-2  animate-bounce duration-400 transition-all ease-linear '> copyshare </h1> account.</h1>
            <form className='flex  flex-col space-y-2 w-[50%] pt-10'>
                {!isLogin ? <FormInput placeholder="Username" value={username} error={errors.username} onChangeHandler={onChangeUsernameHandler} /> : null}

                <FormInput placeholder={"Email"} value={email} error={errors.email} onChangeHandler={onChangeEmailHandler} />
                <FormInput placeholder={"Password"} value={password} error={errors.password} type={"password"} onChangeHandler={onChangePasswordHandler} />
                {!isLogin ? <FormInput placeholder="Confirm password" value={confirmPassword} error={errors.confirmPassword} type={"password"} onChangeHandler={onChangeConfirmPasswordHandler} /> : null}


                <p onClick={toggleLogin} className='text-xs font-thin pt-1 cursor-pointer pl-2'>{isLogin ? "Don't have an account? Register" : "Already have an account? Login"}</p>
                {errors.general.length != 0 ? <p className='pl-1 text-red-700 text-xs font-bold'>{errors.general}</p> : null}

                {loading === false ? <div className='self-center pt-5 '>
                    <button onClick={onClickHandler} className='py-1 px-7 bg-[#053B50]/70 rounded-xl  text-[#EEEEEE] '>{isLogin ? "Login" : "Register"}</button>
                </div> : <div className=' flex justify-center'>
                    <BounceLoader
                        color={"#053B50"}
                        loading={loading}
                        size={28}
                        aria-label="Loading Spinner"
                        data-testid="loader"

                    /></div>}
            </form>
        </Modal >

    )
}

export default AuthForm