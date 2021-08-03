import { Flex, Modal, ModalBody, ModalCloseButton, ModalContent, ModalHeader, ModalOverlay, Text, useDisclosure } from "@chakra-ui/react"
import { useState } from "react"
import { FiTrash, FiTrash2 } from "react-icons/fi"
import QuickAction from "./QuickAction"

const DeleteShortlink = () => {
    const { isOpen, onOpen, onClose } = useDisclosure()
    const [deleteStatus, setDeleteStatus] = useState(false)

    const deleteShortlink = async (link) => {
        link = link.replace(window.location.href, "")
        var location = window.location.href.replace("http://", "").replace("https://", "")
        link = link.replace(location, "")
        var deleteRequest = await fetch(`/api/shortlink/delete/${link}`, {
            method: 'DELETE'
        })

        if (deleteRequest.ok) {
            setDeleteStatus(true)
            onOpen()
        } else {
            setDeleteStatus(false)
            onOpen()
        }
    }

    return (
    <>
        <QuickAction
            title="Quickly revoke a shortlink"
            subtitle={<p><b>Revoke/delete a shortlink by pasting it.</b> For a full list of your shortlinks, switch to the 'Shortlinks' tab.</p>}
            icon={<FiTrash2 />}
            color="red"
            placeholder="Paste your shortlink here"
            buttonLabel="Delete"
            handlerFunction={deleteShortlink}
        />

        <Modal size="4xl" onClose={onClose} isOpen={isOpen}>
            <ModalOverlay />
            <ModalContent pb={5}>
                <ModalHeader><Flex color="red.300"><FiTrash size={25} /><Text ml={2}>Revoke Shortlink</Text></Flex></ModalHeader>
                <ModalCloseButton />
                <ModalBody>
                    {deleteStatus ? <Text color="green.300">Shortlink has been deleted successfully.</Text> : <Text color="red.300">Unable to delete your shortlink. Make sure it exists and you have permissions to delete it.</Text>}
                </ModalBody>
            </ModalContent>
        </Modal>
    </>
    )
}

export default DeleteShortlink
