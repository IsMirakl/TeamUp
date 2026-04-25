export interface User {
  email: string;
  role: UserRole;
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

export interface ProfileData {
  name: string;
  email: string;
  avatar?: string;
  role: UserRole;
  subscriptionPlan: SubscriptionPlan;
}

export interface ProfileResponse {
  user_id: string;
  name: string;
  email: string;
  avatar?: string | null;
  role: UserRole;
  subscriptionPlan: SubscriptionPlan;
}

export type UserRole = 'user' | 'admin' | 'team_lead';
export type SubscriptionPlan = 'Free' | 'Pro' | 'Enterprise';
