

export function useEnvironmentVariable(key: string, dev: string) {
  const envVar = import.meta.env[`VITE_${key}`];
  console.log("We did an ", envVar);
  return envVar ? envVar : dev;
}