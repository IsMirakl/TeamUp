import { Link } from 'react-router-dom';
import { useAuth } from '../hooks/useAuth';

const Header = () => {
  const { user } = useAuth();

  const isAuthenticated = !!user;

  return (
    <header className="relative w-full border-b border-slate-200/70 bg-gradient-to-r from-amber-50 via-white to-sky-50">
      <div className="pointer-events-none absolute inset-0 opacity-70 [background:radial-gradient(90%_140%_at_0%_0%,rgba(14,116,144,0.18),transparent_60%)]" />
      <div className="relative mx-auto flex w-full max-w-6xl flex-col gap-4 px-6 py-4 md:flex-row md:items-center md:justify-between">
        <Link
          className="font-miranda text-2xl font-semibold tracking-tight text-slate-900 transition hover:text-sky-700 md:text-3xl"
          to={'/home'}
        >
          TeamUP
        </Link>
        <nav
          className="flex flex-wrap items-center gap-6 font-miranda text-sm font-semibold text-slate-700 md:text-base"
          aria-label="Основная навигация"
        >
          <Link
            className="transition hover:text-sky-700"
            to={'/'}
          >
            Главная
          </Link>
          <Link
            className="transition hover:text-sky-700"
            to={'/posts/new'}
          >
            Создать пост
          </Link>
          <Link
            className="transition hover:text-sky-700"
            to={'/'}
          >
            Мои посты
          </Link>
          <Link
            className="transition hover:text-sky-700"
            to={'/'}
          >
            Команда
          </Link>
        </nav>
        <div className="flex items-center gap-3">
          {isAuthenticated ? (
            <div className="rounded-full border border-slate-300/70 bg-white/70 px-4 py-2 text-sm font-semibold text-slate-700 shadow-sm">
              Профиль
            </div>
          ) : (
            <Link
              className="rounded-full bg-slate-900 px-4 py-2 text-sm font-semibold text-white shadow-lg shadow-slate-900/20 transition hover:bg-slate-800"
              to={'/login'}
            >
              Вход
            </Link>
          )}
        </div>
      </div>
    </header>
  );
};

export default Header;
