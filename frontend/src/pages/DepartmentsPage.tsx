import { useEffect, useState } from 'react'
import type { FormEvent } from 'react'
import { addDept, listDepts, pageSize, removeDept, updateDept } from '../api'
import type { Department } from '../api'
import Modal from '../components/Modal'
import Pager from '../components/Pager'

export default function DepartmentsPage() {
  const [rows, setRows] = useState<Department[]>([])
  const [q, setQ] = useState('')
  const [search, setSearch] = useState('')
  const [page, setPage] = useState(1)
  const [total, setTotal] = useState(0)
  const [limit, setLimit] = useState(pageSize)
  const [reload, setReload] = useState(0)
  const [loading, setLoading] = useState(true)
  const [busy, setBusy] = useState(false)
  const [error, setError] = useState<string | null>(null)
  const [edit, setEdit] = useState<Department | 'new' | null>(null)
  const [del, setDel] = useState<Department | null>(null)
  const [emps, setEmps] = useState<Department | null>(null)
  const [menu, setMenu] = useState<number | null>(null)
  const [ok, setOk] = useState<string | null>(null)
  const [name, setName] = useState('')
  const [formErr, setFormErr] = useState<string | null>(null)

  useEffect(() => {
    let cancelled = false

    async function run() {
      setLoading(true)
      setError(null)
      setMenu(null)
      try {
        const res = await listDepts(search, page, pageSize)
        if (cancelled) return
        if (res.items.length === 0 && page > 1) {
          setPage(Math.max(1, Math.ceil(res.total / res.limit)))
          return
        }
        setRows(res.items)
        setTotal(res.total)
        setLimit(res.limit)
      } catch (e) {
        if (!cancelled) setError(e instanceof Error ? e.message : 'Gagal memuat departemen.')
      } finally {
        if (!cancelled) setLoading(false)
      }
    }

    void run()
    return () => {
      cancelled = true
    }
  }, [search, page, reload])

  useEffect(() => {
    if (menu === null) return
    const close = () => setMenu(null)
    window.addEventListener('click', close)
    return () => window.removeEventListener('click', close)
  }, [menu])

  function openNew() {
    setName('')
    setFormErr(null)
    setEdit('new')
  }

  function openEdit(d: Department) {
    setName(d.name)
    setFormErr(null)
    setMenu(null)
    setEdit(d)
  }

  async function save(e: FormEvent) {
    e.preventDefault()
    const value = name.trim().toUpperCase()
    if (!value) {
      setFormErr('Nama departemen wajib diisi.')
      return
    }

    setBusy(true)
    setFormErr(null)
    try {
      if (edit && edit !== 'new') {
        await updateDept(edit.id, value)
        setOk('Departemen berhasil diperbarui.')
      } else {
        await addDept(value)
        setOk('Departemen berhasil ditambahkan.')
      }
      setEdit(null)
      setReload((n) => n + 1)
    } catch (err) {
      setFormErr(err instanceof Error ? err.message : 'Gagal menyimpan departemen.')
    } finally {
      setBusy(false)
    }
  }

  async function remove() {
    if (!del) return
    setBusy(true)
    try {
      await removeDept(del.id)
      setOk(`Departemen ${del.name} berhasil dihapus.`)
      setDel(null)
      setReload((n) => n + 1)
    } catch (err) {
      setDel(null)
      setOk(err instanceof Error ? err.message : 'Gagal menghapus departemen.')
    } finally {
      setBusy(false)
    }
  }

  return (
    <main className="page">
      <header className="page-head row">
        <div>
          <h1>Departemen</h1>
          <p>Kelola struktur departemen dan lihat jumlah karyawannya.</p>
        </div>
        <button className="primary" type="button" onClick={openNew}>
          Tambah departemen
        </button>
      </header>

      <section className="panel-card">
        <form
          className="toolbar"
          onSubmit={(e) => {
            e.preventDefault()
            setPage(1)
            setSearch(q.trim())
          }}
        >
          <div className="search-field">
            <input
              type="search"
              placeholder="Cari nama atau karyawan"
              value={q}
              onChange={(e) => setQ(e.target.value)}
              aria-label="Cari departemen"
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
                setPage(1)
              }}
            >
              Reset
            </button>
          )}
        </form>

        {error && (
          <div className="alert" role="alert">
            {error}
            <button className="ghost" type="button" onClick={() => setReload((n) => n + 1)}>
              Coba lagi
            </button>
          </div>
        )}

        <div className="table-wrap">
          <table className="data-table">
            <thead>
              <tr>
                <th style={{ width: '10%' }}>ID</th>
                <th>Nama</th>
                <th style={{ width: '20%' }}>Karyawan</th>
                <th style={{ width: '10%' }} className="align-right">
                  Aksi
                </th>
              </tr>
            </thead>
            <tbody>
              {loading ? (
                <tr>
                  <td colSpan={4} className="table-empty">
                    Memuat departemen...
                  </td>
                </tr>
              ) : rows.length === 0 ? (
                <tr>
                  <td colSpan={4} className="table-empty">
                    {search ? 'Tidak ada departemen yang cocok.' : 'Belum ada departemen.'}
                  </td>
                </tr>
              ) : (
                rows.map((d) => (
                  <tr key={d.id}>
                    <td>{d.id}</td>
                    <td>
                      <strong>{d.name}</strong>
                    </td>
                    <td>
                      <button className="link-btn" type="button" onClick={() => setEmps(d)}>
                        {d.employee_count} karyawan
                      </button>
                    </td>
                    <td className="align-right">
                      <div className="menu">
                        <button
                          className="ghost menu-btn"
                          type="button"
                          aria-label={`Aksi ${d.name}`}
                          aria-expanded={menu === d.id}
                          onClick={(e) => {
                            e.stopPropagation()
                            setMenu(menu === d.id ? null : d.id)
                          }}
                        >
                          ⋯
                        </button>
                        {menu === d.id && (
                          <div className="menu-pop" role="menu" onClick={(e) => e.stopPropagation()}>
                            <button type="button" role="menuitem" onClick={() => openEdit(d)}>
                              Ubah
                            </button>
                            <button
                              type="button"
                              role="menuitem"
                              className="danger-text"
                              disabled={d.employee_count > 0}
                              title={
                                d.employee_count > 0
                                  ? 'Pindahkan atau hapus karyawan terlebih dahulu'
                                  : undefined
                              }
                              onClick={() => {
                                setMenu(null)
                                setDel(d)
                              }}
                            >
                              Hapus
                            </button>
                          </div>
                        )}
                      </div>
                    </td>
                  </tr>
                ))
              )}
            </tbody>
          </table>
        </div>
        {total > 0 && (
          <Pager page={page} limit={limit} total={total} disabled={loading} onPage={setPage} />
        )}
      </section>

      {edit && (
        <Modal
          title={edit === 'new' ? 'Tambah departemen' : 'Ubah departemen'}
          onClose={() => !busy && setEdit(null)}
        >
          {formErr && (
            <div className="alert" role="alert">
              {formErr}
            </div>
          )}
          <form className="modal-form" onSubmit={(e) => void save(e)}>
            <div className="field">
              <label htmlFor="dept-name">
                Nama departemen <span aria-hidden="true">*</span>
              </label>
              <input
                id="dept-name"
                value={name}
                maxLength={100}
                required
                autoFocus
                disabled={busy}
                style={{ textTransform: 'uppercase' }}
                onChange={(e) => setName(e.target.value.toUpperCase())}
              />
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

      {emps && (
        <Modal title={`Karyawan · ${emps.name}`} onClose={() => setEmps(null)}>
          {emps.employees.length === 0 ? (
            <p className="modal-message">Belum ada karyawan di departemen ini.</p>
          ) : (
            <ul className="emp-list">
              {emps.employees.map((e) => (
                <li key={e.id}>{e.name}</li>
              ))}
            </ul>
          )}
          <div className="modal-actions">
            <button className="primary" type="button" onClick={() => setEmps(null)}>
              Tutup
            </button>
          </div>
        </Modal>
      )}

      {del && (
        <Modal title="Hapus departemen?" onClose={() => !busy && setDel(null)}>
          <p className="modal-message">
            Departemen <strong>{del.name}</strong> akan dihapus permanen.
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
