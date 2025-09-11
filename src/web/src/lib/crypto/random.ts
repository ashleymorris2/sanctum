export function shortId(len = 9) {
  return Array.from(crypto.getRandomValues(new Uint8Array(len)))
    .map(x => x.toString(36))
    .join('')
    .slice(0, len);
}