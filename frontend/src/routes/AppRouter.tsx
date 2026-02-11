import { Route, BrowserRouter as Router, Routes } from 'react-router-dom';
import LoginPage from '../pages/LoginPage';

const AppRouter: React.FC = () => {
  return (
    <Router>
      <Routes>
        <Route
          path="/login"
          element={<LoginPage />}
        />
      </Routes>
    </Router>
  );
};

export default AppRouter;
