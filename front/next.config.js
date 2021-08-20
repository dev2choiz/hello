/** @type {import('next').NextConfig} */

// eslint-disable-next-line @typescript-eslint/no-var-requires
const path = require('path')
const resolve = p => path.join(__dirname, p)

module.exports = {
    reactStrictMode: true,
    webpack(conf, options) {
        conf.resolve.alias['@components'] = resolve('components')
        conf.resolve.alias['@pages'] = resolve('pages')
        conf.resolve.alias['@protobuf'] = resolve('protobuf')
        conf.resolve.alias['@'] = resolve('./')
        return conf
    },
}
