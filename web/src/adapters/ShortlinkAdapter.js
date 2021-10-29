/*
    Adapter for handling shortlink-related requests to the API.
*/

/**
 * @typedef ShortlinkData
 * @property {number} id Internal ID of the shortlink
 * @property {string} link The original (long) link behind the shortlink
 * @property {string} short The unique part of the shortlink
 * @property {user} user The internal ID of the user who created the shortlink
 * @property {string} created Timestamp (datetime in ISO format) of shortlink creation
 */

/**
 * Host and protocol of this instance (and if needed port).
 * Example: https://s.example.org:293/
 * @type {String}
 */
export const host = window.location.protocol + '//' + window.location.host + '/';

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
        // Remove possible root URL / path / slash from the URL
        link = link.replace(host, '');
        var location = host.replace('http://', '').replace('https://', '');
        link = link.replace(location, '').replace('/', '');
        
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
 * Gets all shortlinks of the currently logged-in user.
 * 
 * `[GET] /api/shortlink/list`
 * 
 * Reference: https://github.com/fabiancdng/GoShortrr/wiki/%F0%9F%94%8C-REST-API#get-apishortlinklist
 * 
 * @returns {Promise<Array.<ShortlinkData>|number>} The list of shortlinks or the HTTP error code
 */
 export const getShortlinkList = () => {
    return new Promise(async (resolve, reject) => {
        var res = await fetch(`/api/shortlink/list`);
        
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
            // Grab unique part of the shortlink, append the root URL, remove possible double slash,
            // and resolve the promise with it as value
            res = await res.json();
            let short = host.replace('dashboard', '') + res.short;
            // Remove possible double slash from returned short URL
            var shortArray = short.split('://');
            short = shortArray[0] + '://' + shortArray[1].replaceAll('//', '');
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
        link = link.replace(host, '');
        link = link.replace(window.location.href, '');
        var location = host.replace('http://', '').replace('https://', '');
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