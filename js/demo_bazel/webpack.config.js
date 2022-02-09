const path = require('path');
const { merge } = require('webpack-merge');
const common = require('./webpack.config.common.js');

const ForkTsCheckerWebpackPlugin = require('fork-ts-checker-webpack-plugin');
const { CleanWebpackPlugin } = require('clean-webpack-plugin');
const UglifyJsPlugin = require('uglifyjs-webpack-plugin');

module.exports = merge(common, {

    mode: 'production',

    plugins: [
        new ForkTsCheckerWebpackPlugin(),
        new CleanWebpackPlugin()
    ],

    optimization: {
        minimizer: [ new UglifyJsPlugin() ],
    // },

    // output: {
    //    filename: 'bundle.js',
    //    path: path.resolve(__dirname, 'dist')
    }
});
