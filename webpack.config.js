var path = require('path');

module.exports = [
    {
        mode: 'development',
        entry: {
            app: ['./src/app.js', './cmd/js/wasm_exec.js', './cmd/js/go_run.js'],
        },
        module: {
            rules: [
                {
                    test: /\.(js|jsx)$/,
                    exclude: /node_modules/,
                    use: {
                        loader: "babel-loader",
                        options: {
                            presets:[
                                ['@babel/preset-react'],
                                ['@babel/preset-env',
                                    {
                                        useBuiltIns: "usage",
                                        corejs: "3.2.1"
                                    },
                                ]
                            ],
                        }
                    }
                }
            ]
        },
        output: {
            path: path.resolve(__dirname, 'dist/dev/js'),
            filename: 'bundle.js'
        }
    },
    {
        mode: 'production',
        entry: {
            app: ['./src/app.js', './cmd/js/wasm_exec.js', './cmd/js/go_run.js'],
        },
        module: {
            rules: [
                {
                    test: /\.(js|jsx)$/,
                    exclude: /node_modules/,
                    use: {
                        loader: "babel-loader",
                        options: {
                            presets:[
                                ['@babel/preset-react'],
                                ['@babel/preset-env',
                                    {
                                        useBuiltIns: "usage",
                                        corejs: "3.2.1"
                                    },
                                ]
                            ],
                        }
                    }
                }
            ],
        },
        output: {
            path: path.resolve(__dirname, 'dist/deploy/js'),
            filename: 'bundle.js'
        }
    }
];