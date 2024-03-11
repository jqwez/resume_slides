export function useEnvironmentVariable(key: string) {
  const envVar = import.meta.env[`VITE_${key}`];
  return envVar
}