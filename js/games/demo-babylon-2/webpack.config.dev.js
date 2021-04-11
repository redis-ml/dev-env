const { merge } = require('webpack-merge');
const common = require('./webpack.config.common.js');

module.exports = merge(common, {

    mode: 'development',

    devtool: 'eval-source-map',

    module: {
        rules: [
            {
                test: /\.(ts|js)x?$/,
                loader: 'string-replace-loader',
                options: {
                    search: '/* babylonjs-inspector */',
                    replace:
                        `import("@babylonjs/inspector").then(() => {
                            this._scene.debugLayer.show();
                            document.addEventListener('keyup', (e) => {
                                if (e.code === "KeyI") {
                                    this._scene.debugLayer.isVisible() ? this._scene.debugLayer.hide() : this._scene.debugLayer.show();
                                }
                            });
                        });`,
                }
            }
        ]
    },
});
