import React, { createContext, useContext } from 'react'
import useLocalStorage from '../hooks/useLocalStorage';
import { jwtDecode } from 'jwt-decode';

const AuthContext = createContext({})
export const AuthProvider = ({children}) => {
    const [token, setToken] = useLocalStorage('token', '')
    
    let decodedJwt = token ? jwtDecode(token) : null;
    const roles = [decodedJwt?.Role]
    const id = decodedJwt?.UserID
    const isLogin = !!token
  return (
    <>
        <AuthContext.Provider value={{token, setToken , roles , isLogin , id}}>
            {children}
        </AuthContext.Provider>
    </>
  )
}
export default AuthContext;
