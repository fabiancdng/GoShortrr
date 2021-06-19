import { createContext, useState } from "react";

export const UserContext = createContext()

export const UserProvider = ({children}) => {
    const [username, setUsername] = useState('')
    const [permissions, setPermissions] = useState(0)
    const [loggedIn, setLoggedIn] = useState(false)

    return (
        <UserContext.Provider
            value={{
                username,
                setUsername,
                permissions,
                setPermissions,
                loggedIn,
                setLoggedIn
            }}
        >
            {children}
        </UserContext.Provider>
    )
}