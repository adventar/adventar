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
    ecmaFeatures: { legacyDecorators: true },
    extraFileExtensions: [".vue"]
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
    "no-console": ["error", { allow: ["error", "warn"] }],
    "no-unused-vars": "off",
    "@typescript-eslint/no-unused-vars": "error",
    // https://github.com/eslint/eslint/issues/11464
    "no-useless-constructor": "off",
    "@typescript-eslint/no-useless-constructor": "error"
  }
};
