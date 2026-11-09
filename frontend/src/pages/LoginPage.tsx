import { useState } from 'react'
import type { FormEvent } from 'react'
import { login } from '../api'
import type { User } from '../api'
import BrandLogo from '../components/BrandLogo'

type Props = {
  onLogin: (user: User) => void
}

export default function LoginPage({ onLogin }: Props) {
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState<string | null>(null)
  const [loading, setLoading] = useState(false)

  async function submit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault()
    setError(null)
    setLoading(true)

    try {
      onLogin(await login(username.trim(), password))
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
            <label htmlFor="username">
              Username <span aria-hidden="true">*</span>
            </label>
            <input
              id="username"
              type="text"
              autoComplete="username"
              placeholder="Masukkan username"
              required
              value={username}
              onChange={(event) => setUsername(event.target.value)}
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
          Belum punya akses? <strong>Hubungi administrator</strong>
        </p>
      </section>
    </main>
  )
}
