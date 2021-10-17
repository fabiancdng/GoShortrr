import React, { useState } from 'react';
import {
    Flex,
    Modal,
    ModalBody,
    ModalCloseButton,
    ModalContent,
    ModalHeader,
    ModalOverlay,
    Text,
    useDisclosure
} from '@chakra-ui/react';
import { FiTrash, FiTrash2 } from 'react-icons/fi';
import QuickAction from './QuickAction';
import { deleteShortlink } from '../../adapters/ShortlinkAdapter';

const DeleteShortlink = () => {
    // Whether or not the modal is opened and functions to open and close it
    const { isOpen, onOpen, onClose } = useDisclosure();
    // Whether the delete request failed or succeeded
    const [deleteStatus, setDeleteStatus] = useState(false);

    const handleShortlinkDeletion = (link) => {
        deleteShortlink(link)
            .then(deleted => {
                if (deleted) {
                    setDeleteStatus(true);
                    onOpen();
                } else {
                    setDeleteStatus(false);
                    onOpen();
                }
            })
            .catch(httpErrorCode => {
                setDeleteStatus(false);
                onOpen();
            });
    }

    return (
        <>
            <QuickAction
                title='Quickly revoke a shortlink'
                subtitle={ <p><b>Revoke/delete a shortlink by pasting it.</b> For a full list of your shortlinks, switch to the 'Shortlinks' tab.</p> }
                icon={ <FiTrash2 /> }
                color='red'
                placeholder='Paste your shortlink here'
                buttonLabel='Delete'
                handlerFunction={ handleShortlinkDeletion }
            />

            <Modal size='4xl' onClose={ onClose } isOpen={ isOpen }>
                <ModalOverlay />
                <ModalContent pb={5}>
                    <ModalHeader>
                        <Flex color='red.300'>
                            <FiTrash size={25} />
                            <Text ml={2}>Revoke Shortlink</Text>
                        </Flex>
                    </ModalHeader>
                    
                    <ModalCloseButton />
                    
                    <ModalBody>
                        {
                            deleteStatus
                            ? <Text color='green.300'>Shortlink has been deleted successfully.</Text>
                            : <Text color='red.300'>Unable to delete your shortlink. Make sure it exists and you have permissions to delete it.</Text>
                        }
                    </ModalBody>
                </ModalContent>
            </Modal>
        </>
    );
}

export default DeleteShortlink;
