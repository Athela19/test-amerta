type Props = {
  page: number
  limit: number
  total: number
  disabled?: boolean
  onPage: (page: number) => void
}

export default function Pager({ page, limit, total, disabled, onPage }: Props) {
  const pages = Math.max(1, Math.ceil(total / Math.max(limit, 1)))
  const from = total === 0 ? 0 : (page - 1) * limit + 1
  const to = Math.min(page * limit, total)

  return (
    <div className="pager">
      <p>
        Menampilkan {from}–{to} dari {total}
      </p>
      <div className="pager-actions">
        <button
          className="secondary"
          type="button"
          disabled={disabled || page <= 1}
          onClick={() => onPage(page - 1)}
        >
          Sebelumnya
        </button>
        <span>
          Halaman {page} dari {pages}
        </span>
        <button
          className="secondary"
          type="button"
          disabled={disabled || page >= pages}
          onClick={() => onPage(page + 1)}
        >
          Berikutnya
        </button>
      </div>
    </div>
  )
}
