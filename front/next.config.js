/** @type {import('next').NextConfig} */

// eslint-disable-next-line @typescript-eslint/no-var-requires
const path = require('path')
const resolve = p => path.join(__dirname, p)

module.exports = {
    reactStrictMode: true,
    webpack(conf) {
        conf.resolve.alias['@components'] = resolve('components')
        conf.resolve.alias['@config'] = resolve('config')
        conf.resolve.alias['@pages'] = resolve('pages')
        conf.resolve.alias['@protobuf'] = resolve('protobuf')
        conf.resolve.alias['@'] = resolve('./')
        return conf
    },

    /*exportPathMap: async function (
        defaultPathMap,
        { dev, dir, outDir, distDir, buildId }
    ) {
        return {
            '/unary-static': { page: '/unary-static' },
            '/stream-server/10': { page: '/stream-server', query: { number: 10 } },
            '/stream-server/15': { page: '/stream-server', query: { number: 15 } },
            '/stream-server/20': { page: '/stream-server', query: { number: 20 } },
        }
    },*/
}
