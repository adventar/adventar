import NuxtConfiguration from "@nuxt/config";

const config: NuxtConfiguration = {
  mode: process.env.BUILD_MODE === "spa" ? "spa" : "universal",

  env: {
    apiBaseUrl: process.env.API_BASE_URL || "http://localhost:8000",
    firebaseApiKey: process.env.FIREBASE_API_KEY || "AIzaSyCDGROd2Z-2rWjHl-cDepyGnuQ975cpqQo",
    firebaseAuthDomain: process.env.FIREBASE_AUTH_DOMAIN || "api-project-837626752936.firebaseapp.com",
    firebaseProjectId: process.env.FIREBASE_PROJECT_ID || "api-project-837626752936"
  },

  /*
   ** Headers of the page
   */
  head: {
    title: "Adventar",
    meta: [{ charset: "utf-8" }, { name: "viewport", content: "width=device-width, initial-scale=1" }],
    link: [{ rel: "icon", type: "image/x-icon", href: "/favicon.ico" }]
  },

  router: {
    middleware: "signin"
  },

  /*
   ** Customize the progress-bar color
   */
  loading: { color: "#999" },

  /*
   ** Global CSS
   */
  css: [],

  /*
   ** Plugins to load before mounting the App
   */
  plugins: [{ src: "~/plugins/auth", ssr: false }],

  /*
   ** Nuxt.js modules
   */
  modules: ["@nuxtjs/pwa", "nuxt-fontawesome"],

  fontawesome: {
    imports: [
      {
        set: "@fortawesome/free-solid-svg-icons",
        icons: ["fas"]
      },
      {
        set: "@fortawesome/free-brands-svg-icons",
        icons: ["fab"]
      }
    ]
  },

  /*
   ** Build configuration
   */
  build: {
    babel: {
      ignore: [/\/lib\/grpc\//]
    },
    /*
     ** You can extend webpack config here
     */
    extend(config, ctx) {
      if (ctx.isClient) config.devtool = "#source-map";

      // if (config.module) {
      //   config.module.rules[2] = {};
      // }

      // Run ESLint on save
      // if (ctx.isDev && ctx.isClient && config.module) {
      //   config.module.rules.push({
      //     enforce: "pre",
      //     test: /\.(js|vue)$/,
      //     loader: "eslint-loader",
      //     exclude: /(node_modules)/
      //   });
      // }
    }
  }
};

export default config;
