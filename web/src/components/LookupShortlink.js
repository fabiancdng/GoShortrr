import { FiSearch } from "react-icons/fi"
import QuickAction from "./QuickAction"

const LookupShortlink = () => {
    return (
        <QuickAction
            title="Look up a shortlink"
            subtitle={<b>Paste a shortlink here to see it's metadata.</b>}
            icon={<FiSearch />}
            color="blue"
            placeholder="Paste your shortlink here"
            buttonLabel="Look up"
        />
    )
}

export default LookupShortlink
