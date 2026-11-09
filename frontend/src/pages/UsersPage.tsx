import { useEffect, useState } from 'react'
import type { FormEvent } from 'react'
import { Navigate, useOutletContext } from 'react-router-dom'
import { addUser, listUsers, removeUser, updateRole } from '../api'
import type { Role, User } from '../api'
import Modal from '../components/Modal'

type Ctx = { user: User }

export default function UsersPage() {
  const { user: me } = useOutletContext<Ctx>()
  const [rows, setRows] = useState<User[]>([])
  const [q, setQ] = useState('')
  const [search, setSearch] = useState('')
  const [loading, setLoading] = useState(true)
  const [busy, setBusy] = useState(false)
  const [error, setError] = useState<string | null>(null)
  const [edit, setEdit] = useState<User | 'new' | null>(null)
  const [del, setDel] = useState<User | null>(null)
  const [menu, setMenu] = useState<number | null>(null)
  const [ok, setOk] = useState<string | null>(null)
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const [role, setRole] = useState<Role>('user')
  const [formErr, setFormErr] = useState<string | null>(null)

  async function load() {
    setLoading(true)
    setError(null)
    try {
      setRows(await listUsers())
    } catch (e) {
      setError(e instanceof Error ? e.message : 'Gagal memuat pengguna.')
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    if (me.role === 'admin') void load()
  }, [me.role])

  useEffect(() => {
    if (menu === null) return
    const close = () => setMenu(null)
    window.addEventListener('click', close)
    return () => window.removeEventListener('click', close)
  }, [menu])

  if (me.role !== 'admin') {
    return <Navigate to="/departments" replace />
  }

  const filtered = (() => {
    const key = search.trim().toLowerCase()
    if (!key) return rows
    return rows.filter(
      (u) =>
        u.username.toLowerCase().includes(key) ||
        u.role.includes(key) ||
        String(u.id).includes(key),
    )
  })()

  function openNew() {
    setUsername('')
    setPassword('')
    setRole('user')
    setFormErr(null)
    setEdit('new')
  }

  function openEdit(u: User) {
    setRole(u.role)
    setFormErr(null)
    setMenu(null)
    setEdit(u)
  }

  async function save(e: FormEvent) {
    e.preventDefault()
    setBusy(true)
    setFormErr(null)
    try {
      if (edit && edit !== 'new') {
        await updateRole(edit.id, role)
        setOk('Peran pengguna berhasil diperbarui.')
      } else {
        await addUser({ username: username.trim(), password, role })
        setOk('Pengguna berhasil ditambahkan.')
      }
      setEdit(null)
      await load()
    } catch (err) {
      setFormErr(err instanceof Error ? err.message : 'Gagal menyimpan pengguna.')
    } finally {
      setBusy(false)
    }
  }

  async function remove() {
    if (!del) return
    setBusy(true)
    try {
      await removeUser(del.id)
      setOk(`Pengguna ${del.username} berhasil dihapus.`)
      setDel(null)
      await load()
    } catch (err) {
      setDel(null)
      setOk(err instanceof Error ? err.message : 'Gagal menghapus pengguna.')
    } finally {
      setBusy(false)
    }
  }

  return (
    <main className="page">
      <header className="page-head row">
        <div>
          <h1>Pengguna</h1>
          <p>Kelola akun aplikasi, peran, dan akses pengguna.</p>
        </div>
        <button className="primary" type="button" onClick={openNew}>
          Tambah pengguna
        </button>
      </header>

      <section className="panel-card">
        <form
          className="toolbar"
          onSubmit={(e) => {
            e.preventDefault()
            setSearch(q)
          }}
        >
          <div className="search-field">
            <input
              type="search"
              placeholder="Cari username atau peran"
              value={q}
              onChange={(e) => setQ(e.target.value)}
              aria-label="Cari pengguna"
            />
            <button className="secondary" type="submit">
              Cari
            </button>
          </div>
          {search && (
            <button
              className="ghost"
              type="button"
              onClick={() => {
                setQ('')
                setSearch('')
              }}
            >
              Reset
            </button>
          )}
        </form>

        {error && (
          <div className="alert" role="alert">
            {error}
            <button className="ghost" type="button" onClick={() => void load()}>
              Coba lagi
            </button>
          </div>
        )}

        <div className="table-wrap">
          <table className="data-table">
            <thead>
              <tr>
                <th style={{ width: '10%' }}>ID</th>
                <th>Username</th>
                <th style={{ width: '22%' }}>Peran</th>
                <th style={{ width: '10%' }} className="align-right">
                  Aksi
                </th>
              </tr>
            </thead>
            <tbody>
              {loading ? (
                <tr>
                  <td colSpan={4} className="table-empty">
                    Memuat pengguna...
                  </td>
                </tr>
              ) : filtered.length === 0 ? (
                <tr>
                  <td colSpan={4} className="table-empty">
                    {search ? 'Tidak ada pengguna yang cocok.' : 'Belum ada pengguna.'}
                  </td>
                </tr>
              ) : (
                filtered.map((u) => {
                  const self = u.id === me.id
                  return (
                    <tr key={u.id}>
                      <td>{u.id}</td>
                      <td>
                        <strong>{u.username}</strong>
                        {self && <span className="inline-note"> · Anda</span>}
                      </td>
                      <td>
                        <span className={u.role === 'admin' ? 'badge admin' : 'badge user'}>
                          {u.role === 'admin' ? 'Administrator' : 'Pengguna'}
                        </span>
                      </td>
                      <td className="align-right">
                        <div className="menu">
                          <button
                            className="ghost menu-btn"
                            type="button"
                            aria-label={`Aksi ${u.username}`}
                            aria-expanded={menu === u.id}
                            disabled={self}
                            onClick={(e) => {
                              e.stopPropagation()
                              setMenu(menu === u.id ? null : u.id)
                            }}
                          >
                            ⋯
                          </button>
                          {menu === u.id && (
                            <div className="menu-pop" role="menu" onClick={(e) => e.stopPropagation()}>
                              <button type="button" role="menuitem" onClick={() => openEdit(u)}>
                                Ubah
                              </button>
                              <button
                                type="button"
                                role="menuitem"
                                className="danger-text"
                                onClick={() => {
                                  setMenu(null)
                                  setDel(u)
                                }}
                              >
                                Hapus
                              </button>
                            </div>
                          )}
                        </div>
                      </td>
                    </tr>
                  )
                })
              )}
            </tbody>
          </table>
        </div>
      </section>

      {edit && (
        <Modal
          title={edit === 'new' ? 'Tambah pengguna' : 'Ubah peran pengguna'}
          onClose={() => !busy && setEdit(null)}
        >
          {formErr && (
            <div className="alert" role="alert">
              {formErr}
            </div>
          )}
          <form className="modal-form" onSubmit={(e) => void save(e)}>
            {edit === 'new' ? (
              <>
                <div className="field">
                  <label htmlFor="new-username">
                    Username <span aria-hidden="true">*</span>
                  </label>
                  <input
                    id="new-username"
                    value={username}
                    required
                    minLength={3}
                    maxLength={32}
                    pattern="[a-zA-Z0-9._-]{3,32}"
                    autoComplete="off"
                    disabled={busy}
                    onChange={(e) => setUsername(e.target.value)}
                  />
                </div>
                <div className="field">
                  <label htmlFor="new-password">
                    Kata sandi <span aria-hidden="true">*</span>
                  </label>
                  <input
                    id="new-password"
                    type="password"
                    value={password}
                    required
                    minLength={6}
                    maxLength={72}
                    autoComplete="new-password"
                    disabled={busy}
                    onChange={(e) => setPassword(e.target.value)}
                  />
                </div>
              </>
            ) : (
              <p className="modal-message">
                Ubah peran untuk <strong>{edit.username}</strong>.
              </p>
            )}
            <div className="field">
              <label htmlFor="user-role">Peran</label>
              <select
                id="user-role"
                value={role}
                disabled={busy}
                onChange={(e) => setRole(e.target.value as Role)}
              >
                <option value="user">Pengguna</option>
                <option value="admin">Administrator</option>
              </select>
            </div>
            <div className="modal-actions">
              <button className="secondary" type="button" disabled={busy} onClick={() => setEdit(null)}>
                Batal
              </button>
              <button className="primary" type="submit" disabled={busy}>
                {busy ? 'Menyimpan...' : 'Simpan'}
              </button>
            </div>
          </form>
        </Modal>
      )}

      {del && (
        <Modal title="Hapus pengguna?" onClose={() => !busy && setDel(null)}>
          <p className="modal-message">
            Pengguna <strong>{del.username}</strong> akan dihapus permanen.
          </p>
          <div className="modal-actions">
            <button className="secondary" type="button" disabled={busy} onClick={() => setDel(null)}>
              Batal
            </button>
            <button className="danger" type="button" disabled={busy} onClick={() => void remove()}>
              {busy ? 'Memproses...' : 'Hapus'}
            </button>
          </div>
        </Modal>
      )}

      {ok && (
        <Modal title="Info" onClose={() => setOk(null)}>
          <p className="modal-message">{ok}</p>
          <div className="modal-actions">
            <button className="primary" type="button" onClick={() => setOk(null)}>
              Tutup
            </button>
          </div>
        </Modal>
      )}
    </main>
  )
}
