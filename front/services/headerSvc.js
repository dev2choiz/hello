const getCacheControl = (ttl, public = false) => {
    const pub = public ? 'public,' : ''
    return ttl ?
        `${pub}s-maxage=${ttl},stale-while-revalidate`
        : 'no-cache'
}

const getCacheControlObj = (ttl) => ({ key: 'Cache-Control', value: getCacheControl(ttl) })

const exp = {
    getCacheControl,

    generateConfig: () => [
        // SSR pages ==> handled in getServerSideProps() functions
        // SSG/ISR pages ==> handled in getServerStaticProps() functions
        // Automatic optimized pages
        { source: '/server-stream(.*)', headers: [ getCacheControlObj(3600), ], },
        { source: '/client-stream(.*)', headers: [ getCacheControlObj(3600), ], },
        { source: '/bidirectional-stream(.*)', headers: [ getCacheControlObj(3600), ], },
    ]
}

module.exports = exp
