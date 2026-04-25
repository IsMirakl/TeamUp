import type { Post } from '../../../types/Post';
import TagPill from './TagPill';

type PostCardProps = {
  post: Post;
  activeTag: string | null;
  expanded: boolean;
  copied: boolean;
  onToggleTag: (tag: string) => void;
  onToggleExpanded: () => void;
  onRespond: () => void;
};

const PostCard = ({
  post,
  activeTag,
  expanded,
  copied,
  onToggleTag,
  onToggleExpanded,
  onRespond,
}: PostCardProps) => {
  return (
    <article className="group rounded-3xl border border-slate-200/80 bg-white/85 p-6 shadow-xl shadow-slate-900/10 ring-1 ring-slate-900/5 backdrop-blur transition duration-200 hover:-translate-y-0.5 hover:bg-white hover:shadow-2xl hover:shadow-slate-900/10">
      <div className="flex items-start justify-between gap-4">
        <div className="min-w-0">
          <h2 className="truncate text-lg font-semibold text-slate-900 transition group-hover:text-sky-800">
            {post.title}
          </h2>
          <p className="mt-2 text-xs font-semibold tracking-[0.2em] text-slate-500 uppercase">
            {post.author ? `Автор: ${post.author}` : 'Автор: —'}
          </p>
        </div>
      </div>

      <p
        className={`mt-3 text-sm leading-6 text-slate-700 ${
          expanded ? '' : 'max-h-32 overflow-hidden'
        }`}
      >
        {post.description}
      </p>

      {post.tags?.length ? (
        <div className="mt-4 flex flex-wrap gap-2">
          {(expanded ? post.tags : post.tags.slice(0, 8)).map(tag => (
            <TagPill
              key={`${post.id}-${tag}`}
              tag={tag}
              active={activeTag === tag}
              onToggle={onToggleTag}
            />
          ))}
        </div>
      ) : null}

      <div className="mt-5 flex flex-wrap items-center gap-3">
        <button
          type="button"
          className="rounded-full border border-slate-200/80 bg-white/70 px-5 py-2 text-sm font-semibold text-slate-700 shadow-sm transition hover:bg-white"
          onClick={onToggleExpanded}
        >
          {expanded ? 'Свернуть' : 'Подробнее'}
        </button>
        <button
          type="button"
          className="rounded-full bg-slate-900 px-5 py-2 text-sm font-semibold text-white shadow-lg shadow-slate-900/20 transition hover:bg-slate-800"
          onClick={onRespond}
        >
          Откликнуться
        </button>
        {copied ? (
          <span className="text-xs font-semibold text-emerald-700">Скопировано</span>
        ) : (
          <span className="text-xs text-slate-500">Кнопка копирует шаблон отклика</span>
        )}
      </div>
    </article>
  );
};

export default PostCard;
