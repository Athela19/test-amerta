import { useEffect, useState } from 'react'
import { BrowserRouter, Navigate, Route, Routes } from 'react-router-dom'
import { me, signOut } from './api'
import type { User } from './api'
import AppLayout from './components/AppLayout'
import LoginPage from './pages/LoginPage'
import DepartmentsPage from './pages/DepartmentsPage'
import UsersPage from './pages/UsersPage'

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

  return (
    <BrowserRouter>
      <Routes>
        <Route element={<AppLayout user={user} onLogout={logout} />}>
          <Route index element={<Navigate to="/departments" replace />} />
          <Route path="departments" element={<DepartmentsPage />} />
          <Route path="users" element={<UsersPage />} />
          <Route path="*" element={<Navigate to="/departments" replace />} />
        </Route>
      </Routes>
    </BrowserRouter>
  )
}
