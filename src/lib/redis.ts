import Redis from "ioredis";
 
export const KEY_MAIN_PAGE = "main";
 
export function getMainKey(): string {
  return `main`;
}

export default new Redis();