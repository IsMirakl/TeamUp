export interface User {
  email: string;
  role: UserRoles;
  name: string;
  avatarUrl?: string;
  subscriptionPlan: SubscriptionPlan;
}

export interface LoginData {
  email: string;
  password: string;
}

export interface RegisterData {
  name: string;
  email: string;
  password: string;
  confirmPassword: string;
  avatarUrl?: string;
}

export interface AuthResponse {
  user: User;
  accessToken: string;
}

enum UserRoles {
  USER,
  ADMIN,
}

enum SubscriptionPlan {
  FREE,
  PRO,
  ENTERPRISE,
}
