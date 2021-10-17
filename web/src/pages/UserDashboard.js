import React, { useEffect, useState } from 'react';
import { Flex, Heading, IconButton, useBreakpointValue } from '@chakra-ui/react';
import { FiMenu } from 'react-icons/fi';
import CreateShortlink from '../components/user-dashboard/CreateShortlink';
import DeleteShortlink from '../components/user-dashboard/DeleteShortlink';
import LookupShortlink from '../components/user-dashboard/LookupShortlink';
import Sidebar from '../components/user-dashboard/Sidebar/Sidebar';

const UserDashboard = ({ username }) => {
    const [displayMobileNav, setDisplayMobileNav] = useState(false);
    const navSize = useBreakpointValue({ base: 'small', lg: 'large' });

    const getTimeGreeting = () => {
        var currentTime = new Date().getHours();

        if (currentTime < 12) return 'Good morning';
        if (currentTime < 18) return 'Good afternoon';
        if (currentTime < 23) return 'Good evening';
        return 'Good night';
    }

    useEffect(() => {
        if (navSize !== 'small') setDisplayMobileNav(false);
    }, [navSize]);

    const timeGreeting = getTimeGreeting();

    return (
        <>
        { 
            navSize === 'small'
            && <IconButton
                rounded='md'
                variant='ghost'
                onClick={ e => setDisplayMobileNav(displayMobileNav ? false : true) }
                ml={5}
                mt={5}
                icon={ <FiMenu /> }
            />
        }
            
        <Flex>
            <Sidebar mobileNav={ displayMobileNav } />
            <Flex
                display={ displayMobileNav ? 'none' : 'flex' }
                flexDir='column'
                alignItems='center'
                width='100%'
                mt={10}
                ml={ navSize === 'small' ? '10px' : '280px' }
            >
                <Heading width='90%'>{ timeGreeting }, { username }.</Heading>
                <CreateShortlink />
                <LookupShortlink />
                <DeleteShortlink />
            </Flex>
        </Flex>
        </>
    );
}

export default UserDashboard;
