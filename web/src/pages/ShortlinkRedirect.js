import React, { useEffect } from 'react';
import { useLocation } from 'react-router';
import { getShortlinkData } from '../adapters/ShortlinkAdapter';

const Redirect = () => {
    // Read current URL (as that is the shortlink to look up)
    const shortlink = useLocation().pathname;

    useEffect(() => {
        getShortlinkData(shortlink)
            .then(shortlinkData => {
                window.location.replace(shortlinkData.link);
            })
            // TODO: Error handling in case shortlink lookup failed
            .catch(httpErrorCode => {});
    // eslint-disable-next-line
    }, []);

    return (
        <div>
            <p>Redirect in progress...</p>
        </div>
    );
}

export default Redirect;
