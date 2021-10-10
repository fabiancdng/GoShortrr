import { Flex, Link, Modal, ModalBody, ModalCloseButton, ModalContent, ModalHeader, ModalOverlay, Table, Tbody, Td, Text, Th, Tr, useDisclosure } from '@chakra-ui/react';
import { useState } from 'react';
import { FiSearch } from 'react-icons/fi';
import QuickAction from './QuickAction';

const LookupShortlink = () => {
    const { isOpen, onOpen, onClose } = useDisclosure();
    const [lookupStatus, setLookupStatus] = useState(false);
    const [shortlinkData, setShortlinkData] = useState('');

    const fetchShortlinkData = async (link) => {
        link = link.replace(window.location.href, '');
        var location = window.location.href.replace('http://', '').replace('https://', '');
        link = link.replace(location, '');
        var linkData = await fetch(`/api/shortlink/get/${link}`);
        
        if (linkData.ok) {
            linkData = await linkData.json();
            setShortlinkData(linkData);
            setLookupStatus(true);
            onOpen();
        } else {
            setShortlinkData('');
            setLookupStatus(false);
            onOpen();
        }
    }

    return (
        <>
            <QuickAction
                title='Look up a shortlink'
                subtitle={<b>Paste a shortlink here to see it's metadata.</b>}
                icon={<FiSearch />}
                color='blue'
                placeholder='Paste your shortlink here'
                buttonLabel='Look up'
                handlerFunction={fetchShortlinkData}
            />
            
            <Modal size='4xl' onClose={onClose} isOpen={isOpen}>
                <ModalOverlay />
                <ModalContent pb={5}>
                    <ModalHeader><Flex color='blue.300'><FiSearch size={25} /><Text ml={2}>Shortlink Lookup</Text></Flex></ModalHeader>
                    <ModalCloseButton />
                    <ModalBody>
                    {!lookupStatus && <Text color='red.300'>Unable to look up shortlink. Make sure it exists.</Text>}

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
                                    <Td><Link href={ shortlinkData.link } isExternal={true}>{ shortlinkData.link }</Link></Td>
                                </Tr>
                                <Tr>
                                    <Th>Shortlink</Th>
                                    <Td><Link href={ window.location.href + shortlinkData.short } isExternal={true}>{ window.location.href + shortlinkData.short }</Link></Td>
                                </Tr>
                                <Tr>
                                    <Th>Created</Th>
                                    <Td>{ shortlinkData.created.replace('T', ' - ').replace('Z', '') }</Td>
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
