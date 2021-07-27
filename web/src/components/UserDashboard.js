import { Flex, Heading } from "@chakra-ui/react"
import CreateShortlink from "./CreateShortlink"
import DeleteShortlink from "./DeleteShortlink"
import LookupShortlink from "./LookupShortlink"
import Sidebar from "./Sidebar"

const UserDashboard = () => {
    return (
        <>
        <Flex>
            <Sidebar />
            <Flex flexDir="column" alignItems="center" width="100%" mt={10} ml="280px">
                <Heading width="90%">Good evening, Fabian.</Heading>

                <CreateShortlink />
                <LookupShortlink />
                <DeleteShortlink />
            </Flex>
        </Flex>
        </>
    )
}

export default UserDashboard
