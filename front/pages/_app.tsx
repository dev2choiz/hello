import Head from 'next/head'
import Link from 'next/link'
import type { AppProps } from 'next/app'
import { useEffect, useState } from 'react'
import faker from 'faker'
import { CacheProvider, EmotionCache } from '@emotion/react'
import createEmotionCache from '@/styles/createEmotionCache'
import { ThemeProvider } from '@mui/material/styles'
import CssBaseline from '@mui/material/CssBaseline'
import '@/styles/globals.css'
import theme from '@/styles/theme'

const clientSideEmotionCache = createEmotionCache()

interface MyAppProps extends AppProps {
    emotionCache?: EmotionCache;
}

const App = ({ Component, pageProps/*, emotionCache = clientSideEmotionCache*/ }: MyAppProps) => {
    const [randomName, setRandomName] = useState<string>('rand')
    const emotionCache = clientSideEmotionCache

    useEffect(() => {
        setRandomName(faker.name.lastName())
        const interval = setInterval(() => { setRandomName(faker.name.lastName()) }, 5000)
        return () => { clearInterval(interval) }
    }, [])

    return <CacheProvider value={emotionCache}>
        <Head>
            <title>HelloFront</title>
            <meta name="viewport" content="initial-scale=1, width=device-width" />
        </Head>
        <ThemeProvider theme={theme}>
            {/* CssBaseline kickstart an elegant, consistent, and simple baseline to build upon. */}
            <CssBaseline />
            <div>
                <table><tbody>
                    <tr>
                        <td>unary ssr</td>
                        <td><Link prefetch={false} href={{pathname: '/'}}><a>/</a></Link><br/></td>
                    </tr>
                    <tr>
                        <td>unary ssr with param</td>
                        <td><Link prefetch={false} href={{ pathname: '/', query: { name: 'John' }}}><a>/?name=John</a></Link><br/></td>
                    </tr>
                    <tr>
                        <td>unary static</td>
                        <td><Link prefetch={false} href="/unary-static"><a>/unary-static</a></Link><br/></td>
                    </tr>
                    <tr>
                        <td>unary static with param</td>
                        <td><Link prefetch={false} href={`/unary-static/${randomName}`}><a>/unary-static/{randomName}</a></Link><br/></td>
                    </tr>
                    <tr>
                        <td>server stream</td>
                        <td><Link prefetch={false} href="/server-stream"><a>/server-stream</a></Link><br/></td>
                    </tr>
                    <tr>
                        <td>server stream with param</td>
                        <td><Link prefetch={false} href="/server-stream/5"><a>/server-stream/5</a></Link><br/></td>
                    </tr>
                </tbody></table>
            </div>
            <Component {...pageProps} />
        </ThemeProvider>
    </CacheProvider>
}

export default App
