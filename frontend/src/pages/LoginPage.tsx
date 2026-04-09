import { Link } from 'react-router-dom';
import LoginForm from '../components/forms/LoginForm';

const LoginPage = () => {
  return (
    <section className="relative flex min-h-screen items-center justify-center overflow-hidden bg-gradient-to-br from-amber-50 via-white to-sky-50 px-6 py-16">
      <div className="pointer-events-none absolute -top-40 right-10 h-80 w-80 rounded-full bg-sky-200/40 blur-3xl" />
      <div className="pointer-events-none absolute -bottom-32 left-0 h-72 w-72 rounded-full bg-amber-200/50 blur-3xl" />
      <div className="relative w-full max-w-md rounded-3xl border border-slate-200/80 bg-white/80 p-8 shadow-2xl shadow-slate-900/10 backdrop-blur">
        <div className="mb-6 text-center">
          <p className="text-xs font-semibold tracking-[0.25em] text-slate-500 uppercase">
            TeamUP
          </p>
          <h1 className="mt-3 text-3xl font-semibold text-slate-900">
            С возвращением
          </h1>
          <p className="mt-2 text-sm text-slate-600">
            Войдите, чтобы продолжить работу с командой.
          </p>
        </div>
        <LoginForm />
        <div className="mt-6 text-center text-sm text-slate-600">
          Нет аккаунта?{' '}
          <Link
            className="font-semibold text-sky-700 transition hover:text-sky-600"
            to={'/register'}
          >
            Создать
          </Link>
        </div>
      </div>
    </section>
  );
};

export default LoginPage;
