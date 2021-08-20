module.exports = {
    extends: [
        'next',
        'next/core-web-vitals',
        'plugin:@typescript-eslint/recommended',
        'prettier',
    ],
    plugins: [
        'react',
        'react-hooks',
        '@typescript-eslint',
    ],
    rules: {
        'react-hooks/rules-of-hooks': 'error',
        indent: ['error', 4, { SwitchCase: 1 }],
        'react/jsx-indent': ['error', 4],
        yoda: 0,
        quotes: ['error', 'single'],
        semi: ['error', 'never'],
    }
}
