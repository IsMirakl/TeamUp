import { useEffect } from 'react';
import { useAuth } from './hooks/useAuth';
import AppRouter from './routes/AppRouter';

function App() {
  const { checkAuth } = useAuth();

  useEffect(() => {
    checkAuth();
  }, []);
  return (
    <>
      <AppRouter />
    </>
  );
}

export default App;
