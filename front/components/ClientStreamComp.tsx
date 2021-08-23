import { useEffect } from 'react'
import styles from './UnaryComp.module.css'
import { UnaryRequest } from '@protobuf/sandbox_pb'
import { SandboxClient } from '@protobuf/sandbox_pb_service'

const ClientStream = () => {
    const url = process.env.NEXT_PUBLIC_API_BASE_URL as string
    useEffect(() => {
        const client = new SandboxClient(url)
        const reqStream = client.clientStream()
        for (let i = 1; i < 10; i++) {
            setTimeout(() => {
                if (9 === i) {
                    reqStream.end()
                    return
                }
                const req = new UnaryRequest()
                req.setMessage('msg' + i)
                reqStream.write(req)
            }, 5000 * i)
        }
    }, [url])

    return (
        <div className={styles.container}>
            <pre>streaming to server</pre>
        </div>
    )
}

export default ClientStream
