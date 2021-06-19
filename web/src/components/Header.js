import { MoonIcon, SunIcon } from "@chakra-ui/icons"
import { Box, Flex, Heading } from "@chakra-ui/layout"
import { IconButton, useColorMode } from '@chakra-ui/react'

const Header = () => {
    const { colorMode, toggleColorMode } = useColorMode();

    return (
        <Box textAlign="right" py={4} mr={5}>
        <Flex p={4}  justifyContent="space-between">
            <Heading size="lg">GoShortrr</Heading>
            <IconButton
            icon={colorMode === 'light' ? <MoonIcon /> : <SunIcon />}
            onClick={toggleColorMode}
            variant="ghost"
            borderRadius={5}
            />
        </Flex>
        </Box>
    )
}

export default Header
