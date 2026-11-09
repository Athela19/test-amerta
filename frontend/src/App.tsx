import { useEffect, useState } from 'react'
import { me, signOut } from './api'
import type { User } from './api'
import LoginPage from './pages/LoginPage'
import HomePage from './pages/HomePage'

export default function App() {
  const [user, setUser] = useState<User | null>(null)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    let active = true

    me()
      .then((current) => {
        if (active) setUser(current)
      })
      .catch(() => {
        if (active) setUser(null)
      })
      .finally(() => {
        if (active) setLoading(false)
      })

    return () => {
      active = false
    }
  }, [])

  async function logout() {
    try {
      await signOut()
    } finally {
      setUser(null)
    }
  }

  if (loading) {
    return (
      <main className="auth-page">
        <section className="auth-card" aria-busy="true" aria-label="Memeriksa sesi">
          <div className="skeleton logo" />
          <div className="skeleton title" />
          <div className="skeleton line" />
          <div className="skeleton input" />
          <div className="skeleton input" />
          <div className="skeleton button" />
        </section>
      </main>
    )
  }

  if (!user) {
    return <LoginPage onLogin={setUser} />
  }

  return <HomePage user={user} onLogout={logout} />
}
