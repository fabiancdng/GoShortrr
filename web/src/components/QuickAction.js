import { Box, Button, Heading, HStack, Input, Text } from "@chakra-ui/react"

const QuickAction = ({title, subtitle, icon, color, placeholder, buttonLabel}) => {
    return (
        <Box mt={10} width="80%" height="min" p={8} borderWidth={1} borderRadius={8} boxShadow="lg">
            <Heading size="lg">{ title }</Heading>
            <Text mt={0.5} pl={0.5} size="xs">{ subtitle }</Text>
            <HStack mt={3}>
                <Input placeholder={ placeholder } size="md" />
                <Button colorScheme={color} variant="outline" leftIcon={icon}>{ buttonLabel }</Button>
            </HStack>
        </Box>
    )
}

export default QuickAction
