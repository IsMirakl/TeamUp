import { useMemo, useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import CreatePostForm from '../components/forms/CreatePostForm';
import Header from '../components/Header';
import type { Post } from '../types/Post';

const CreatePostPage = () => {
  const navigate = useNavigate();
  const [createdPost, setCreatedPost] = useState<Post | null>(null);
  const [formKey, setFormKey] = useState(0);

  const accessToken = useMemo(() => localStorage.getItem('accessToken'), []);

  return (
    <section className="min-h-screen bg-gradient-to-br from-amber-50 via-white to-sky-50">
      <Header />
      <main className="mx-auto w-full max-w-5xl px-6 py-10">
        <div className="mb-8">
          <p className="text-xs font-semibold tracking-[0.25em] text-slate-500 uppercase">
            TeamUP
          </p>
          <h1 className="mt-3 text-3xl font-semibold text-slate-900">Новый пост</h1>
          <p className="mt-2 max-w-2xl text-sm text-slate-600">
            Опишите, кого ищете и что хотите собрать. Чем конкретнее описание, тем проще
            людям откликнуться.
          </p>
        </div>

        {!accessToken ? (
          <div className="w-full max-w-2xl rounded-3xl border border-slate-200/80 bg-white/80 p-6 shadow-xl shadow-slate-900/10 backdrop-blur">
            <p className="text-sm font-semibold text-slate-900">
              Чтобы публиковать посты, нужно войти.
            </p>
            <p className="mt-2 text-sm text-slate-600">
              Перейдите на страницу входа и вернитесь сюда после авторизации.
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
          <div className="grid grid-cols-1 gap-8 lg:grid-cols-5">
            <div className="lg:col-span-3">
              <div className="rounded-3xl border border-slate-200/80 bg-white/80 p-6 shadow-xl shadow-slate-900/10 backdrop-blur">
                <CreatePostForm
                  key={formKey}
                  onCreated={post => {
                    setCreatedPost(post);
                  }}
                />
              </div>
            </div>

            <aside className="lg:col-span-2">
              <div className="rounded-3xl border border-slate-200/80 bg-white/70 p-6 shadow-xl shadow-slate-900/10 backdrop-blur">
                <p className="text-xs font-semibold tracking-[0.25em] text-slate-500 uppercase">
                  Подсказки
                </p>
                <ul className="mt-4 space-y-3 text-sm text-slate-700">
                  <li>Заголовок: роль или задача (например, "Ищу UI/UX на MVP").</li>
                  <li>Описание: формат, сроки, что уже сделано, стек.</li>
                  <li>Теги: 3-6 ключевых слов, через запятую.</li>
                </ul>
              </div>

              {createdPost ? (
                <div className="mt-6 rounded-3xl border border-emerald-200/70 bg-emerald-50/70 p-6 shadow-xl shadow-slate-900/10">
                  <p className="text-sm font-semibold text-emerald-900">Пост опубликован</p>
                  <p className="mt-2 text-sm text-emerald-900/80">{createdPost.title}</p>
                  <p className="mt-1 text-xs text-emerald-900/70">id: {createdPost.id}</p>
                  <div className="mt-5 flex flex-wrap items-center gap-3">
                    <button
                      type="button"
                      className="rounded-full bg-slate-900 px-5 py-2 text-sm font-semibold tracking-[0.15em] text-white uppercase shadow-lg shadow-slate-900/20 transition hover:bg-slate-800"
                      onClick={() => navigate('/home')}
                    >
                      На главную
                    </button>
                    <button
                      type="button"
                      className="rounded-full border border-slate-200 bg-white/70 px-5 py-2 text-sm font-semibold tracking-[0.15em] text-slate-700 uppercase transition hover:bg-white"
                      onClick={() => {
                        setCreatedPost(null);
                        setFormKey(v => v + 1);
                      }}
                    >
                      Еще пост
                    </button>
                  </div>
                </div>
              ) : null}
            </aside>
          </div>
        )}
      </main>
    </section>
  );
};

export default CreatePostPage;
