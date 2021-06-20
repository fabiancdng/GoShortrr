import { MoonIcon, SunIcon } from "@chakra-ui/icons"
import { Box, Flex, Heading } from "@chakra-ui/layout"
import { IconButton, useColorMode } from '@chakra-ui/react'

const Header = () => {
    const { colorMode, toggleColorMode } = useColorMode();

    return (
        <Box position="absolute" top={0} right={0} textAlign="right" py={4} mr={5}>
            <Flex p={4}  justifyContent="space-between">
                <IconButton
                    icon={colorMode === 'light' ? <MoonIcon /> : <SunIcon />}
                    onClick={toggleColorMode}
                    variant="outline"
                    borderRadius={5}
                />
            </Flex>
        </Box>
    )
}

export default Header
