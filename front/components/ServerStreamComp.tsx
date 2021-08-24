import { useEffect, useState } from 'react'
import styles from './UnaryComp.module.css'
import { ServerStreamRequest } from '@protobuf/sandbox_pb'
import { SandboxClient } from '@protobuf/sandbox_pb_service'
import { useRouter } from 'next/router'

const ServerStream = () => {
    const router = useRouter()
    const nbResp = !!router.query.nb ? parseInt(router.query.nb as string) : 15
    const [data, setData] = useState<string>('initial')
    const url = process.env.NEXT_PUBLIC_API_BASE_URL as string
    useEffect(() => {
        const client = new SandboxClient(url)
        const req = new ServerStreamRequest()
        req.setNumber(nbResp)
        const res = client.serverStream(req)
        res.on('data', msg => {
            setData(d => d + ' ' + msg.getMessage())
        })
        return () => {
            res.cancel()
        }
    }, [url])

    return (
        <div className={styles.container}>
            <pre>{JSON.stringify(data, null, 2)}</pre>
        </div>
    )
}

export default ServerStream
