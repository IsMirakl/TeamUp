import { useEffect } from 'react';
import { useAuth } from './hooks/useAuth';
import AppRouter from './routes/AppRouter';

function App() {
  const { initialize } = useAuth();

  useEffect(() => {
    initialize();
  }, [initialize]);
  return (
    <>
      <AppRouter />
    </>
  );
}

export default App;
