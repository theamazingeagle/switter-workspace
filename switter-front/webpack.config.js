const path = require('path');
const VueLoaderPlugin = require('vue-loader/lib/plugin');
const HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = {
    mode: 'development',
    entry: './src/index.js',
    //watch: true,
    devServer: {
        contentBase: [
            path.join(__dirname, 'dist'),
            path.join(__dirname, 'assets'),
          ],
        compress: true,
        hot: true,
        disableHostCheck: true,
        host: '0.0.0.0',
        port: 8080
    },
    module: {
        rules: [
            {
                test: /\.vue$/,
                loader: 'vue-loader',
            },
            {
                test: /\.js$/,
                use: {
                    loader: "babel-loader"
                }
            },
            {
                test: /.css$/,
                use: [
                  'style-loader',
                  'css-loader',
                ]
            },
            {
                test: /\.s[ac]ss$/i,
                use: [
                  'style-loader',
                  'sass-loader',
                ]
            },
            {
                test: /\.(png|svg|jpg|jpeg|gif)$/i,
                use: [
                    {
                      loader: 'file-loader',
                      options: {
                        name: '[name].[contenthash].[ext]',
                        esModule: false, // <- here
                      }
                    },
                ],
            },         
        ],
    },    
    plugins: [
        new HtmlWebpackPlugin({
            template: "./src/index.html",
            title: 'Development',
        }),
        new VueLoaderPlugin()
    ],
}
