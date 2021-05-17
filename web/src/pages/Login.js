import { Button } from "@chakra-ui/button"
import { Box, Flex, Heading } from "@chakra-ui/layout"
import { FormControl, FormLabel, Input, Spinner } from '@chakra-ui/react'
import { useState } from "react"

const Login = () => {
    const [username, setUsername] = useState('')
    const [password, setPassword] = useState('')
    const [submitted, setSubmitted] = useState(false)

    const submitHandler = (e) => {
        e.preventDefault()
        setSubmitted(true)
        setUsername('')
        setPassword('')
    }

    return (
        <Flex mt={10} width="full" flexDirection="column" align="center" justifyContent="center">
            <Box p={8} minW={{ base: "90%", md: "600px" }} borderWidth={1} borderRadius={8} boxShadow="lg">
                <Box textAlign="center">
                    <Heading>Login</Heading>
                </Box>
                <Box my={4} textAlign="left">
                    <form onSubmit={e => submitHandler(e)}>
                    <FormControl>
                        <FormLabel>Username</FormLabel>
                        <Input value={username} type="email" placeholder="Your username" onChange={e => setUsername(e.target.value)} />
                    </FormControl>
                    <FormControl mt={6}>
                        <FormLabel>Password</FormLabel>
                        <Input value={password} type="password" placeholder="Your password" onChange={e => setPassword(e.target.value)} />
                    </FormControl>
                    {submitted ? <Button disabled={true} width="full" mt={4} type="submit"><Spinner mr={3} /> Signing in...</Button> : <Button width="full" mt={4} type="submit">Sign In</Button>}
                    </form>
                </Box>
            </Box>
        </Flex>
    )
}

export default Login
