import { useState } from 'react'
import type { User } from '../api'
import BrandLogo from '../components/BrandLogo'

type Props = {
  user: User
  onLogout: () => Promise<void>
}

export default function HomePage({ user, onLogout }: Props) {
  const [loading, setLoading] = useState(false)
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
    <>
      <header className="topbar">
        <BrandLogo className="topbar-brand" />
        <div className="topbar-user">
          <div>
            <strong>{user.email}</strong>
            <span>{role}</span>
          </div>
          <button className="secondary" type="button" onClick={logout} disabled={loading}>
            {loading ? 'Keluar...' : 'Keluar'}
          </button>
        </div>
      </header>

      <main className="page">
        <header className="page-head">
          <p className="eyebrow">Beranda</p>
          <h1>Selamat datang</h1>
          <p>Anda masuk sebagai {role.toLowerCase()}.</p>
        </header>

        <section className="card">
          <div className="card-head">
            <h2>Profil akun</h2>
            <span className={admin ? 'badge admin' : 'badge user'}>{role}</span>
          </div>

          <dl className="profile">
            <div>
              <dt>Email</dt>
              <dd>{user.email}</dd>
            </div>
            <div>
              <dt>Peran</dt>
              <dd>{role}</dd>
            </div>
          </dl>

          <div className={admin ? 'panel admin' : 'panel'}>
            <strong>{admin ? 'Akses administrator' : 'Akses pengguna'}</strong>
            <p>
              {admin
                ? 'Akun ini dapat mengelola data dan pengguna aplikasi.'
                : 'Akun ini dapat melihat data yang diizinkan.'}
            </p>
          </div>
        </section>
      </main>
    </>
  )
}
