import '@/styles/globals.css'
import type { AppProps } from 'next/app'
import Link from 'next/link'

const MyApp = ({ Component, pageProps }: AppProps) => {
    return <div>
        <div>
            <Link href="/"><a>unary</a></Link><br/>
            <Link href="/unary-static"><a>unary static</a></Link><br/>
            <Link href="/server-stream"><a>serveur stream</a></Link>
        </div>
        <Component {...pageProps} />
    </div>
}

export default MyApp
