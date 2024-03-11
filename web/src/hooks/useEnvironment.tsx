

export function useEnvironmentVariable(key: string, dev: string) {
  const envVar = import.meta.env[`VITE_${key}`];
  return envVar ? envVar : dev;
}