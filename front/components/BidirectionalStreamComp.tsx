import styles from './UnaryComp.module.css'
import { useEffect, useState } from 'react'
import { SandboxClient } from '@protobuf/sandbox_pb_service'
import { UnaryRequest } from '@protobuf/sandbox_pb'

const BidirectionalStream = () => {
    const [data, setData] = useState<string>('initial')
    const url = process.env.NEXT_PUBLIC_API_BASE_URL as string
    useEffect(() => {
        const client = new SandboxClient(url)
        const res = client.bidirectionalStream()
        res.on('data', msg => {
            setData(data + ' ' + msg.getResponse())
        })

        for (let i = 0; i < 20; i++) {
            const req = new UnaryRequest()
            req.setMessage('msg' + i)
            res.write(req)
        }

    }, [])

    return <div className={styles.container}>
        <pre>{JSON.stringify(data, null, 2)}</pre>
    </div>
}

export default BidirectionalStream
