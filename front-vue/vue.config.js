const path = require('path')
const resolve = dir => path.join(__dirname, dir)

module.exports = {
    chainWebpack: conf => {
        conf.resolve.alias.set('@components', resolve('src/components'))
        conf.resolve.alias.set('@pages', resolve('src/pages'))
    }
}
