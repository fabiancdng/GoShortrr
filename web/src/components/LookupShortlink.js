import { HStack, Link, Modal, ModalBody, ModalCloseButton, ModalContent, ModalHeader, ModalOverlay, Table, TableCaption, Tbody, Td, Text, Th, Thead, Tr, useDisclosure } from "@chakra-ui/react"
import { useState } from "react"
import { FiSearch } from "react-icons/fi"
import QuickAction from "./QuickAction"

const LookupShortlink = () => {
    const { isOpen, onOpen, onClose } = useDisclosure()
    const [shortlinkData, setShortlinkData] = useState('')

    const fetchShortlinkData = async (link) => {
        link = link.replace(window.location.href, "")
        var location = window.location.href.replace("http://", "").replace("https://", "")
        link = link.replace(location, "")
        var linkData = await fetch(`/api/shortlink/get/${link}`)
        linkData = await linkData.json()
        setShortlinkData(linkData)

        onOpen()
    }

    return (
    <>
        <QuickAction
            title="Look up a shortlink"
            subtitle={<b>Paste a shortlink here to see it's metadata.</b>}
            icon={<FiSearch />}
            color="blue"
            placeholder="Paste your shortlink here"
            buttonLabel="Look up"
            handlerFunction={fetchShortlinkData}
        />
        {
            shortlinkData !== '' && 
            <Modal size="4xl" onClose={onClose} isOpen={isOpen}>
            <ModalOverlay />
            <ModalContent pb={5}>
                <ModalHeader><HStack color="blue.300"><FiSearch size={28} /> <Text>Shortlink Lookup</Text></HStack></ModalHeader>
                <ModalCloseButton />
                <ModalBody>
                <Table variant="simple">
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
                            <Td>{ shortlinkData.created.replace("T", " - ").replace("Z", "") }</Td>
                        </Tr>
                    </Tbody>
                </Table>
                </ModalBody>
            </ModalContent>
        </Modal>
        }
    </>
    )
}

export default LookupShortlink
