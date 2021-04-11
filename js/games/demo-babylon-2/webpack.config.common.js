const path = require('path');

const HtmlWebpackPlugin = require('html-webpack-plugin');
const CopyWebpackPlugin = require("copy-webpack-plugin");

module.exports = {

    entry: path.resolve(__dirname, 'src'),

    resolve: {
        extensions: [
            '.ts',
            '.tsx',
            '.js',
            '.json'
        ]
    },

    module: {
        rules: [
            {
                test: /\.(ts|js)x?$/,
                loader: 'babel-loader',
                exclude: /node_modules/
            },
            {
                test: /\.pug$/,
                use: [
                    "html-loader",
                    "pug-html-loader"
                ]
            },
            {
                test: /\.s[ac]ss$/i,
                use: [
                    'style-loader',
                    'css-loader',
                    'sass-loader'
                ],
            }
        ]
    },

    plugins: [
        new HtmlWebpackPlugin({
            template: "./src/templates/index.pug",
            filename: "./index.html"
        }),
        new CopyWebpackPlugin({
            patterns: [
                {
                    from: 'src/assets',
                    to: 'assets',
                    globOptions: {
                        dot: false,
                        ignore: [
                            '**/*.psd',
                            '**/*.ai',
                            '**/*.blend',
                            '**/*.max',
                            '**/*.fbx',
                            '**/*.obj',
                            '**/*.spp',
                            '**/*.xd',
                            '**/*.afdesign',
                            '**/*.pur'
                        ]
                    }
                },
                {
                    from: 'src/data',
                    to: 'data',
                    globOptions: {
                        dot: false,
                        ignore: []
                    }
                },
            ]
        })
    ],

    performance: {
        hints: false
    },
};
