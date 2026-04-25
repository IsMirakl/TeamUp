import { useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import Header from '../components/Header';
import { useAuth } from '../hooks/useAuth';
import { useProfile } from '../hooks/useProfile';

const ROLE_LABELS: Record<string, string> = {
  user: 'Пользователь',
  admin: 'Администратор',
  team_lead: 'Тимлид',
};

const PLAN_LABELS: Record<string, string> = {
  Free: 'Free',
  Pro: 'Pro',
  Enterprise: 'Enterprise',
};

const labelValue = (value: unknown, dict: Record<string, string>) => {
  if (typeof value !== 'string' || !value) return '—';
  return dict[value] ?? value;
};

const ProfilePage = () => {
  const navigate = useNavigate();
  const { user, logout } = useAuth();
  const { profile, getMyProfile, isLoading, error } = useProfile();

  const isAuthenticated = !!localStorage.getItem('accessToken');

  useEffect(() => {
    if (!isAuthenticated) return;
    void getMyProfile();
  }, [getMyProfile, isAuthenticated]);

  const data =
    profile ??
    (user
      ? {
          name: user.name,
          email: user.email,
          role: user.role,
          subscriptionPlan: user.subscriptionPlan,
          avatar: user.avatarUrl,
        }
      : null);

  return (
    <section className="relative min-h-screen overflow-hidden bg-gradient-to-br from-amber-50 via-white to-sky-50">
      <div className="pointer-events-none absolute -top-48 left-10 h-96 w-96 rounded-full bg-amber-200/50 blur-3xl" />
      <div className="pointer-events-none absolute right-0 -bottom-40 h-96 w-96 rounded-full bg-sky-200/40 blur-3xl" />

      <div className="relative">
        <Header />
      </div>

      <main className="relative mx-auto w-full max-w-5xl px-6 py-10">
        <div className="mb-8">
          <p className="text-xs font-semibold tracking-[0.25em] text-slate-500 uppercase">
            TeamUP
          </p>
          <h1 className="mt-3 text-3xl font-semibold text-slate-900">
            Профиль
          </h1>
          <p className="mt-2 max-w-2xl text-sm text-slate-600">
            Настройки аккаунта и информация о подписке.
          </p>
        </div>

        {!isAuthenticated ? (
          <div className="w-full max-w-2xl rounded-3xl border border-slate-200/80 bg-white/80 p-6 shadow-xl ring-1 shadow-slate-900/10 ring-slate-900/5 backdrop-blur">
            <p className="text-sm font-semibold text-slate-900">Нужно войти</p>
            <p className="mt-2 text-sm text-slate-600">
              Перейдите на страницу входа, чтобы увидеть свой профиль.
            </p>
            <div className="mt-5 flex flex-wrap items-center gap-3">
              <Link
                className="rounded-full bg-slate-900 px-5 py-2 text-sm font-semibold tracking-[0.15em] text-white uppercase shadow-lg shadow-slate-900/20 transition hover:bg-slate-800"
                to={'/login'}
              >
                Войти
              </Link>
              <button
                type="button"
                className="rounded-full border border-slate-200 bg-white/70 px-5 py-2 text-sm font-semibold tracking-[0.15em] text-slate-700 uppercase transition hover:bg-white"
                onClick={() => navigate('/home')}
              >
                На главную
              </button>
            </div>
          </div>
        ) : (
          <div className="w-full rounded-3xl border border-slate-200/80 bg-white/85 p-7 shadow-xl ring-1 shadow-slate-900/10 ring-slate-900/5 backdrop-blur">
            {error ? (
              <div className="mb-5 rounded-2xl border border-rose-200/70 bg-rose-50/70 p-4 text-sm text-rose-800">
                {error}
              </div>
            ) : null}

            <div className="flex flex-col gap-6 sm:flex-row sm:items-start">
              <div className="h-18 w-18 shrink-0 overflow-hidden rounded-2xl border border-slate-200/80 bg-gradient-to-br from-slate-900 to-slate-700 shadow-lg shadow-slate-900/10 sm:h-20 sm:w-20">
                {data?.avatar ? (
                  <img
                    src={data.avatar}
                    alt="Avatar"
                    className="h-full w-full object-cover"
                  />
                ) : (
                  <div className="grid h-full w-full place-items-center text-xl font-semibold text-white">
                    {(data?.name?.trim()?.[0] ?? 'T').toUpperCase()}
                  </div>
                )}
              </div>

              <div className="min-w-0 flex-1">
                <div className="flex flex-col gap-4 lg:flex-row lg:items-start lg:justify-between">
                  <div className="min-w-0">
                    <p className="text-xs font-semibold tracking-[0.2em] text-slate-500 uppercase">
                      Аккаунт
                    </p>
                    <h2 className="mt-2 truncate text-2xl font-semibold text-slate-900">
                      {data?.name ?? '—'}
                    </h2>
                    <p className="mt-1 truncate text-sm text-slate-600">
                      {data?.email ?? '—'}
                    </p>
                    {isLoading ? (
                      <p className="mt-3 text-xs text-slate-500">
                        Загружаем профиль…
                      </p>
                    ) : null}
                  </div>

                  <div className="flex flex-wrap items-center gap-3">
                    <button
                      type="button"
                      className="rounded-full bg-slate-900 px-5 py-2 text-sm font-semibold text-white shadow-lg shadow-slate-900/20 transition hover:bg-slate-800"
                      onClick={() => {
                        logout();
                        navigate('/home');
                      }}
                    >
                      Выйти
                    </button>
                  </div>
                </div>

                <div className="mt-6 grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-3">
                  <div className="rounded-2xl border border-slate-200/80 bg-white/70 p-4 shadow-sm">
                    <p className="text-xs font-semibold tracking-[0.18em] text-slate-500 uppercase">
                      Роль
                    </p>
                    <p className="mt-2 text-sm font-semibold text-slate-900">
                      {labelValue(data?.role, ROLE_LABELS)}
                    </p>
                  </div>
                  <div className="rounded-2xl border border-slate-200/80 bg-white/70 p-4 shadow-sm">
                    <p className="text-xs font-semibold tracking-[0.18em] text-slate-500 uppercase">
                      Подписка
                    </p>
                    <p className="mt-2 text-sm font-semibold text-slate-900">
                      {labelValue(data?.subscriptionPlan, PLAN_LABELS)}
                    </p>
                  </div>
                  <div className="rounded-2xl border border-slate-200/80 bg-white/70 p-4 shadow-sm">
                    <p className="text-xs font-semibold tracking-[0.18em] text-slate-500 uppercase">
                      Статус
                    </p>
                    <p className="mt-2 text-sm font-semibold text-slate-900">
                      {isLoading ? 'Обновляем…' : 'Активен'}
                    </p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        )}
      </main>
    </section>
  );
};

export default ProfilePage;
