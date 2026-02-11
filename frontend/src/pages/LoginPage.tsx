import LoginForm from '../components/forms/LoginForm';

const LoginPage = () => {
  return (
    <section className="flex min-h-screen items-center justify-center bg-white">
      <div className="flex w-full max-w-md flex-col items-center justify-center rounded-3xl border-2 border-gray-200 bg-white p-8 shadow-2xl">
        <h1 className="mb-8 text-3xl font-bold text-gray-900">Логин</h1>
        <LoginForm />
      </div>
    </section>
  );
};

export default LoginPage;
