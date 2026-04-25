type PostsToolbarProps = {
  searchQuery: string;
  normalizedQuery: string;
  activeTag: string | null;
  onChangeSearchQuery: (value: string) => void;
  onClearSearch: () => void;
  onResetFilters: () => void;
};

const PostsToolbar = ({
  searchQuery,
  normalizedQuery,
  activeTag,
  onChangeSearchQuery,
  onClearSearch,
  onResetFilters,
}: PostsToolbarProps) => {
  return (
    <div className="mb-6 grid grid-cols-1 gap-3 md:grid-cols-[1fr_auto] md:items-center">
      <div className="rounded-3xl border border-slate-200/80 bg-white/80 shadow-xl shadow-slate-900/10 ring-1 ring-slate-900/5 backdrop-blur">
        <div className="flex items-center gap-3 px-5 py-3">
          <div className="grid h-10 w-10 place-items-center rounded-2xl bg-slate-900 text-white shadow-lg shadow-slate-900/20">
            <svg
              viewBox="0 0 24 24"
              className="h-5 w-5"
              fill="none"
              stroke="currentColor"
              strokeWidth="2"
              strokeLinecap="round"
              strokeLinejoin="round"
              aria-hidden="true"
            >
              <circle cx="11" cy="11" r="7" />
              <path d="M21 21l-4.3-4.3" />
            </svg>
          </div>
          <div className="min-w-0 flex-1">
            <label
              htmlFor="posts-search"
              className="block text-xs font-semibold tracking-[0.18em] text-slate-500 uppercase"
            >
              Поиск по ленте
            </label>
            <div className="mt-1 flex items-center gap-2">
              <input
                id="posts-search"
                value={searchQuery}
                onChange={e => onChangeSearchQuery(e.target.value)}
                placeholder='Заголовок, теги, "дизайн", "react"…'
                className="w-full bg-transparent text-sm font-semibold text-slate-900 placeholder:text-slate-400 focus:outline-none"
              />
              {searchQuery.trim() ? (
                <button
                  type="button"
                  className="rounded-full border border-slate-200/80 bg-white/70 px-3 py-1 text-xs font-semibold text-slate-700 shadow-sm transition hover:bg-white"
                  onClick={onClearSearch}
                >
                  Очистить
                </button>
              ) : null}
            </div>
          </div>
        </div>
      </div>

      {activeTag || normalizedQuery ? (
        <div className="flex flex-wrap items-center gap-3">
          {activeTag ? (
            <div className="rounded-2xl border border-slate-200/70 bg-white/80 px-4 py-2 text-sm text-slate-700 shadow-sm ring-1 ring-slate-900/5 backdrop-blur">
              <span className="text-xs font-semibold tracking-[0.18em] text-slate-500 uppercase">
                Тег
              </span>
              <span className="ml-2 font-semibold text-slate-900">{activeTag}</span>
            </div>
          ) : null}
          {normalizedQuery ? (
            <div className="rounded-2xl border border-slate-200/70 bg-white/80 px-4 py-2 text-sm text-slate-700 shadow-sm ring-1 ring-slate-900/5 backdrop-blur">
              <span className="text-xs font-semibold tracking-[0.18em] text-slate-500 uppercase">
                Запрос
              </span>
              <span className="ml-2 font-semibold text-slate-900">{normalizedQuery}</span>
            </div>
          ) : null}
          <button
            type="button"
            className="rounded-full border border-slate-200/80 bg-white/70 px-4 py-2 text-sm font-semibold text-slate-700 shadow-sm transition hover:bg-white"
            onClick={onResetFilters}
          >
            Сбросить
          </button>
        </div>
      ) : null}
    </div>
  );
};

export default PostsToolbar;
