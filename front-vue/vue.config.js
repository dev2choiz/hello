const path = require('path')
const resolve = dir => path.join(__dirname, dir)

/**
 * @type {import('@vue/cli-service').ProjectOptions}
 */
module.exports = {
    chainWebpack: conf => {
        conf.resolve.alias.set('@components', resolve('src/components'))
        conf.resolve.alias.set('@pages', resolve('src/pages'))
        conf.resolve.alias.set('@protobuf', resolve('src/protobuf'))
        conf.resolve.alias.set('@services', resolve('src/services'))
    },
}
