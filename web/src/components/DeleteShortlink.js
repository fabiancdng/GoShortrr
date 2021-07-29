import { FiTrash2 } from "react-icons/fi"
import QuickAction from "./QuickAction"

const DeleteShortlink = () => {
    return (
        <QuickAction
            title="Quickly revoke a shortlink"
            subtitle={<p><b>Revoke/delete a shortlink by pasting it.</b> For a full list of your shortlinks, switch to the 'Shortlinks' tab.</p>}
            icon={<FiTrash2 />}
            color="red"
            placeholder="Paste your shortlink here"
            buttonLabel="Delete"
        />
    )
}

export default DeleteShortlink
