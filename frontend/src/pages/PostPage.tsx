import { useEffect, useMemo, useState } from 'react';
import { Link, useNavigate, useParams } from 'react-router-dom';
import Header from '../components/Header';
import RespondModal from '../features/posts/components/RespondModal';
import TagPill from '../features/posts/components/TagPill';
import { usePost } from '../hooks/usePost';
import type { Post } from '../types/Post';

const PostPage = () => {
  const navigate = useNavigate();
  const { id } = useParams();
  const { post, fetchPost, isLoading, error } = usePost();

  const [responded, setResponded] = useState(false);
  const [respondModalPost, setRespondModalPost] = useState<Post | null>(null);
  const [activeTag, setActiveTag] = useState<string | null>(null);

  useEffect(() => {
    if (!id) return;
    void fetchPost(id);
  }, [fetchPost, id]);

  const tags = useMemo(() => post?.tags ?? [], [post?.tags]);

  return (
    <section className="relative min-h-screen overflow-hidden bg-gradient-to-br from-amber-50 via-white to-sky-50">
      <div className="pointer-events-none absolute -top-48 right-10 h-96 w-96 rounded-full bg-sky-200/40 blur-3xl" />
      <div className="pointer-events-none absolute -bottom-40 left-0 h-96 w-96 rounded-full bg-amber-200/50 blur-3xl" />
      <div className="pointer-events-none absolute inset-0 opacity-60 [background:radial-gradient(90%_140%_at_0%_0%,rgba(14,116,144,0.12),transparent_60%),radial-gradient(90%_140%_at_100%_0%,rgba(234,179,8,0.10),transparent_55%)]" />

      <div className="relative">
        <Header />
      </div>

      <main className="relative mx-auto w-full max-w-4xl px-6 py-10">
        <div className="mb-8 flex flex-wrap items-center justify-between gap-4">
          <div>
            <p className="text-xs font-semibold tracking-[0.25em] text-slate-500 uppercase">TeamUP</p>
            <h1 className="mt-3 text-3xl font-semibold text-slate-900">Пост</h1>
          </div>
          <div className="flex flex-wrap items-center gap-3">
            <Link
              to={'/home'}
              className="rounded-full border border-slate-200 bg-white/70 px-5 py-2 text-sm font-semibold text-slate-700 shadow-sm transition hover:bg-white"
            >
              Лента
            </Link>
            <button
              type="button"
              onClick={() => navigate(-1)}
              className="rounded-full bg-slate-900 px-5 py-2 text-sm font-semibold text-white shadow-lg shadow-slate-900/20 transition hover:bg-slate-800"
            >
              Назад
            </button>
          </div>
        </div>

        {error ? (
          <div className="mb-6 rounded-3xl border border-rose-200/70 bg-rose-50/60 p-5 text-sm text-rose-800">
            {error}
          </div>
        ) : null}

        {isLoading || !post ? (
          <div className="rounded-3xl border border-slate-200/80 bg-white/85 p-8 shadow-xl shadow-slate-900/10 ring-1 ring-slate-900/5 backdrop-blur">
            <p className="text-sm text-slate-600">Загружаем…</p>
          </div>
        ) : (
          <article className="relative overflow-hidden rounded-3xl border border-slate-200/70 bg-[linear-gradient(135deg,rgba(255,255,255,0.92),rgba(240,249,255,0.66),rgba(255,251,235,0.62))] p-8 shadow-xl shadow-slate-900/10 ring-1 ring-slate-900/5 backdrop-blur">
            <div className="pointer-events-none absolute inset-0 opacity-70 [background:radial-gradient(120%_120%_at_0%_0%,rgba(14,116,144,0.14),transparent_60%),radial-gradient(110%_110%_at_100%_0%,rgba(234,179,8,0.12),transparent_58%)]" />
            <div className="relative">
              <h2 className="text-2xl font-semibold text-slate-900">{post.title}</h2>
              <p className="mt-2 text-xs font-semibold tracking-[0.2em] text-slate-500 uppercase">
                {post.author ? `Автор: ${post.author}` : 'Автор: —'}
              </p>

              <p className="mt-6 whitespace-pre-wrap text-sm leading-7 text-slate-700">
                {post.description}
              </p>

              {tags.length ? (
                <div className="mt-6 flex flex-wrap gap-2">
                  {tags.map(tag => (
                    <TagPill
                      key={`${post.id}-${tag}`}
                      tag={tag}
                      active={activeTag === tag}
                      onToggle={clicked => setActiveTag(prev => (prev === clicked ? null : clicked))}
                    />
                  ))}
                </div>
              ) : null}

              <div className="mt-8 flex flex-wrap items-center gap-3">
                <button
                  type="button"
                  className="rounded-full bg-slate-900 px-6 py-2.5 text-sm font-semibold text-white shadow-lg shadow-slate-900/20 transition hover:bg-slate-800"
                  onClick={() => setRespondModalPost(post)}
                >
                  Откликнуться
                </button>
                {responded ? (
                  <span className="text-xs font-semibold text-emerald-700">Отклик отправлен</span>
                ) : (
                  <span className="text-xs text-slate-500">Отправьте сообщение автору</span>
                )}
              </div>
            </div>
          </article>
        )}
      </main>

      <RespondModal
        open={Boolean(respondModalPost)}
        post={respondModalPost}
        onClose={() => setRespondModalPost(null)}
        onSent={() => {
          setResponded(true);
          window.setTimeout(() => setResponded(false), 2200);
        }}
      />
    </section>
  );
};

export default PostPage;
