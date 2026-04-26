import { memo } from 'react';
import { Link } from 'react-router-dom';
import type { Post } from '../../../types/Post';
import TagPill from './TagPill';

type PostCardProps = {
  post: Post;
  activeTag: string | null;
  responded: boolean;
  onToggleTag: (tag: string) => void;
  onRespond: () => void;
};

const PostCard = ({
  post,
  activeTag,
  responded,
  onToggleTag,
  onRespond,
}: PostCardProps) => {
  return (
    <article className="group relative overflow-hidden rounded-3xl border border-slate-200/70 bg-[linear-gradient(135deg,rgba(255,255,255,0.92),rgba(240,249,255,0.66),rgba(255,251,235,0.62))] p-6 shadow-xl shadow-slate-900/10 ring-1 ring-slate-900/5 backdrop-blur transition duration-200 hover:-translate-y-0.5 hover:shadow-2xl hover:shadow-slate-900/10">
      <div className="pointer-events-none absolute inset-0 opacity-70 [background:radial-gradient(120%_120%_at_0%_0%,rgba(14,116,144,0.14),transparent_60%),radial-gradient(110%_110%_at_100%_0%,rgba(234,179,8,0.12),transparent_58%)]" />
      <div className="relative">
      <div className="flex items-start justify-between gap-4">
        <div className="min-w-0">
          <Link to={`/posts/${post.id}`} className="block">
            <h2 className="truncate text-lg font-semibold text-slate-900 transition group-hover:text-sky-800">
              {post.title}
            </h2>
          </Link>
          <p className="mt-2 text-xs font-semibold tracking-[0.2em] text-slate-500 uppercase">
            {post.author ? `Автор: ${post.author}` : 'Автор: —'}
          </p>
        </div>
      </div>

      <p
        className="mt-3 text-sm leading-6 text-slate-700 max-h-32 overflow-hidden"
      >
        {post.description}
      </p>

      {post.tags?.length ? (
        <div className="mt-4 flex flex-wrap gap-2">
          {post.tags.slice(0, 8).map(tag => (
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
        <Link
          className="rounded-full border border-slate-200/80 bg-white/70 px-5 py-2 text-sm font-semibold text-slate-700 shadow-sm transition hover:bg-white"
          to={`/posts/${post.id}`}
        >
          Подробнее
        </Link>
        <button
          type="button"
          className="rounded-full bg-slate-900 px-5 py-2 text-sm font-semibold text-white shadow-lg shadow-slate-900/20 transition hover:bg-slate-800"
          onClick={onRespond}
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
  );
};

export default memo(PostCard);
