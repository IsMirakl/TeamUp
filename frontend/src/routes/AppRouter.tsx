import { lazy, Suspense } from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';

const HomePage = lazy(() => import('../pages/HomePage'));
const LoginPage = lazy(() => import('../pages/LoginPage'));
const RegisterPage = lazy(() => import('../pages/RegisterPage'));
const ProfilePage = lazy(() => import('../pages/ProfilePage'));
const CreatePostPage = lazy(() => import('../pages/CreatePostPage'));
const PostPage = lazy(() => import('../pages/PostPage'));
const MyPostsPage = lazy(() => import('../pages/MyPostsPage'));
const PostResponsesPage = lazy(() => import('../pages/PostResponsesPage'));

const PageLoader = () => (
  <div className="min-h-screen bg-gradient-to-br from-amber-50 via-white to-sky-50 px-6 py-10">
    <p className="text-sm text-slate-600">Загружаем…</p>
  </div>
);

const AppRouter: React.FC = () => {
  return (
    <Router>
      <Suspense fallback={<PageLoader />}>
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/login" element={<LoginPage />} />
          <Route path="/register" element={<RegisterPage />} />
          <Route path="/profile" element={<ProfilePage />} />
          <Route path="/posts/new" element={<CreatePostPage />} />
          <Route path="/posts/:id" element={<PostPage />} />
          <Route path="/my/posts" element={<MyPostsPage />} />
          <Route path="/posts/:id/responses" element={<PostResponsesPage />} />
          <Route path="/home" element={<HomePage />} />
        </Routes>
      </Suspense>
    </Router>
  );
};

export default AppRouter;
