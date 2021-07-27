import { Flex, Heading } from "@chakra-ui/react"
import CreateShortlink from "./CreateShortlink"
import DeleteShortlink from "./DeleteShortlink"
import LookupShortlink from "./LookupShortlink"
import Sidebar from "./Sidebar"

const UserDashboard = ({ username }) => {
    const getTimeGreeting = () => {
        var currentTime = new Date().getHours()

        if(currentTime < 12) return "Good morning, "
        if(currentTime < 18) return "Good afternoon, "
        if(currentTime < 23) return "Good evening, "
        return "Good night, "
    }

    const timeGreeting = getTimeGreeting()

    return (
        <>
        <Flex>
            <Sidebar />
            <Flex flexDir="column" alignItems="center" width="100%" mt={10} ml="280px">
                <Heading width="90%">{timeGreeting} {username}.</Heading>

                <CreateShortlink />
                <LookupShortlink />
                <DeleteShortlink />
            </Flex>
        </Flex>
        </>
    )
}

export default UserDashboard
