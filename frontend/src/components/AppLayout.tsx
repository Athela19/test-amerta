import { useState } from 'react'
import { NavLink, Outlet } from 'react-router-dom'
import type { User } from '../api'
import BrandLogo from './BrandLogo'

type Props = {
  user: User
  onLogout: () => Promise<void>
}

export default function AppLayout({ user, onLogout }: Props) {
  const [loading, setLoading] = useState(false)
  const [open, setOpen] = useState(false)
  const admin = user.role === 'admin'
  const role = admin ? 'Administrator' : 'Pengguna'

  async function logout() {
    setLoading(true)
    try {
      await onLogout()
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="app-shell">
      {open && (
        <button className="sidebar-backdrop" type="button" aria-label="Tutup menu" onClick={() => setOpen(false)} />
      )}

      <aside className={open ? 'sidebar open' : 'sidebar'}>
        <div className="sidebar-brand">
          <BrandLogo />
        </div>
        <nav className="sidebar-nav" aria-label="Navigasi utama">
          <NavLink to="/departments" onClick={() => setOpen(false)}>
            Departemen
          </NavLink>
          {admin && (
            <NavLink to="/users" onClick={() => setOpen(false)}>
              Pengguna
            </NavLink>
          )}
        </nav>
      </aside>

      <div className="app-main">
        <header className="topbar">
          <button className="menu-toggle secondary" type="button" onClick={() => setOpen(true)}>
            Menu
          </button>
          <div className="topbar-user">
            <div>
              <strong>{user.username}</strong>
              <span>{role}</span>
            </div>
            <button className="secondary" type="button" onClick={logout} disabled={loading}>
              {loading ? 'Keluar...' : 'Keluar'}
            </button>
          </div>
        </header>
        <Outlet context={{ user }} />
      </div>
    </div>
  )
}
