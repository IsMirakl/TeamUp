import { useEffect, useMemo, useState } from 'react';
import Header from '../components/Header';
import PostCard from '../features/posts/components/PostCard';
import PostsToolbar from '../features/posts/components/PostsToolbar';
import { usePost } from '../hooks/usePost';

const HomePage = () => {
  const { posts, fetchPosts, isLoading, error } = usePost();
  const [activeTag, setActiveTag] = useState<string | null>(null);
  const [searchQuery, setSearchQuery] = useState('');
  const [expandedPostId, setExpandedPostId] = useState<string | null>(null);
  const [copiedPostId, setCopiedPostId] = useState<string | null>(null);

  useEffect(() => {
    fetchPosts(50, 0);
  }, [fetchPosts]);

  const normalizedQuery = useMemo(() => searchQuery.trim().toLowerCase(), [searchQuery]);

  const filteredPosts = useMemo(() => {
    const byTag = activeTag
      ? posts.filter(post => post.tags?.some(tag => tag === activeTag))
      : posts;

    if (!normalizedQuery) return byTag;

    return byTag.filter(post => {
      const haystack = [
        post.title,
        post.description,
        post.author ?? '',
        ...(post.tags ?? []),
      ]
        .join(' ')
        .toLowerCase();

      return haystack.includes(normalizedQuery);
    });
  }, [posts, activeTag, normalizedQuery]);

  return (
    <section className="relative min-h-screen overflow-hidden bg-gradient-to-br from-amber-50 via-white to-sky-50">
      <div className="pointer-events-none absolute -top-48 right-10 h-96 w-96 rounded-full bg-sky-200/40 blur-3xl" />
      <div className="pointer-events-none absolute -bottom-40 left-0 h-96 w-96 rounded-full bg-amber-200/50 blur-3xl" />
      <div className="pointer-events-none absolute inset-0 opacity-60 [background:radial-gradient(90%_140%_at_0%_0%,rgba(14,116,144,0.14),transparent_60%),radial-gradient(90%_140%_at_100%_0%,rgba(234,179,8,0.12),transparent_55%)]" />

      <div className="relative">
        <Header />
      </div>

      <main className="relative mx-auto w-full max-w-6xl px-6 py-10">
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

        <PostsToolbar
          searchQuery={searchQuery}
          normalizedQuery={normalizedQuery}
          activeTag={activeTag}
          onChangeSearchQuery={setSearchQuery}
          onClearSearch={() => setSearchQuery('')}
          onResetFilters={() => {
            setActiveTag(null);
            setSearchQuery('');
          }}
        />

        {error ? (
          <div className="mb-6 rounded-3xl border border-rose-200/70 bg-rose-50/60 p-5 text-sm text-rose-800">
            {error}
          </div>
        ) : null}

        {isLoading ? (
          <p className="text-sm text-slate-500">Загружаем посты...</p>
        ) : null}

        {!isLoading && filteredPosts.length === 0 ? (
          <div className="rounded-3xl border border-slate-200/80 bg-white/85 p-8 shadow-xl shadow-slate-900/10 ring-1 ring-slate-900/5 backdrop-blur">
            <p className="text-sm font-semibold text-slate-900">
              {activeTag || normalizedQuery
                ? 'По вашему фильтру пока нет постов'
                : 'Пока нет постов'}
            </p>
            <p className="mt-2 text-sm text-slate-600">
              {activeTag || normalizedQuery
                ? 'Попробуйте изменить запрос, выбрать другой тег или сбросить фильтр.'
                : 'Создайте первый пост — он появится здесь.'}
            </p>
          </div>
        ) : (
          <div className="grid grid-cols-1 gap-6 md:grid-cols-2">
            {filteredPosts.map(post => (
              <PostCard
                key={post.id}
                post={post}
                activeTag={activeTag}
                expanded={expandedPostId === post.id}
                copied={copiedPostId === post.id}
                onToggleTag={clicked =>
                  setActiveTag(prev => (prev === clicked ? null : clicked))
                }
                onToggleExpanded={() =>
                  setExpandedPostId(prev => (prev === post.id ? null : post.id))
                }
                onRespond={async () => {
                  try {
                    await navigator.clipboard.writeText(
                      `Отклик на пост #${post.id}: ${post.title}`,
                    );
                    setCopiedPostId(post.id);
                    window.setTimeout(() => {
                      setCopiedPostId(current => (current === post.id ? null : current));
                    }, 1400);
                  } catch {
                    setCopiedPostId(post.id);
                    window.setTimeout(() => {
                      setCopiedPostId(current => (current === post.id ? null : current));
                    }, 1400);
                  }
                }}
              />
            ))}
          </div>
        )}
      </main>
    </section>
  );
};

export default HomePage;
