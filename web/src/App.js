import { useContext, useEffect, useState } from 'react'
import { Switch, Route, BrowserRouter, Redirect } from 'react-router-dom'
import Login from "./pages/Login"
import UserOnlyRoute from './components/UserOnlyRoute'
import { UserContext } from './context/UserContext'
import UserDashboard from './components/UserDashboard'

const App = () => {
  const { username, setUsername, permissions, setPermissions, loggedIn, setLoggedIn } = useContext(UserContext)
  // True if the API hasn't responded yet (to the /auth/user request)
  const [pending, setPending] = useState(true)

  useEffect(async () => {
    var res = await fetch('/api/auth/user', {
      method: 'POST',
      credentials: 'include'
    })
    
    if(res.status === 401) {
          setLoggedIn(false)
          setPending(false)
    } else {
      res = await res.json()
      setLoggedIn(true)
      setUsername(res.username)
      setPermissions(res.role)
      setPending(false)
    }
  }, [])

  if(pending) return null

  return (
    <BrowserRouter>
        <div className="App">
          <Switch>
            <UserOnlyRoute loggedIn={loggedIn} path="/" exact>
              <UserDashboard />
            </UserOnlyRoute>

            <Route path="/login" exact>
              {loggedIn ? <Redirect to="/" /> : <Login />}
            </Route>
          </Switch>

        </div>
    </BrowserRouter>
  );
}

export default App;
