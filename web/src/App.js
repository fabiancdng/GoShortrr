import { useContext, useEffect } from 'react'
import { Switch, Route, BrowserRouter } from 'react-router-dom'
import Login from "./pages/Login"
import Header from './components/Header'
import UserOnlyRoute from './components/UserOnlyRoute'
import { UserContext } from './context/UserContext'

const App = () => {
  const { username, setUsername, permissions, setPermissions, loggedIn, setLoggedIn } = useContext(UserContext)

  useEffect(() => {
    fetch('/api/auth/user', {
      method: 'POST',
      credentials: 'include'
    })
      .then(async res => {
        if(res.status === 401) {
          setLoggedIn(false)
        } else {
          console.log(await res.json())
        }
      })
  }, [])

  return (
    <BrowserRouter>
        <div className="App">
          <Header />
          
          <Switch>
            <Route path="/login">
              <Login />
            </Route>

            <UserOnlyRoute loggedIn={loggedIn} path="/">
              <p>Welcome, {username}!</p>
            </UserOnlyRoute>
          </Switch>

        </div>
    </BrowserRouter>
  );
}

export default App;
