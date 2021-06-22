import React, { useState } from 'react'
import {
    Flex,
    Text,
    Divider,
    Avatar,
    Heading,
    useBreakpointValue,
    useColorMode
} from '@chakra-ui/react'
import {
    FiHome,
    FiCalendar,
    FiUser,
    FiDollarSign,
    FiBriefcase,
    FiSettings
} from 'react-icons/fi'
import { IoPawOutline } from 'react-icons/io5'
import NavItem from '../components/NavItem'

export default function Sidebar() {
    const navSize = useBreakpointValue({ base: "small", lg: "large" })

    return (
        <Flex
            display={navSize === 'small' ? 'none' : 'flex'}
            pos="sticky"
            left="5"
            h="95vh"
            pr={5}
            marginTop="2.5vh"
            borderRight='1px solid #4b4b4b'
            w="240px"
            flexDir="column"
            justifyContent="space-between"
        >
            <Heading mt={10} textAlign="center">🔗</Heading>
            <Flex
                p="5%"
                flexDir="column"
                w="100%"
                alignItems="flex-start"
                as="nav"
            >
                <NavItem icon={FiHome} title="Dashboard" description="This is the description for the dashboard." />
                <NavItem icon={FiCalendar} title="Calendar" active />
                <NavItem icon={FiUser} title="Clients" />
                <NavItem icon={IoPawOutline} title="Animals" />
                <NavItem icon={FiDollarSign} title="Stocks" />
                <NavItem icon={FiBriefcase} title="Reports" />
                <NavItem icon={FiSettings} title="Settings" />
            </Flex>

            <Flex
                p="5%"
                flexDir="column"
                w="100%"
                alignItems="flex-start"
                mb={4}
            >
                <Divider />
                <Flex mt={4} align="center">
                    <Avatar size="sm" src="avatar-1.jpg" />
                    <Flex flexDir="column" ml={4}>
                        <Heading as="h3" size="sm">Sylwia Weller</Heading>
                        <Text color="gray">Admin</Text>
                    </Flex>
                </Flex>
            </Flex>
        </Flex>
    )
}