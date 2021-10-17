import React, { useState } from 'react';
import { Button, Flex, Input, InputGroup, InputLeftElement, Modal, ModalBody, ModalCloseButton, ModalContent, ModalHeader, ModalOverlay, Text, useClipboard, useDisclosure } from '@chakra-ui/react';
import { FiCheck, FiLink } from 'react-icons/fi';
import QuickAction from './QuickAction';
import { createShortlink } from '../../adapters/ShortlinkAdapter';

const CreateShortlink = () => {
    // Whether or not the modal is opened and functions to open and close it
    const { isOpen, onOpen, onClose } = useDisclosure();
    // The generated shortlink returned by the API
    const [shortlink, setShortlink] = useState('');
    // Whether or not the shortlink has been copied to the clipboard and function to do so
    const { hasCopied, onCopy } = useClipboard(shortlink);

    const handleShortlinkCreation = (link) => {
        createShortlink(link)
            .then(short => {
                setShortlink(short);
                onOpen();
            })
            // TODO: Error handling in case shortlink creation failed
            .catch(httpErrorCode => {});
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
                handlerFunction={ handleShortlinkCreation }
            />

            <Modal size='4xl' onClose={ onClose } isOpen={ isOpen }>
                    <ModalOverlay />
                    <ModalContent pb={5}>
                        <ModalHeader>
                            <Flex color='green.300'>
                                <FiCheck size={25} />
                                <Text>Shortlink created!</Text>
                            </Flex>
                        </ModalHeader>

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
