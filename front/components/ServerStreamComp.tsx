import { useEffect, useState } from 'react'
import styles from './UnaryComp.module.css'
import { UnaryRequest } from '@protobuf/sandbox_pb'
import { SandboxClient } from '@protobuf/sandbox_pb_service'

const ServerStream = () => {
    const [data, setData] = useState<string>('initial')
    const url = process.env.NEXT_PUBLIC_API_BASE_URL as string
    useEffect(() => {
        const client = new SandboxClient(url)
        const res = client.serverStream(new UnaryRequest())
        res.on('data', msg => {
            setData(d => d + ' ' + msg.getResponse())
        })
    }, [url])

    return (
        <div className={styles.container}>
            <pre>{JSON.stringify(data, null, 2)}</pre>
        </div>
    )
}

export default ServerStream
