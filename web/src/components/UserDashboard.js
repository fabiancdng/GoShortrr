import { Box, Button, Flex, Heading, HStack, Input, Text } from "@chakra-ui/react"
import { FiLink, FiSearch, FiTrash2 } from "react-icons/fi"
import Sidebar from "./Sidebar"

const UserDashboard = () => {
    return (
        <>
        <Flex>
            <Sidebar />
            <Flex flexDir="column" alignItems="center" width="100%" mt={10}>
                <Heading width="90%">Good evening, Fabian.</Heading>
                <Box mt={15} width="80%" height="min" p={8} borderWidth={1} borderRadius={8} boxShadow="lg">
                    <Heading size="lg">Quickly create a shortlink</Heading>
                    <Text mt={0.5} pl={0.5} size="xs"><b>Paste. Click the button. Done.</b></Text>
                    <HStack mt={3}>
                        <Input placeholder="Paste your long link here" size="md" />
                        <Button colorScheme="green" variant="outline" leftIcon={<FiLink />}>Shorten</Button>
                    </HStack>
                </Box>

                <Box mt={14} width="80%" height="min" p={8} borderWidth={1} borderRadius={8} boxShadow="lg">
                    <Heading size="lg">Look up a shortlink</Heading>
                    <Text mt={0.5} pl={0.5} size="xs"><b>Paste a shortlink here to see it's metadata.</b></Text>
                    <HStack mt={3}>
                        <Input placeholder="Paste your shortlink here" size="md" />
                        <Button colorScheme="blue" variant="outline" leftIcon={<FiSearch />}>Look up</Button>
                    </HStack>
                </Box>

                <Box mt={14} width="80%" height="min" p={8} borderWidth={1} borderRadius={8} boxShadow="lg">
                    <Heading size="lg">Quickly revoke a shortlink</Heading>
                    <Text mt={0.5} pl={0.5} size="xs"><b>Revoke/delete a shortlink by pasting it.</b> For a full list of your shortlinks, switch to the 'Shortlinks' tab.</Text>
                    <HStack mt={3}>
                        <Input placeholder="Paste your shortlink here" size="md" />
                        <Button colorScheme="red" variant="outline" leftIcon={<FiTrash2 />}>Delete</Button>
                    </HStack>
                </Box>
            </Flex>
        </Flex>
        </>
    )
}

export default UserDashboard
