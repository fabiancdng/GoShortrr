import { Button } from '@chakra-ui/button';
import { Box, Flex, Heading } from '@chakra-ui/layout';
import { FormControl, FormLabel, Input, Spinner, Text } from '@chakra-ui/react';
import { useContext, useState } from 'react';
import { UserContext } from '../context/UserContext';

const Login = () => {
    const { setUsername, setLoggedIn, setPermissions } = useContext(UserContext);

    const [usernameInput, setUsernameInput] = useState('');
    const [passwordInput, setPasswordInput] = useState('');
    const [status, setStatus] = useState(0);
    
    /*
        Possible status values:

        0 = unsubmitted
        1 = submitted but no response yet
        2 = submitted but unauthorized
        3 = submitted but error
        4 = submitted and session created
    */

    const submitHandler = async (e) => {
        e.preventDefault();
        setStatus(1);
        setUsernameInput('');
        setPasswordInput('');

        var res = await fetch('/api/auth/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({username: usernameInput, password: passwordInput})
        });

        if(res.status === 200) {
            setStatus(4);
            res = await fetch('/api/auth/user', {
                method: 'POST',
                credentials: 'include'
            });

            if(res.status === 401) {
                setLoggedIn(false);
            } else {
                res = await res.json();
                setLoggedIn(true);
                setUsername(res.username);
                setPermissions(res.role);
            }
        } else if(res.status === 401) setStatus(2);
        else setStatus(3);
    }

    return (
        <Flex mt={10} width='full' flexDirection='column' align='center' justifyContent='center'>
            <Flex
                mt={15}
                mb={35}
                flexDir='column'
                alignItems='center'
                justifyContent='center'
            >
                <img width='30%' alt='GoShortrr logo' src='/goshortrr-logo-bg-circle.png' />
            </Flex>

            <Box p={8} minW={{ base: '90%', md: '600px' }} borderWidth={1} borderRadius={8} boxShadow='lg'>
                <Box textAlign='center'>
                    <Heading>Login</Heading>
                </Box>
                <Box my={4} textAlign='left'>
                    <form onSubmit={e => submitHandler(e)}>
                    {status === 2 ? <Text textAlign='center' my={5} color='red'>Invalid username or password.</Text> : status === 3 && <Text my={5} textAlign='center' color='red'>Internal Server Error! Please try again later.</Text>}
                    <FormControl>
                        <FormLabel>Username</FormLabel>
                        <Input value={usernameInput} type='text' placeholder='Your username' onChange={e => setUsernameInput(e.target.value)} />
                    </FormControl>
                    <FormControl mt={6}>
                        <FormLabel>Password</FormLabel>
                        <Input value={passwordInput} type='password' placeholder='Your password' onChange={e => setPasswordInput(e.target.value)} />
                    </FormControl>
                    {status === 1 ? <Button disabled={true} width='full' mt={4} type='submit'><Spinner mr={3} /> Signing in...</Button> : <Button width='full' mt={4} type='submit'>Sign In</Button>}
                    </form>
                </Box>
            </Box>
        </Flex>
    );
}

export default Login;
