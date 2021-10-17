import { Heading } from '@chakra-ui/layout';
import React from 'react';
import CreateShortlink from './CreateShortlink';
import DeleteShortlink from './DeleteShortlink';
import LookupShortlink from './LookupShortlink';

const Home = ({ username }) => {
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
