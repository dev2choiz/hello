module.exports = {
    root: true,
    env: {
        node: true
    },
    extends: [
        'plugin:vue/vue3-essential',
        '@vue/standard',
        '@vue/typescript/recommended'
    ],
    parserOptions: {
        ecmaVersion: 2020
    },
    rules: {
        'no-console': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
        'no-debugger': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
        indent: [ 'error', 4, { SwitchCase: 1 } ],
        yoda: 0,
        quotes: [ 'error', 'single' ],
        semi: [ 'error', 'never' ],
        "comma-dangle": 0,
        "spaced-comment": 0,
        '@typescript-eslint/no-empty-function': 0,
    },
    ignorePatterns: ['./src/protobuf/**',]
}
