import { Route, BrowserRouter as Router, Routes } from 'react-router-dom';
import CreatePostPage from '../pages/CreatePostPage';
import HomePage from '../pages/HomePage';
import LoginPage from '../pages/LoginPage';
import MyPostsPage from '../pages/MyPostsPage';
import PostResponsesPage from '../pages/PostResponsesPage';
import ProfilePage from '../pages/ProfilePage';
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
          path="/profile"
          element={<ProfilePage />}
        />

        <Route
          path="/posts/new"
          element={<CreatePostPage />}
        />

        <Route
          path="/my/posts"
          element={<MyPostsPage />}
        />

        <Route
          path="/posts/:id/responses"
          element={<PostResponsesPage />}
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
