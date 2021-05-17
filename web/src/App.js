import { ChakraProvider } from "@chakra-ui/react"
import Login from "./pages/Login"

const App = () => {
  return (
    <ChakraProvider>
      <div className="App">
        <Login />
      </div>
    </ChakraProvider>
  );
}

export default App;
