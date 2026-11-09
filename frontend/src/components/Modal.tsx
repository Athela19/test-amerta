import { useEffect } from 'react'
import type { ReactNode } from 'react'

type Props = {
  title: string
  children: ReactNode
  onClose: () => void
}

export default function Modal({ title, children, onClose }: Props) {
  useEffect(() => {
    const onKey = (e: KeyboardEvent) => {
      if (e.key === 'Escape') onClose()
    }
    window.addEventListener('keydown', onKey)
    return () => window.removeEventListener('keydown', onKey)
  }, [onClose])

  return (
    <div className="modal-root">
      <button className="modal-backdrop" type="button" aria-label="Tutup" onClick={onClose} />
      <div className="modal-card" role="dialog" aria-modal="true">
        <div className="modal-head">
          <h2>{title}</h2>
          <button className="ghost" type="button" onClick={onClose} aria-label="Tutup">
            ✕
          </button>
        </div>
        {children}
      </div>
    </div>
  )
}
