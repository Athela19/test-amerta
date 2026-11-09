export type Role = 'admin' | 'user'

export type User = {
  id: number
  email: string
  role: Role
}

type Auth = {
  user: User
}

const base = import.meta.env.VITE_API_URL ?? '/api'
let pending: Promise<boolean> | null = null

async function refresh(): Promise<boolean> {
  pending ??= fetch(`${base}/auth/refresh`, {
    method: 'POST',
    credentials: 'include',
  })
    .then((res) => res.ok)
    .catch(() => false)
    .finally(() => {
      pending = null
    })

  return pending
}

async function request<T>(path: string, init: RequestInit = {}, retry = true): Promise<T> {
  const headers = new Headers(init.headers)
  headers.set('Content-Type', 'application/json')

  let res: Response
  try {
    res = await fetch(`${base}${path}`, {
      ...init,
      headers,
      credentials: 'include',
    })
  } catch {
    throw new Error('Tidak dapat terhubung ke server. Periksa koneksi Anda.')
  }

  const authless = path === '/auth/login' || path === '/auth/refresh'
  if (res.status === 401 && retry && !authless && (await refresh())) {
    return request<T>(path, init, false)
  }

  if (!res.ok) {
    const body = (await res.json().catch(() => null)) as { error?: string } | null
    throw new Error(body?.error ?? 'Terjadi kesalahan. Coba lagi.')
  }

  return res.status === 204 ? (undefined as T) : ((await res.json()) as T)
}

export async function login(email: string, password: string): Promise<User> {
  const { user } = await request<Auth>(
    '/auth/login',
    {
      method: 'POST',
      body: JSON.stringify({ email, password }),
    },
    false,
  )
  return user
}

export async function me(): Promise<User> {
  const { user } = await request<Auth>('/auth/me')
  return user
}

export async function signOut(): Promise<void> {
  await request<void>('/auth/logout', { method: 'POST' }, false)
}
