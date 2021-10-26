import { Heading } from '@chakra-ui/layout';
import React, { useContext } from 'react';
import CreateShortlink from '../../components/dashboard/CreateShortlink';
import DeleteShortlink from '../../components/dashboard/DeleteShortlink';
import LookupShortlink from '../../components/dashboard/LookupShortlink';
import { UserContext } from '../../context/UserContext';

const Home = () => {
    // Get user-specific states from global user context
    const { username } = useContext(UserContext);

    const getTimeGreeting = () => {
        var currentTime = new Date().getHours();

        if (currentTime < 12) return 'Good morning';
        if (currentTime < 18) return 'Good afternoon';
        if (currentTime < 23) return 'Good evening';
        return 'Good night';
    }


    const timeGreeting = getTimeGreeting();

    return (
        <>
            <Heading width='90%'>{ timeGreeting }, { username }.</Heading>
            
            <CreateShortlink />
            <LookupShortlink />
            <DeleteShortlink />
        </>
    );
}

export default Home;
