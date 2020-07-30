const { CleanWebpackPlugin } = require("clean-webpack-plugin");
const HtmlWebpackPlugin = require("html-webpack-plugin");
const path = require("path");

module.exports = {
  entry: "./src/index.tsx",
  plugins: [
    new CleanWebpackPlugin({
      cleanAfterEveryBuildPatterns: ["public/build"],
    }),
    new HtmlWebpackPlugin({
      template: "src/index.html",
    }),
  ],
  output: {
    path: __dirname + "/public",
    filename: "build/[name].[contenthash].js",
  },
  resolve: {
    extensions: [".ts", ".tsx", ".js"],
    alias: {
      "~": path.resolve(__dirname, "src"),
    },
  },
  module: {
    rules: [{ test: /\.tsx?$/, loader: "ts-loader" }],
  },
  devServer: {
    historyApiFallback: true,
    port: 3001,
  },
};
