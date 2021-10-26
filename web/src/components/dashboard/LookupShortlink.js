import React, { useState } from 'react';
import {
    Flex,
    Link,
    Modal,
    ModalBody,
    ModalCloseButton,
    ModalContent,
    ModalHeader,
    ModalOverlay,
    Table,
    Tbody,
    Td,
    Text,
    Th,
    Tr,
    useDisclosure } from '@chakra-ui/react';
import { FiSearch } from 'react-icons/fi';
import QuickAction from './QuickAction';
import { host, getShortlinkData } from '../../adapters/ShortlinkAdapter';

const LookupShortlink = () => {
    // Whether or not modal is opened and functions to open and close it
    const { isOpen, onOpen, onClose } = useDisclosure();
    // Whether the lookup failed or succeeded
    const [lookupStatus, setLookupStatus] = useState(false);
    // The data of the shortlink returned by the API
    const [shortlinkData, setShortlinkData] = useState('');

    const handleShortlinkLookup = (link) => {
        getShortlinkData(link)
            .then(data => {
                setShortlinkData(data);
                setLookupStatus(true);
                onOpen();
            })
            .catch(httpErrorCode => {
                setShortlinkData('');
                setLookupStatus(false);
                onOpen();
            });
    }

    return (
        <>
            <QuickAction
                title='Look up a shortlink'
                subtitle={ <b>Paste a shortlink here to see it's metadata.</b> }
                icon={ <FiSearch /> }
                color='blue'
                placeholder='Paste your shortlink here'
                buttonLabel='Look up'
                handlerFunction={ handleShortlinkLookup }
            />
            
            <Modal size='4xl' onClose={ onClose } isOpen={ isOpen }>
                <ModalOverlay />
                <ModalContent pb={5}>
                    <ModalHeader>
                        <Flex color='blue.300'>
                            <FiSearch size={25} />
                            <Text ml={2}>Shortlink Lookup</Text>
                        </Flex>
                    </ModalHeader>
                    
                    <ModalCloseButton />
                    
                    <ModalBody>
                        {
                            !lookupStatus
                            && <Text color='red.300'>Unable to look up shortlink. Make sure it exists.</Text>
                        }

                        {
                            shortlinkData !== '' && 
                            <Table variant='simple'>
                                <Tbody>
                                    <Tr>
                                        <Th>ID</Th>
                                        <Td>{ shortlinkData.id }</Td>
                                    </Tr>
                                    
                                    <Tr>
                                        <Th>Link</Th>
                                        <Td>
                                            <Link
                                                href={ shortlinkData.link }
                                                isExternal={true}
                                            >
                                                { shortlinkData.link }
                                            </Link>
                                        </Td>
                                    </Tr>

                                    <Tr>
                                        <Th>Shortlink</Th>
                                        <Td>
                                            <Link
                                                href={ host + shortlinkData.short }
                                                isExternal={true}
                                            >
                                                { host + shortlinkData.short }
                                            </Link>
                                        </Td>
                                    </Tr>
                                    
                                    <Tr>
                                        <Th>Created</Th>
                                        <Td>
                                            { shortlinkData.created.replace('T', ' - ').replace('Z', '') }
                                        </Td>
                                    </Tr>
                                </Tbody>
                            </Table>
                        }
                    </ModalBody>
                </ModalContent>
            </Modal>
        </>
    );
}

export default LookupShortlink;
