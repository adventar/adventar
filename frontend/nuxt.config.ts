import NuxtConfiguration from "@nuxt/config";

const lang = "ja";
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
    API_BASE_URL: process.env.API_BASE_URL || "http://localhost:8000",
    FIREBASE_API_KEY: process.env.FIREBASE_API_KEY || "AIzaSyCDGROd2Z-2rWjHl-cDepyGnuQ975cpqQo",
    FIREBASE_AUTH_DOMAIN: process.env.FIREBASE_AUTH_DOMAIN || "api-project-837626752936.firebaseapp.com",
    FIREBASE_PROJECT_ID: process.env.FIREBASE_PROJECT_ID || "api-project-837626752936",
    CURRENT_DATE: process.env.CURRENT_DATE || ""
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
      { property: "og:site_name", content: "Adventar" },
      { name: "apple-mobile-web-app-status-bar-style", content: "white" },
      { name: "mobile-web-app-capable", content: "yes" }
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
    lang
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
  modules: [
    "@nuxtjs/pwa",
    "nuxt-fontawesome",
    "@nuxtjs/style-resources",
    ["@nuxtjs/google-analytics", { id: "UA-1474271-8" }]
  ],

  pwa: {
    manifest: {
      name: defaultTitle,
      short_name: defaultTitle,
      description: defaultDescription,
      lang,
      background_color: "#ffffff",
      theme_color: "#ffffff",
      icons: [
        { src: "/icon512.png", sizes: "512x512", type: "image/png" },
        { src: "/icon192.png", sizes: "192x192", type: "image/png" }
      ]
    }
  },

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
          "faLink",
          "faCircleNotch"
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
    }
  }
};

export default config;
