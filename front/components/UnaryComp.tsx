import { useEffect, useState } from 'react'
import styles from './UnaryComp.module.css'
import {UnaryRequest, UnaryResponse} from '@protobuf/sandbox_pb'
import { SandboxClient } from '@protobuf/sandbox_pb_service'
import config from '@config/config'

type Props = {
    unaryResult: UnaryResponse.AsObject
}


const Unary = (props: Props) => {
    console.log(props)
    const [data, setData] = useState<any>(props.unaryResult)
    const url = config.grpcBaseUrl

    useEffect(() => {
        if (!!data) return
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
