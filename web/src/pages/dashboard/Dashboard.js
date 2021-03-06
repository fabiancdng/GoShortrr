import React, { useEffect, useState } from 'react';
import { Flex, IconButton, useBreakpointValue } from '@chakra-ui/react';
import { FiMenu } from 'react-icons/fi';
import Sidebar from '../../components/dashboard/Sidebar/Sidebar';

const Dashboard = ({ component }) => {
    const [displayMobileNav, setDisplayMobileNav] = useState(false);
    const navSize = useBreakpointValue({ base: 'small', lg: 'large' });

    useEffect(() => {
        if (navSize !== 'small') setDisplayMobileNav(false);
    }, [navSize]);

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
                { component }
            </Flex>
        </Flex>
        </>
    );
}

export default Dashboard;
