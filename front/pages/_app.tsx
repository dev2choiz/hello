import '@/styles/globals.css'
import type { AppProps } from 'next/app'
import Link from 'next/link'

const MyApp = ({ Component, pageProps }: AppProps) => {
    return <div>
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
                    <td><Link prefetch={false} href="/unary-static/rand"><a>/unary-static/rand</a></Link><br/></td>
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
    </div>
}

export default MyApp
