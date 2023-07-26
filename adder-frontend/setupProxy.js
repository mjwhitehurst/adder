const { createProxyMiddleware } = require('http-proxy-middleware');

module.exports = function(app) {
  app.use(
    '/proxy',
    createProxyMiddleware({
      target: 'http://mawhit1:8080',
      changeOrigin: true,
      pathRewrite: {
        '^/proxy': '',
      },
    })
  );
};