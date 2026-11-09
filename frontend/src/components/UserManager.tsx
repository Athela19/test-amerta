import { useEffect, useState } from 'react'
import type { FormEvent } from 'react'
import { addUser, listUsers, removeUser, updateRole } from '../api'
import type { Role, User } from '../api'

type Props = {
  current: User
}

export default function UserManager({ current }: Props) {
  const [users, setUsers] = useState<User[]>([])
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const [role, setRole] = useState<Role>('user')
  const [loading, setLoading] = useState(true)
  const [busy, setBusy] = useState(false)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    listUsers()
      .then(setUsers)
      .catch((caught) => {
        setError(caught instanceof Error ? caught.message : 'Gagal memuat pengguna.')
      })
      .finally(() => setLoading(false))
  }, [])

  async function submit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault()
    setBusy(true)
    setError(null)

    try {
      const created = await addUser({ username: username.trim(), password, role })
      setUsers((list) => [...list, created])
      setUsername('')
      setPassword('')
      setRole('user')
    } catch (caught) {
      setError(caught instanceof Error ? caught.message : 'Gagal menambah pengguna.')
    } finally {
      setBusy(false)
    }
  }

  async function change(id: number, next: Role) {
    setBusy(true)
    setError(null)

    try {
      const updated = await updateRole(id, next)
      setUsers((list) => list.map((user) => (user.id === id ? updated : user)))
    } catch (caught) {
      setError(caught instanceof Error ? caught.message : 'Gagal mengubah peran.')
    } finally {
      setBusy(false)
    }
  }

  async function remove(id: number, name: string) {
    if (!window.confirm(`Hapus pengguna ${name}?`)) return

    setBusy(true)
    setError(null)

    try {
      await removeUser(id)
      setUsers((list) => list.filter((user) => user.id !== id))
    } catch (caught) {
      setError(caught instanceof Error ? caught.message : 'Gagal menghapus pengguna.')
    } finally {
      setBusy(false)
    }
  }

  return (
    <section className="card">
      <div className="card-head">
        <h2>Manajemen pengguna</h2>
      </div>

      {error && (
        <div className="alert" role="alert">
          {error}
        </div>
      )}

      <form className="user-form" onSubmit={submit}>
        <div className="field">
          <label htmlFor="new-username">Username</label>
          <input
            id="new-username"
            type="text"
            autoComplete="off"
            placeholder="username"
            required
            minLength={3}
            maxLength={32}
            pattern="[a-zA-Z0-9._-]{3,32}"
            value={username}
            onChange={(event) => setUsername(event.target.value)}
            disabled={busy}
          />
        </div>

        <div className="field">
          <label htmlFor="new-password">Kata sandi</label>
          <input
            id="new-password"
            type="password"
            autoComplete="new-password"
            placeholder="Minimal 6 karakter"
            required
            minLength={6}
            maxLength={72}
            value={password}
            onChange={(event) => setPassword(event.target.value)}
            disabled={busy}
          />
        </div>

        <div className="field">
          <label htmlFor="new-role">Peran</label>
          <select
            id="new-role"
            value={role}
            onChange={(event) => setRole(event.target.value as Role)}
            disabled={busy}
          >
            <option value="user">Pengguna</option>
            <option value="admin">Administrator</option>
          </select>
        </div>

        <button className="primary" type="submit" disabled={busy}>
          {busy ? 'Menyimpan...' : 'Tambah'}
        </button>
      </form>

      {loading ? (
        <p className="muted">Memuat pengguna...</p>
      ) : users.length === 0 ? (
        <p className="muted">Belum ada pengguna.</p>
      ) : (
        <div className="user-list">
          {users.map((user) => {
            const self = user.id === current.id

            return (
              <div className="user-row" key={user.id}>
                <div className="user-info">
                  <strong>{user.username}</strong>
                  <span>
                    {user.role === 'admin' ? 'Administrator' : 'Pengguna'}
                    {self && ' · Anda'}
                  </span>
                </div>

                <div className="user-actions">
                  <select
                    value={user.role}
                    onChange={(event) => void change(user.id, event.target.value as Role)}
                    disabled={self || busy}
                    aria-label={`Peran ${user.username}`}
                  >
                    <option value="user">Pengguna</option>
                    <option value="admin">Administrator</option>
                  </select>
                  <button
                    className="danger"
                    type="button"
                    onClick={() => void remove(user.id, user.username)}
                    disabled={self || busy}
                  >
                    Hapus
                  </button>
                </div>
              </div>
            )
          })}
        </div>
      )}
    </section>
  )
}
