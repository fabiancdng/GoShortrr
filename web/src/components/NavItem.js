import React from 'react'
import {
    Flex,
    Text,
    Icon,
    Menu,
    MenuButton,
} from '@chakra-ui/react'

export default function NavItem({ icon, title, active, disabled }) {
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
                    color={disabled && 'gray.500'}
                    w="100%"
                    cursor={disabled ? 'not-allowed' : 'pointer'}
                    _hover={{ textDecor: 'none', backgroundColor: `${!active && !disabled ? 'gray.700' : ''}`}}
                >
                    <Flex>
                        <Icon as={icon} fontSize="xl"/>
                        <Text ml={5}>{title}</Text>
                    </Flex>
                </MenuButton>
            </Menu>
        </Flex>
    )
}