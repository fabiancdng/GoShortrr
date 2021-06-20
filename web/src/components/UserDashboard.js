import { useContext } from "react"
import { UserContext } from "../context/UserContext"
import Header from './Header'
import { IconButton } from "@chakra-ui/button"
import { Text, Flex } from "@chakra-ui/layout"
import { useDisclosure } from "@chakra-ui/react"
import Sidebar from "./Sidebar"
import { FiMenu } from "react-icons/fi"

const UserDashboard = () => {
    const { username, setUsername, permissions, setPermissions, loggedIn, setLoggedIn } = useContext(UserContext)
    const { isOpen, onToggle } = useDisclosure()

    return (
        <>
            <Header />
            <Sidebar />
        </>
    )
}

export default UserDashboard
