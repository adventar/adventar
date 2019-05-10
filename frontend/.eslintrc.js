module.exports = {
  root: true,
  env: {
    browser: true,
    node: true
  },
  parserOptions: {
    parser: "@typescript-eslint/parser",
    sourceType: "module",
    project: "./tsconfig.json",
    ecmaFeatures: { legacyDecorators: true }
  },
  extends: [
    "@nuxtjs",
    "plugin:nuxt/recommended",
    "plugin:prettier/recommended",
    "prettier/vue",
    "prettier/@typescript-eslint"
  ],
  plugins: ["prettier", "@typescript-eslint"],
  // add your custom rules here
  rules: {
    "no-unused-vars": "off",
    "@typescript-eslint/no-unused-vars": "error"
  }
};
