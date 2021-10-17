import React, { useContext } from 'react';
import {
    Flex,
    Text,
    Divider,
    Avatar,
    Heading,
    useBreakpointValue,
    IconButton,
} from '@chakra-ui/react';
import {
    FiHome,
    FiUser,
    FiSettings,
    FiLink,
    FiInfo,
    FiLogOut
} from 'react-icons/fi';
import NavItem from './NavItem';
import { UserContext } from '../../../context/UserContext';
import { logoutUser } from '../../../adapters/UserAdapter';

export default function Sidebar({ mobileNav }) {
    const { username, permissions } = useContext(UserContext);
    const navSize = useBreakpointValue({ base: 'small', lg: 'large' });

    const logout = async () => {
        logoutUser()
            .then(loggedOut => {
                if (loggedOut) {
                    window.location.reload();
                }
            })
            .catch(err => window.location.reload());
    }

    return (
        <Flex
            display={ navSize === 'small' ? mobileNav ? 'block' : 'none' : 'flex' }
            left='5'
            pos={ mobileNav ? 'unset' : 'fixed' }
            h='95vh'
            pr={5}
            marginTop='2.5vh'
            borderRight='1px solid #4b4b4b'
            w={ mobileNav ? '100%' : '280px' }
            flexDir='column'
        >   
            <Flex
                mt={15}
                flexDir='column'
                alignItems='center'
                justifyContent='center'
            >
                <img width='30%' alt='GoShortrr logo' src='/goshortrr-logo-bg-circle.png' />
                <Heading mt={8} size='lg' textAlign='center'>GoShortrr</Heading>
            </Flex>

            <Flex
                h={ mobileNav ? '70vh' : '95vh' }
                pr={5}
                marginTop='2vh'
                w={ mobileNav ? '100%' : '280px' }
                flexDir='column'
                justifyContent='space-between'
            >
                <Flex
                    p='5%'
                    flexDir='column'
                    w='100%'
                    alignItems='flex-start'
                    as='nav'
                >
                    <NavItem icon={ FiHome } title='Dashboard' active />
                    <NavItem icon={ FiLink } title='Shortlinks' />
                    <NavItem icon={ FiUser } title='Users' disabled />
                    <NavItem icon={ FiInfo } title='Infos' disabled />
                    <NavItem icon={ FiSettings } title='Settings' disabled />
                </Flex>

                <Flex
                    p='5%'
                    flexDir='column'
                    w='100%'
                    alignItems='flex-start'
                >
                    <Divider />
                    <Flex mt={4} align='center' width='100%' justifyContent='space-between'>
                        <Flex flexDir='row'>
                            <Avatar size='sm' />
                            <Flex flexDir='column' ml={4}>
                                <Heading as='h3' size='sm'>{ username }</Heading>
                                <Text color='gray' size='xs'>{ permissions === 0 ? 'User' : 'Admin' }</Text>
                            </Flex>
                        </Flex>
                        
                        <IconButton rounded='md' onClick={ logout } variant='ghost' icon={ <FiLogOut /> } />
                    </Flex>
                </Flex>
            </Flex>
        </Flex>
    );
}