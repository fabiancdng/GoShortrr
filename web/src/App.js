import { useContext, useEffect, useState } from 'react'
import { Switch, Route, BrowserRouter, Redirect } from 'react-router-dom'
import Login from "./pages/Login"
import Header from './components/Header'
import UserOnlyRoute from './components/UserOnlyRoute'
import { UserContext } from './context/UserContext'

const App = () => {
  const { username, setUsername, permissions, setPermissions, loggedIn, setLoggedIn } = useContext(UserContext)
  // True if the API hasn't responded yet (to the /auth/user request)
  const [pending, setPending] = useState(true)

  useEffect(() => {
    fetch('/api/auth/user', {
      method: 'POST',
      credentials: 'include'
    })
      .then(async res => {
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
      })
  }, [])

  if(pending) return null

  return (
    <BrowserRouter>
        <div className="App">
          <Header />
          
          <Switch>
            <UserOnlyRoute loggedIn={loggedIn} path="/" exact>
              <p>Welcome, {username}!</p>
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
