const purgecss = require('@fullhuman/postcss-purgecss')({
    content: [
        './html/**/*.html'
    ],

    defaultExtractor: content => content.match(/[\w-/:]+(?<!:)/g) || []
})

module.exports = {
    plugins: [
        require('tailwindcss'),
        // require('autoprefixer'),
        require('postcss-preset-env')({ stage: 1 }),
        ...process.env.NODE_ENV === 'production'
        ? [purgecss]
        : []
    ]
}
