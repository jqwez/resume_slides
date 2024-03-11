

export function useEnvironmentVariable(key: string, dev: string) {
  const envVar = import.meta.env[key];
  return envVar ? envVar : dev;
}