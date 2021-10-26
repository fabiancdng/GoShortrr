import React, { useContext } from 'react';
import { UserContext } from './context/UserContext';
import Router from './Router';

const App = () => {
  // Get user-specific states from global user context
  const { pending } = useContext(UserContext);

  // Login/user status isn't checked yet
  if (pending) return null;

  return <Router />;
}

export default App;
