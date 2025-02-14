const API_BASE = 'http://localhost:8080/api';

export async function register(username: string, password: string, email: string) {
  const res = await fetch(`${API_BASE}/register`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ username, password, email }),
  });
  
  if (!res.ok) {
    const error = await res.json();
    throw new Error(error.error || '注册失败');
  }
  
  return res.json();
}

export async function login(username: string, password: string) {
  const res = await fetch(`${API_BASE}/login`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ username, password }),
  });

  if (!res.ok) {
    const error = await res.json();
    throw new Error(error.error || '登录失败');
  }

  return res.json();
}

export async function getProfile(token: string) {
  const res = await fetch(`${API_BASE}/profile`, {
    headers: {
      'Authorization': `Bearer ${token}`,
    },
  });

  if (!res.ok) {
    const error = await res.json();
    throw new Error(error.error || '获取个人资料失败');
  }

  return res.json();
} 