import { useDisclosure } from "@chakra-ui/react"
import { useState } from "react"
import { FiTrash2 } from "react-icons/fi"
import QuickAction from "./QuickAction"

const DeleteShortlink = () => {
    const { isOpen, onOpen, onClose } = useDisclosure()

    const deleteShortlink = async (link) => {
        link = link.replace(window.location.href, "")
        var location = window.location.href.replace("http://", "").replace("https://", "")
        link = link.replace(location, "")
        await fetch(`/api/shortlink/delete/${link}`, {
            method: 'DELETE'
        })

        onOpen()
    }

    return (
        <QuickAction
            title="Quickly revoke a shortlink"
            subtitle={<p><b>Revoke/delete a shortlink by pasting it.</b> For a full list of your shortlinks, switch to the 'Shortlinks' tab.</p>}
            icon={<FiTrash2 />}
            color="red"
            placeholder="Paste your shortlink here"
            buttonLabel="Delete"
            handlerFunction={deleteShortlink}
        />
    )
}

export default DeleteShortlink
