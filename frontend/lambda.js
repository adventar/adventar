const { Nuxt } = require("nuxt");
const serverless = require("serverless-http");
const config = require("./nuxt.config").default;

const nuxt = new Nuxt({ ...config, dev: false });

module.exports.handler = serverless(async (req, res) => {
  await nuxt.ready();
  nuxt.server.app(req, res);
});
