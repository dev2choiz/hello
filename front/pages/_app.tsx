import type { AppProps } from 'next/app'
import Link from 'next/link'
import { useEffect, useState } from 'react'
import faker from 'faker'
import '@/styles/globals.css'

const MyApp = ({ Component, pageProps }: AppProps) => {
    const [randomName, setRandomName] = useState<string>('rand')

    useEffect(() => {
        setRandomName(faker.name.lastName())
        const interval = setInterval(() => { setRandomName(faker.name.lastName()) }, 5000)
        return () => { clearInterval(interval) }
    }, [])
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
    </div>
}

export default MyApp
