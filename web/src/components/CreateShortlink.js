import { FiLink } from "react-icons/fi"
import QuickAction from "./QuickAction"

const CreateShortlink = () => {
    return (
        <QuickAction
            title="Quickly create a shortlink"
            subtitle={<b>Paste. Click the button. Done.</b>}
            icon={<FiLink />}
            color="green"
            placeholder="Paste your long link here"
            buttonLabel="Shorten"
        />
    )
}

export default CreateShortlink
