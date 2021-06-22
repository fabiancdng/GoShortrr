import {
    Flex,
    Text,
    Divider,
    Avatar,
    Heading,
    useBreakpointValue,
} from '@chakra-ui/react'
import { useContext } from 'react'
import {
    FiHome,
    FiUser,
    FiSettings,
    FiLink,
    FiInfo
} from 'react-icons/fi'
import NavItem from '../components/NavItem'
import { UserContext } from '../context/UserContext'

export default function Sidebar() {
    const { username, permissions } = useContext(UserContext)
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
            w="280px"
            flexDir="column"
        >
            <Flex
                mt={15}
                flexDir="column"
                alignItems="center"
                justifyContent="center"
            >
                <Heading textAlign="center">🔗</Heading>
                <Heading mt={5} size="lg" textAlign="center">GoShortrr</Heading>
            </Flex>

            <Flex
                h="95vh"
                pr={5}
                marginTop="2vh"
                w="280px"
                flexDir="column"
                justifyContent="space-between"
            >
                <Flex
                    p="5%"
                    flexDir="column"
                    w="100%"
                    alignItems="flex-start"
                    as="nav"
                >
                    <NavItem icon={FiHome} title="Dashboard" active />
                    <NavItem icon={FiLink} title="Shortlinks" />
                    <NavItem icon={FiUser} title="Users" disabled />
                    <NavItem icon={FiInfo} title="Infos" disabled />
                    <NavItem icon={FiSettings} title="Settings" disabled />
                </Flex>

                <Flex
                    p="5%"
                    flexDir="column"
                    w="100%"
                    alignItems="flex-start"
                >
                    <Divider />
                    <Flex mt={4} align="center">
                        <Avatar size="sm" />
                        <Flex flexDir="column" ml={4}>
                            <Heading as="h3" size="sm">{ username }</Heading>
                            <Text color="gray" size="xs">{ permissions === 0 ? 'User' : 'Admin' }</Text>
                        </Flex>
                    </Flex>
                </Flex>
            </Flex>
        </Flex>
    )
}