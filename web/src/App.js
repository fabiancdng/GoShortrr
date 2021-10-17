import React, { useContext } from 'react';
import { Switch, Route, BrowserRouter, Redirect } from 'react-router-dom';
import Login from './pages/Login';
import UserOnlyRoute from './components/shared/UserOnlyRoute';
import { UserContext } from './context/UserContext';
import UserDashboard from './pages/UserDashboard';
import ShortlinkRedirect from './pages/ShortlinkRedirect';

const App = () => {
  // Get user-specific states from global user context
  const { username, loggedIn, pending } = useContext(UserContext);

  // Login/user status isn't checked yet
  if(pending) return null;

  return (
    <BrowserRouter>
        <div className='App'>
          <Switch>
            <UserOnlyRoute loggedIn={ loggedIn } path='/' exact>
              <UserDashboard username={ username } />
            </UserOnlyRoute>

            <Route path='/login' exact>
              { loggedIn ? <Redirect to='/' /> : <Login /> }
            </Route>

            <Route path='/'>
              <ShortlinkRedirect />
            </Route>
          </Switch>

        </div>
    </BrowserRouter>
  );
}

export default App;
