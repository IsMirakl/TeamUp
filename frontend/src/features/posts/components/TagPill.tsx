const TAG_STYLES = [
  {
    base: 'border-sky-200/70 bg-sky-50/80 text-sky-800 ring-sky-200/50',
    active:
      'border-sky-300 bg-sky-100/80 text-sky-900 ring-sky-300/70 shadow-sky-500/10',
  },
  {
    base: 'border-amber-200/70 bg-amber-50/80 text-amber-900 ring-amber-200/50',
    active:
      'border-amber-300 bg-amber-100/80 text-amber-950 ring-amber-300/70 shadow-amber-500/10',
  },
  {
    base: 'border-emerald-200/70 bg-emerald-50/80 text-emerald-900 ring-emerald-200/50',
    active:
      'border-emerald-300 bg-emerald-100/80 text-emerald-950 ring-emerald-300/70 shadow-emerald-500/10',
  },
  {
    base: 'border-rose-200/70 bg-rose-50/80 text-rose-900 ring-rose-200/50',
    active:
      'border-rose-300 bg-rose-100/80 text-rose-950 ring-rose-300/70 shadow-rose-500/10',
  },
  {
    base: 'border-indigo-200/70 bg-indigo-50/80 text-indigo-900 ring-indigo-200/50',
    active:
      'border-indigo-300 bg-indigo-100/80 text-indigo-950 ring-indigo-300/70 shadow-indigo-500/10',
  },
  {
    base: 'border-slate-200/80 bg-slate-50/80 text-slate-800 ring-slate-200/50',
    active:
      'border-slate-300 bg-white text-slate-900 ring-slate-300/70 shadow-slate-900/10',
  },
] as const;

const hashToIndex = (value: string, mod: number) => {
  let hash = 0;
  for (let i = 0; i < value.length; i += 1) {
    hash = (hash * 31 + value.charCodeAt(i)) >>> 0;
  }
  return hash % mod;
};

type TagPillProps = {
  tag: string;
  active: boolean;
  onToggle: (tag: string) => void;
};

const TagPill = ({ tag, active, onToggle }: TagPillProps) => {
  const style = TAG_STYLES[hashToIndex(tag.toLowerCase(), TAG_STYLES.length)];

  return (
    <button
      type="button"
      onClick={() => onToggle(tag)}
      className={`rounded-full border px-3 py-1 text-xs font-semibold shadow-sm ring-1 transition focus:outline-none focus-visible:ring-2 focus-visible:ring-sky-500/50 ${
        active
          ? `${style.active} shadow-lg`
          : `${style.base} hover:-translate-y-px hover:shadow-md`
      }`}
      aria-pressed={active}
    >
      {tag}
    </button>
  );
};

export default TagPill;
