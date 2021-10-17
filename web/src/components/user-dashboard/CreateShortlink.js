import React, { useState } from 'react';
import { Button, Flex, Input, InputGroup, InputLeftElement, Modal, ModalBody, ModalCloseButton, ModalContent, ModalHeader, ModalOverlay, Text, useClipboard, useDisclosure } from '@chakra-ui/react';
import { FiCheck, FiLink } from 'react-icons/fi';
import QuickAction from './QuickAction';

const CreateShortlink = () => {
    const { isOpen, onOpen, onClose } = useDisclosure();
    const [shortlink, setShortlink] = useState('');
    const { hasCopied, onCopy } = useClipboard(shortlink);

    const createShortlink = async (link) => {
        var short = await fetch('/api/shortlink/create', {
            method: 'POST',
            credentials: 'include',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({link: link})
        });

        short = await short.json();
        short = window.location.href + short.short;

        setShortlink(short);

        onOpen();
    }

    return (
        <>
            <QuickAction
                title='Quickly create a shortlink'
                subtitle={ <b>Paste. Click the button. Done.</b> }
                icon={ <FiLink /> }
                color='green'
                placeholder='Paste your long link here'
                buttonLabel='Shorten'
                handlerFunction={ createShortlink }
            />

            <Modal size='4xl' onClose={ onClose } isOpen={ isOpen }>
                    <ModalOverlay />
                    <ModalContent pb={5}>
                        <ModalHeader><Flex color='green.300'><FiCheck size={25} /> <Text>Shortlink created!</Text></Flex></ModalHeader>
                        <ModalCloseButton />
                        <ModalBody>
                            <InputGroup>
                                <InputLeftElement
                                    pointerEvents='none'
                                    children={ <FiLink color='gray.300' /> }
                                />
                                <Input isReadOnly={true} type='text' value={ shortlink } />
                                <Button onClick={ onCopy } ml={2}>
                                    { hasCopied ? 'Copied' : 'Copy' }
                                </Button>
                            </InputGroup>
                        </ModalBody>
                    </ModalContent>
                </Modal>
        </>
    );
}

export default CreateShortlink;
