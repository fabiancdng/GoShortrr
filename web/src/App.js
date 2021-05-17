import { Switch, Route, BrowserRouter } from 'react-router-dom'
import { ChakraProvider } from "@chakra-ui/react"
import Login from "./pages/Login"
import Header from './components/Header'

const App = () => {
  return (
    <BrowserRouter>
      <ChakraProvider>
        <div className="App">
          <Header />
          
          <Switch>
            <Route path="/login">
              <Login />
            </Route>

            <Route path="/">
              <p>Welcome!</p>
            </Route>
          </Switch>

        </div>
      </ChakraProvider>
    </BrowserRouter>
  );
}

export default App;
