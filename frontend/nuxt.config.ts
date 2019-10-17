import NuxtConfiguration from "@nuxt/config";

const defaultTitle = "Adventar";
const defaultDescription =
  "12月1日から25日まで1日に1つ、みんなで記事を投稿するAdvent Calendarの作成や管理をおこなうことができます。";

const config: NuxtConfiguration = {
  dev: process.env.NODE_ENV !== "production",
  buildDir: process.env.BUILD_DIR || ".nuxt",

  server: {
    port: 3333
  },

  env: {
    API_BASE_URL: process.env.NODE_ENV === "development" ? "http://localhost:8000" : "https://api.adventar.org",
    FIREBASE_API_KEY: process.env.FIREBASE_API_KEY || "AIzaSyCDGROd2Z-2rWjHl-cDepyGnuQ975cpqQo",
    FIREBASE_AUTH_DOMAIN: process.env.FIREBASE_AUTH_DOMAIN || "api-project-837626752936.firebaseapp.com",
    FIREBASE_PROJECT_ID: process.env.FIREBASE_PROJECT_ID || "api-project-837626752936"
  },

  /*
   ** Headers of the page
   */
  head: {
    title: defaultTitle,
    meta: [
      { charset: "utf-8" },
      { name: "viewport", content: "width=device-width, initial-scale=1" },
      { hid: "description", name: "description", content: defaultDescription },
      { hid: "og:title", property: "og:title", content: defaultTitle },
      { hid: "og:description", property: "og:description", content: defaultDescription },
      { property: "og:image", content: "https://adventar.org/og_image.png" },
      { property: "og:site_name", content: "Adventar" }
    ],
    link: [
      {
        rel: "icon",
        type: "image/x-icon",
        href: process.env.NODE_ENV === "development" ? "/favicon-dev.ico" : "/favicon.ico"
      }
    ]
  },

  htmlAttrs: {
    lang: "ja"
  },

  router: {
    middleware: "signin",
    extendRoutes(routes, resolve) {
      routes.push({
        name: "404",
        path: "*",
        // @ts-ignore: Remove this line after update nuxt. https://github.com/nuxt/nuxt.js/pull/5841
        component: resolve(__dirname, "pages/404.vue")
      });
    }
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
  modules: ["@nuxtjs/pwa", "nuxt-fontawesome", "@nuxtjs/style-resources"],

  fontawesome: {
    imports: [
      {
        set: "@fortawesome/free-solid-svg-icons",
        icons: [
          "faSearch",
          "faBars",
          "faCalendarPlus",
          "faUser",
          "faCog",
          "faSignOutAlt",
          "faCalendarMinus",
          "faQuestionCircle",
          "faEdit",
          "faTimes",
          "faComment",
          "faLink"
        ]
      },
      {
        set: "@fortawesome/free-brands-svg-icons",
        icons: ["faGoogle", "faGithub", "faTwitter", "faFacebook"]
      },
      {
        set: "@fortawesome/free-regular-svg-icons",
        icons: ["faCalendar"]
      }
    ]
  },

  styleResources: {
    scss: ["./assets/scss/*.scss"]
  },

  /*
   ** Build configuration
   */
  build: {
    babel: {
      ignore: ["./lib/grpc"]
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
