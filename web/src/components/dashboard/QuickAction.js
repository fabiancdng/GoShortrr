import React, { useState } from 'react';
import { Box, Button, Heading, Flex, Input, Text } from '@chakra-ui/react';
import { useBreakpointValue } from '@chakra-ui/media-query';

const QuickAction = ({ title, subtitle, icon, color, placeholder, buttonLabel, handlerFunction }) => {
    // Whether the mobile or desktop version of the navbar is (or should be) displayed.
    const navSize = useBreakpointValue({ base: 'small', lg: 'large' });

    const [link, setLink] = useState('');

    return (
        <Box
            mt={10}
            width={ navSize === 'small' ? '90%' : '80%' }
            height='min'
            p={8}
            borderWidth={1}
            borderRadius={8}
            boxShadow='lg'
        >
            <Heading size='lg'>{ title }</Heading>
            <Text mt={0.5} pl={0.5} size='xs'>{ subtitle }</Text>
            <Flex
                flexDir={ navSize === 'small' ? 'column' : 'row' }
                mt={3}
            >
                <Input
                    placeholder={ placeholder }
                    size='md'
                    value={ link }
                    onChange={ e => setLink(e.target.value) }
                />

                <Button
                    mt={ navSize === 'small' ? 5 : 0 }
                    colorScheme={ color }
                    variant='outline'
                    leftIcon={ icon }
                    onClick={ e => { handlerFunction(link) } }
                >
                    { buttonLabel }
                </Button>
            </Flex>
        </Box>
    );
}

export default QuickAction;
