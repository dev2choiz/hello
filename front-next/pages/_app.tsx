import Head from 'next/head'
import type { AppProps } from 'next/app'
import { CacheProvider, EmotionCache } from '@emotion/react'
import createEmotionCache from '@/styles/createEmotionCache'
import '@/styles/globals.css'
import Layout from '@components/Layout'

const clientSideEmotionCache = createEmotionCache()

interface MyAppProps extends AppProps {
    emotionCache?: EmotionCache;
}

const App = ({ Component, pageProps }: MyAppProps) => {
    const emotionCache = clientSideEmotionCache

    return <CacheProvider value={emotionCache}>
        <Head>
            <title>HelloFront</title>
            <meta name="viewport" content="initial-scale=1, width=device-width" />
        </Head>
        <Layout>
            <Component {...pageProps} />
        </Layout>
    </CacheProvider>
}

export default App
