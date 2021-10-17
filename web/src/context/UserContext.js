import React, { createContext, useEffect, useState } from 'react';
import { getUserData } from '../adapters/global/user';

export const UserContext = createContext();

export const UserProvider = ({ children }) => {
    // Username of the currently logged in user
    const [username, setUsername] = useState('');
    // Permissions level of the user (permissions integer)
    const [permissions, setPermissions] = useState(0);
    // Whether or not the user is logged in
    const [loggedIn, setLoggedIn] = useState(false);
    // Whether or not the user data has already been retrieved from the API
    const [pending, setPending] = useState(true);

    // Check if user is logged in. And if so, get their user data from the API
    // and store it in the global user context.
    useEffect(() => {
        // Get login status and if logged in also data of the currently logged-in user
        getUserData()
            .then(userData => {
                setUsername(userData.username);
                setPermissions(userData.role);
                setLoggedIn(true);
                setPending(false);
            })
            .catch(httpErrorCode => {
                setPending(false);
                setLoggedIn(false);
            });
    // eslint-disable-next-line
    }, []);

    return (
        <UserContext.Provider
            value={{
                username,
                setUsername,
                permissions,
                setPermissions,
                loggedIn,
                setLoggedIn,
                pending,
                setPending
            }}
        >
            { children }
        </UserContext.Provider>
    );
}