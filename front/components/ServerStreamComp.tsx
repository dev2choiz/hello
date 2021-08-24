import { useEffect, useState } from 'react'
import styles from './UnaryComp.module.css'
import { ServerStreamRequest } from '@protobuf/sandbox_pb'
import { SandboxClient } from '@protobuf/sandbox_pb_service'
import { useRouter } from 'next/router'
import config from '@config/config'

const ServerStream = () => {
    const router = useRouter()
    const nbResp = !!router.query.number ? parseInt(router.query.number[0] as string) : 15
    const [data, setData] = useState<string>('server responses:')
    const url = config.grpcBaseUrl
    useEffect(() => {
        const client = new SandboxClient(url)
        const req = new ServerStreamRequest()
        req.setNumber(nbResp)
        req.setMsPerResponse(1000)
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
            <div>nb: { nbResp }</div>
            <pre>{JSON.stringify(data, null, 2)}</pre>
        </div>
    )
}

export default ServerStream
