/** @type {import('next').NextConfig} */

/* eslint @typescript-eslint/no-var-requires: 0 */
const path = require('path')
const resolve = p => path.join(__dirname, p)

module.exports = {
    reactStrictMode: true,
    assetPrefix: 'production' === process.env.NODE_ENV ? process.env.CDN_BASE_URL : '',

    webpack(conf) {
        conf.resolve.alias['@components'] = resolve('components')
        conf.resolve.alias['@config'] = resolve('config')
        conf.resolve.alias['@pages'] = resolve('pages')
        conf.resolve.alias['@protobuf'] = resolve('protobuf')
        conf.resolve.alias['@'] = resolve('./')
        return conf
    },

    publicRuntimeConfig: {
        cdnBaseUrl: process.env.CDN_BASE_URL,
        grpcBaseUrl: process.env.BROWSER_API_BASE_URL,
    },

    serverRuntimeConfig: {
        apiKey: process.env.GRPC_API_KEY,
        serverGrpcBaseUrl: process.env.SERVER_API_BASE_URL,
    },

    /*exportPathMap: async function () {
        const ret = {}
        const names = ['rand', 'richard', 'fitz', 'belgarion', 'gerald', '']
        names.forEach((name) => {
            ret['/unary-static/' + name] = { page: '/unary-static/[[...name]]', query: { name } }
        })
        ret['/bidirectional-stream'] = { page: '/bidirectional-stream' }
        ret['/client-stream'] = { page: '/client-stream' }
        for (let i = 1; i <= 20; i++) {
            ret['/server-stream/' + i] = { page: '/server-stream/[[...number]]', query: { number: i } }
        }
        return ret
    },*/
}
