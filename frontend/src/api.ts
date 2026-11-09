export type Role = 'admin' | 'user'

export type User = {
  id: number
  username: string
  role: Role
}

export type NewUser = {
  username: string
  password: string
  role: Role
}

type UserResponse = {
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

export async function login(username: string, password: string): Promise<User> {
  const { user } = await request<UserResponse>(
    '/auth/login',
    {
      method: 'POST',
      body: JSON.stringify({ username, password }),
    },
    false,
  )
  return user
}

export async function me(): Promise<User> {
  const { user } = await request<UserResponse>('/auth/me')
  return user
}

export async function signOut(): Promise<void> {
  await request<void>('/auth/logout', { method: 'POST' }, false)
}

export async function listUsers(): Promise<User[]> {
  const { users } = await request<{ users: User[] }>('/users')
  return users
}

export async function addUser(input: NewUser): Promise<User> {
  const { user } = await request<UserResponse>('/users', {
    method: 'POST',
    body: JSON.stringify(input),
  })
  return user
}

export async function updateRole(id: number, role: Role): Promise<User> {
  const { user } = await request<UserResponse>(`/users/${id}`, {
    method: 'PATCH',
    body: JSON.stringify({ role }),
  })
  return user
}

export async function removeUser(id: number): Promise<void> {
  await request<void>(`/users/${id}`, { method: 'DELETE' })
}

export type Employee = {
  id: number
  name: string
}

export type Department = {
  id: number
  name: string
  employee_count: number
  employees: Employee[]
}

export async function listDepts(): Promise<Department[]> {
  const { departments } = await request<{ departments: Department[] }>('/departments')
  return departments
}

export async function addDept(name: string): Promise<Department> {
  const { department } = await request<{ department: Department }>('/departments', {
    method: 'POST',
    body: JSON.stringify({ name }),
  })
  return department
}

export async function updateDept(id: number, name: string): Promise<Department> {
  const { department } = await request<{ department: Department }>(`/departments/${id}`, {
    method: 'PATCH',
    body: JSON.stringify({ name }),
  })
  return department
}

export async function removeDept(id: number): Promise<void> {
  await request<void>(`/departments/${id}`, { method: 'DELETE' })
}
