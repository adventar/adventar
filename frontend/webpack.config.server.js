const path = require("path");
const nodeExternals = require("webpack-node-externals");

module.exports = {
  mode: "production",
  entry: {
    server: "./server.ts"
  },
  devtool: "source-map",
  resolve: {
    alias: {
      "~": path.resolve(__dirname)
    },
    extensions: [".js", ".jsx", ".json", ".ts", ".tsx"]
  },
  output: {
    libraryTarget: "commonjs",
    path: __dirname,
    filename: "[name].js"
  },
  target: "node",
  externals: [nodeExternals()],
  module: {
    rules: [
      {
        test: /\.tsx?$/,
        loader: "ts-loader",
        options: {
          compilerOptions: { target: "ES2018" }
        }
      }
    ]
  }
};
