/*
    Adapter for handling user-specific requests to the API.
*/

/**
 * @typedef UserData
 * @property {number} role Role of the current user (permissions integer) https://github.com/fabiancdng/GoShortrr/wiki/%F0%9F%94%8C-REST-API#permissions
 * @property {string} username Username of the current user
 */

/**
 * Requests user data of the user that is currently logged in
 * from the API.
 * 
 * `[POST] /api/auth/user`
 * 
 * Reference: https://github.com/fabiancdng/GoShortrr/wiki/%F0%9F%94%8C-REST-API#post-apiauthuser
 * 
 * @returns {Promise<UserData|number>} User data of the current user or HTTP error code
 */
export const getUserData = async () => {
    return new Promise(async (resolve, reject) => {
        var res = await fetch('/api/auth/user', {
            method: 'POST',
            credentials: 'include'
        });
        
        if(!res.ok) {
            reject(res.status);
        } else {
            res = await res.json();
            resolve(res);
        }
    });
}

/**
 * Sends the passed credentials to the login endpoint of the
 * API to (in case the credentials entered are correct) exchange
 * them against a session cookie.
 * 
 * `[POST] /api/auth/login`
 * 
 * Reference: https://github.com/fabiancdng/GoShortrr/wiki/%F0%9F%94%8C-REST-API#post-apiauthlogin
 *
 * @param {string} username The entered username
 * @param {string} password The entered password
 * 
 * @returns {Promise<boolean|number>} `true` if the login succeeded or otherwise the HTTP error code
 */
export const loginUser = (username, password) => {
    return new Promise(async (resolve, reject) => {
        var res = await fetch('/api/auth/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ username: username, password: password })
        });

        if (!res.ok) {
            // Credentials are incorrect or an error occurred
            reject(res.status);
        } else {
            // Credentials are correct and the login succeeded
            resolve(true);
        }
    });
}

/**
 * Sends a logout request to the API that destroys the
 * session (& session cookie) and therefore logs out the
 * user.
 * 
 * TODO: Corresponding API documentation in the GitHub Wiki
 * TODO: Better implementation by checking for errors (as the API could have network failures, etc.)
 * 
 * @returns {Promise<boolean>} Always `true`
 */
export const logoutUser = () => {
    return new Promise(async (resolve, reject) => {
        await fetch('/api/auth/logout', {
            method: 'POST',
            credentials: 'include'
        });

        resolve(true);
    });
}