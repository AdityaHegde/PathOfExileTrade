const path = require("path");
const webpack = require("webpack");
const HtmlWebpackPlugin = require("html-webpack-plugin");
const { CleanWebpackPlugin } = require("clean-webpack-plugin");

module.exports = {
  mode: "development",
  devtool: "source-map",
  entry: [
    "./src/typescript/index.tsx",
  ],
  resolve: {
    extensions: ["*", ".ts", ".tsx", ".js"],
    alias: {
      "@model": path.resolve(__dirname, "src/typescript/model"),
      "@component": path.resolve(__dirname, "src/typescript/component"),
      "@typescript": path.resolve(__dirname, "src/typescript"),
    },
  },

  module: {
    rules: [{
      test: /src\/.*\.tsx?$/,
      loader: "ts-loader",
      options: {
        configFile: "tsconfig-webpack.json",
      },
    }, {
      test: /\.(css|less)$/,
      use: [{
        loader: "style-loader",
      }, {
        loader: "css-loader",
      }, {
        loader: "less-loader",
        options: {
          lessOptions: {
            javascriptEnabled: true,
          },
        },
      }],
    }],
  },

  watch: true,

  plugins: [
    new CleanWebpackPlugin({
      cleanStaleWebpackAssets: false,
    }),
    new HtmlWebpackPlugin({
      title: "Game Engine UI",
    }),
  ],

  optimization: {
    minimize: false,
    splitChunks: {
      chunks: "all",
    },
  },

  output: {
    filename: "[name].bundle.js",
    path: path.resolve(__dirname, "public"),
    publicPath: "/",
  },
};
