import React from 'react'
import {
    Flex,
    Text,
    Icon,
    Menu,
    MenuButton,
} from '@chakra-ui/react'

export default function NavItem({ icon, title, description, active }) {
    return (
        <Flex
            mt={30}
            flexDir="column"
            w="100%"
            alignItems="flex-start"
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
                        <Text ml={5}>{title}</Text>
                    </Flex>
                </MenuButton>
            </Menu>
        </Flex>
    )
}