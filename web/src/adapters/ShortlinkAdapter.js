/*
    Adapter for handling shortlink-related requests to the API.
*/

/**
 * @typedef ShortlinkData
 * @property {number} id Internal ID of the shortlink
 * @property {string} link The original (long) link behind the shortlink
 * @property {string} short The unique part of the shortlink
 * @property {user} user The internal ID of the user who created the shortlink
 * @property {boolean} password Whether the shortlink is protected by a password
 * @property {string} created Timestamp (datetime in ISO format) of shortlink creation
 */

/**
 * Gets the data behind a shortlink.
 * 
 * `[GET] /api/shortlink/get/SHORT`
 * 
 * Reference: https://github.com/fabiancdng/GoShortrr/wiki/%F0%9F%94%8C-REST-API#get-apishortlinkgetshort
 * 
 * @param {string} link The shortlink to look up
 * 
 * @returns {Promise<ShortlinkData|number>} The shortlink's data or the HTTP error code
 */
export const getShortlinkData = (link) => {
    return new Promise(async (resolve, reject) => {
        // Remove possible root URL / path from the URL
        link = link.replace(window.location.href, '');
        var location = window.location.href.replace('http://', '').replace('https://', '');
        link = link.replace(location, '');
        
        var res = await fetch(`/api/shortlink/get/${link}`);
        
        if (!res.ok) {
            // Reject promise with the HTTP error code as value
            reject(res.status);
        } else {
            // Resolve promise with shortlink data as value
            resolve(await res.json());
        }
    });
}

/**
 * Shortens a URL by sending a shortlink create request to the API.
 * 
 * `[POST] /api/shortlink/create`
 * 
 * Reference: https://github.com/fabiancdng/GoShortrr/wiki/%F0%9F%94%8C-REST-API#post-apishortlinkcreate
 * 
 * @param {string} link The long URL that a shortlink should be generated for
 * 
 * @returns {Promise<string|number>} The generated shortlink (with root URL) or the HTTP error code
 */
export const createShortlink = (link) => {
    return new Promise(async (resolve, reject) => {
        var res = await fetch('/api/shortlink/create', {
            method: 'POST',
            credentials: 'include',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({link: link})
        });

        if (!res.ok) {
            // Reject promise with the HTTP error code as value
            reject(res.status);
        } else {
            // Grab unique part of the shortlink, append the root URL, and
            // resolve the promise with it as value
            res = await res.json();
            let short = window.location.href + res.short;
            resolve(short);
        }
    });
}

/**
 * Deletes a shortlink by sending a shortlink delete request to the API.
 * 
 * `[DELETE] /api/shortlink/delete/SHORT`
 * 
 * Reference: https://github.com/fabiancdng/GoShortrr/wiki/%F0%9F%94%8C-REST-API#delete-apishortlinkdeleteshort
 * 
 * @param {string} link The shortlink that should be deleted (no matter if with or without root URL)
 * 
 * @returns {Promise<boolean|number>} `true` if the shortlink has been deleted successfully or the HTTP error code
 */
export const deleteShortlink = (link) => {
    return new Promise(async (resolve, reject) => {
        // Remove possible root URL / path from the URL
        link = link.replace(window.location.href, '');
        var location = window.location.href.replace('http://', '').replace('https://', '');
        link = link.replace(location, '');
        
        var res = await fetch(`/api/shortlink/delete/${link}`, {
            method: 'DELETE'
        });
    
        if (!res.ok) {
            // Reject promise with HTTP error code
            reject(res.status);
        } else {
            // The shortlink has been deleted successfully
            resolve(true);
        }
    });
}