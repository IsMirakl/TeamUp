import { Route, BrowserRouter as Router, Routes } from 'react-router-dom';
import CreatePostPage from '../pages/CreatePostPage';
import HomePage from '../pages/HomePage';
import LoginPage from '../pages/LoginPage';
import RegisterPage from '../pages/RegisterPage';

const AppRouter: React.FC = () => {
  return (
    <Router>
      <Routes>
        <Route
          path="/"
          element={<HomePage />}
        />

        <Route
          path="/login"
          element={<LoginPage />}
        />
        <Route
          path="/register"
          element={<RegisterPage />}
        />

        <Route
          path="/posts/new"
          element={<CreatePostPage />}
        />

        <Route
          path="/home"
          element={<HomePage />}
        />
      </Routes>
    </Router>
  );
};

export default AppRouter;
