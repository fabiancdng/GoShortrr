import React from 'react'
import {
    Flex,
    Text,
    Icon,
    Link,
    Menu,
    MenuButton,
    MenuList,
    Button
} from '@chakra-ui/react'

export default function NavItem({ icon, title, description, active, navSize }) {
    return (
        <Flex
            mt={30}
            flexDir="column"
            w="100%"
            alignItems={navSize == "small" ? "center" : "flex-start"}
        >
            <Menu placement="right">
                <MenuButton
                    backgroundColor={active && "gray.600"}
                    p={4}
                    borderRadius={8}
                    w="100%"
                    _hover={{ textDecor: 'none' }}
                >
                    <Flex>
                        <Icon as={icon} fontSize="xl" color={active && "gray.800"} />
                        <Text ml={5} display={navSize == "small" ? "none" : "flex"}>{title}</Text>
                    </Flex>
                </MenuButton>
            </Menu>
        </Flex>
    )
}