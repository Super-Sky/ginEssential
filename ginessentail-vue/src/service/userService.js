import request from '@/utils/request';

// 用户注册
const register = ({ name, telephone, password }) => request.post('auth/register', { name, telephone, password });

// 获取用户信息
const info = () => {
  return request.get('auth/info');
};

export default {
  info,
  register,
};
