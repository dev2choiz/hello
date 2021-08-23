import { useEffect, useState } from 'react'
import styles from './UnaryComp.module.css'
import { UnaryRequest } from '@protobuf/sandbox_pb'
import { SandboxClient } from '@protobuf/sandbox_pb_service'

const Unary = () => {
    const [data, setData] = useState<any>(null)
    const url = process.env.NEXT_PUBLIC_API_BASE_URL as string

    useEffect(() => {
        const req = new UnaryRequest()
        const client = new SandboxClient(url)
        client.unary(req, (error, res) => {
            if (!!error) {
                console.log(error.message)
            }
            setData(res?.toObject())
        })
    }, [url])

    return (
        <div className={styles.container}>
            <pre>{JSON.stringify(data, null, 2)}</pre>
        </div>
    )
}

export default Unary
