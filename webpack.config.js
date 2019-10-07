var path = require('path');

module.exports = [
    {
        mode: 'development',
        entry: {
            app: ['./cmd/js/wasm_exec.js', './cmd/js/go_run.js'],
        },
        module: {
            rules: [
                {
                    test: /\.(js|jsx)$/,
                    exclude: /node_modules/,
                    use: {
                        loader: "babel-loader"
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
            app: ['./cmd/js/wasm_exec.js', './cmd/js/go_run.js'],
        },
        module: {
            rules: [
                {
                    test: /\.(js|jsx)$/,
                    exclude: /node_modules/,
                    use: {
                        loader: "babel-loader"
                    }
                }
            ]
        },
        output: {
            path: path.resolve(__dirname, 'dist/deploy/js'),
            filename: 'bundle.js'
        }
    }
];