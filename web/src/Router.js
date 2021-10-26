import React, { useContext } from 'react';
import { BrowserRouter, Route, Switch, Redirect } from 'react-router-dom';
import { UserContext } from './context/UserContext';
import Dashboard from './pages/dashboard/Dashboard';
import Login from './pages/Login';
import ShortlinkRedirect from './pages/ShortlinkRedirect';
import Shortlinks from './pages/dashboard/Shortlinks';
import Home from './pages/dashboard/Home';

const Router = () => {
    // Get user-specific states from global user context
    const { loggedIn } = useContext(UserContext);
    
    return (
        <BrowserRouter>
            <div className='App'>
            <Switch>
                { /* TODO: Nested routing */}
                <Route path='/dashboard' exact>
                    { loggedIn ? <Dashboard component={ <Home /> } /> : <Redirect to='/login' /> }
                </Route>

                <Route path='/dashboard/shortlinks' exact>
                    { loggedIn ? <Dashboard component={ <Shortlinks /> } /> : <Redirect to='/login' /> }
                </Route>

                <Route path='/login' exact>
                    { loggedIn ? <Redirect to='/' /> : <Login /> }
                </Route>

                <Route path='/' exact>
                    { loggedIn ? <Redirect to='/dashboard' /> : <Redirect to='login' /> }
                </Route>

                <Route path='/'>
                    <ShortlinkRedirect />
                </Route>
            </Switch>

            </div>
        </BrowserRouter>
    );
}

export default Router;
