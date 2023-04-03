import config from './config';

export async function fetchResult(url) {
  const response = await fetch(`${config.apiEndpoint}?url=${encodeURIComponent(url)}`);
  const data = await response.json();
  return data;
}
