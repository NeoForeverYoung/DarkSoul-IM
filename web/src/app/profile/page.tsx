'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { useAuth } from '@/components/AuthProvider';
import { getProfile } from '@/lib/api';
import { User } from '@/types';

export default function Profile() {
  const router = useRouter();
  const { token, user: authUser } = useAuth();
  const [user, setUser] = useState<User | null>(null);
  const [error, setError] = useState('');

  useEffect(() => {
    if (!token) {
      router.push('/login');
      return;
    }

    getProfile(token)
      .then(response => setUser(response.user))
      .catch(err => setError(err.message));
  }, [token, router]);

  if (!token || !authUser) {
    return null;
  }

  if (error) {
    return <div className="text-red-500 text-center mt-8">{error}</div>;
  }

  return (
    <div className="max-w-2xl mx-auto mt-8 p-6 bg-white rounded-lg shadow">
      <h1 className="text-2xl font-bold mb-6">个人资料</h1>
      {user && (
        <div className="space-y-4">
          <div>
            <label className="font-medium">用户名：</label>
            <span>{user.username}</span>
          </div>
          <div>
            <label className="font-medium">邮箱：</label>
            <span>{user.email}</span>
          </div>
          <div>
            <label className="font-medium">注册时间：</label>
            <span>{new Date(user.created_at!).toLocaleString()}</span>
          </div>
        </div>
      )}
    </div>
  );
} 