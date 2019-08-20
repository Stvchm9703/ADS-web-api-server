module.exports = {
  root: true,
  env: {
    browser: true,
    node: true
  },
  parserOptions: {
    parser: 'babel-eslint'
  },
  extends: [
    '@nuxtjs',
    'plugin:nuxt/recommended'
  ],
  // add your custom rules here
  rules: {
    'no-console': process.env.NODE_ENV === 'production' ? 'error' : 'off',
        'no-debugger': process.env.NODE_ENV === 'production' ? 'error' : 'off',
        'semi': [1, 'always'],
        "indent": ["error", 4, {
            "outerIIFEBody": 0,
            "MemberExpression": 1,
            "FunctionDeclaration": {
                "parameters": "first"
            },
            "CallExpression": {
                "arguments": "first"
            },
            "ObjectExpression": 1,
            "flatTernaryExpressions": false,
            "SwitchCase": 1

        }],
        'array-element-newline': ["error", "never"],
        'compat/compat': "error"
  }
}
