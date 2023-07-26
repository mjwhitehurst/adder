function addHttp(url = '') {
  if (!url.startsWith('http')) {
    return `http://${url}`;
  }
  return url;
}

module.exports = { addHttp };