import { useEffect } from 'react';
import { useLocation } from 'react-router';

const Redirect = () => {
    const shortlink = useLocation().pathname;

    useEffect(() => {
        const fetchShortlinkData = async () => {
            var res = await fetch(`/api/shortlink/get${shortlink}`);
            res = await res.json();

            window.location.replace(res.link);
        }
      
        fetchShortlinkData();
    }, [])

    return (
        <div>
            <p>Redirect in progress...</p>
        </div>
    );
}

export default Redirect;
