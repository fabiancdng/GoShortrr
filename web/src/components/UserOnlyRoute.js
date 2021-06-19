import { Route, Redirect } from 'react-router-dom'

const UserOnlyRoute = ({ path, loggedIn, children }) => {
    if(loggedIn) {
        return(
            <Route path={path}>
                {children}
            </Route>
        )
    } else {
        return <Redirect to="/login" />
    }
}

export default UserOnlyRoute
