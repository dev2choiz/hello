/* eslint @typescript-eslint/no-var-requires: 0 */
/* eslint @typescript-eslint/no-unused-vars: 0 */
const express = require('express')
const app = express()
const cors = require('cors')
app.use(cors())

const argv = require('minimist')(process.argv.slice(2))

if (!argv['port']) throw new Error('missing --port argument')
if (!argv['path-mapping']) throw new Error('missing --path-mapping argument')
const port = argv['port']
const mapping = argv['path-mapping']

app.use(function(req, res, next) {
    console.log(`${req.url}`)
    next()
})

mapping
    .split(';')
    .map(map => map.split(':'))
    .forEach(v => {
        console.log(`${v[0]} ==> ${v[1]}`)
        app.use(v[0], express.static(v[1]))
    })

app.listen(port, () => {
    console.log(`static server listening at :${port}`)
})
