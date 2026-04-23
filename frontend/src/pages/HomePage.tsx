import { useEffect } from 'react';
import Header from '../components/Header';
import { usePost } from '../hooks/usePost';

const Tag = ({ tag }: { tag: string }) => (
  <span className="rounded-full border border-slate-200 bg-white/70 px-3 py-1 text-xs font-semibold text-slate-700">
    {tag}
  </span>
);

const HomePage = () => {
  const { posts, fetchPosts, isLoading, error } = usePost();

  useEffect(() => {
    fetchPosts(50, 0);
  }, [fetchPosts]);

  return (
    <>
      <header>
        <Header />
      </header>

      <main className="mx-auto w-full max-w-6xl px-6 py-10">
        <div className="mb-8">
          <p className="text-xs font-semibold tracking-[0.25em] text-slate-500 uppercase">
            TeamUP
          </p>
          <h1 className="mt-3 text-3xl font-semibold text-slate-900">
            Лента постов
          </h1>
          <p className="mt-2 max-w-2xl text-sm text-slate-600">
            Идеи, запросы и вакансии в командах. Откликайтесь по тегам и
            собирайте состав.
          </p>
        </div>

        {error ? (
          <div className="mb-6 rounded-3xl border border-rose-200/70 bg-rose-50/60 p-5 text-sm text-rose-800">
            {error}
          </div>
        ) : null}

        {isLoading ? (
          <p className="text-sm text-slate-500">Загружаем посты...</p>
        ) : null}

        {!isLoading && posts.length === 0 ? (
          <div className="rounded-3xl border border-slate-200/80 bg-white/70 p-8 shadow-xl shadow-slate-900/10 backdrop-blur">
            <p className="text-sm font-semibold text-slate-900">
              Пока нет постов
            </p>
            <p className="mt-2 text-sm text-slate-600">
              Создайте первый пост — он появится здесь.
            </p>
          </div>
        ) : (
          <div className="grid grid-cols-1 gap-6 md:grid-cols-2">
            {posts.map(post => (
              <article
                key={post.id}
                className="group rounded-3xl border border-slate-200/80 bg-white/70 p-6 shadow-xl shadow-slate-900/10 backdrop-blur transition hover:bg-white"
              >
                <h2 className="text-lg font-semibold text-slate-900 transition group-hover:text-sky-800">
                  {post.title}
                </h2>
                <p className="mt-2 text-xs font-semibold tracking-[0.2em] text-slate-500 uppercase">
                  {post.author ? `Автор: ${post.author}` : 'Автор: —'}
                </p>
                <p className="mt-3 max-h-32 overflow-hidden text-sm leading-6 text-slate-700">
                  {post.description}
                </p>

                {post.tags?.length ? (
                  <div className="mt-4 flex flex-wrap gap-2">
                    {post.tags.slice(0, 8).map(tag => (
                      <Tag
                        key={`${post.id}-${tag}`}
                        tag={tag}
                      />
                    ))}
                  </div>
                ) : null}
              </article>
            ))}
          </div>
        )}
      </main>
    </>
  );
};

export default HomePage;
