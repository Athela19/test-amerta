import { useState } from 'react'
import type { FormEvent } from 'react'
import { login } from '../api'
import type { User } from '../api'
import BrandLogo from '../components/BrandLogo'

type Props = {
  onLogin: (user: User) => void
}

export default function LoginPage({ onLogin }: Props) {
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState<string | null>(null)
  const [loading, setLoading] = useState(false)

  async function submit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault()
    setError(null)
    setLoading(true)

    try {
      onLogin(await login(email.trim(), password))
    } catch (caught) {
      setError(caught instanceof Error ? caught.message : 'Gagal masuk.')
    } finally {
      setLoading(false)
    }
  }

  return (
    <main className="auth-page">
      <section className="auth-card">
        <BrandLogo />

        <h1>Masuk</h1>
        <p className="subtitle">Masuk untuk mengelola data departemen dan karyawan.</p>

        {error && (
          <div className="alert" role="alert">
            {error}
          </div>
        )}

        <form className="auth-form" onSubmit={submit}>
          <div className="field">
            <label htmlFor="email">
              Email <span aria-hidden="true">*</span>
            </label>
            <input
              id="email"
              type="email"
              autoComplete="email"
              placeholder="nama@amerta.local"
              required
              value={email}
              onChange={(event) => setEmail(event.target.value)}
              disabled={loading}
            />
          </div>

          <div className="field">
            <label htmlFor="password">
              Kata sandi <span aria-hidden="true">*</span>
            </label>
            <input
              id="password"
              type="password"
              autoComplete="current-password"
              placeholder="Masukkan kata sandi"
              required
              value={password}
              onChange={(event) => setPassword(event.target.value)}
              disabled={loading}
            />
          </div>

          <button className="primary" type="submit" disabled={loading}>
            {loading && <span className="spinner" />}
            {loading ? 'Memproses...' : 'Masuk'}
          </button>
        </form>

        <p className="auth-help">
          Belum punya akses? <a href="mailto:admin@amerta.local">Hubungi administrator</a>
        </p>
      </section>
    </main>
  )
}
