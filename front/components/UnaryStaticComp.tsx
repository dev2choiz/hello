import { memo } from 'react'
import { UnaryResponse } from '@protobuf/sandbox_pb'
import styles from './Style.module.css'
import { useRouter } from 'next/router'

export type UnaryStaticProps = {
    result: UnaryResponse.AsObject
    dateTime: string
}

const UnaryStatic = (props: UnaryStaticProps) => {
    const router = useRouter()

    if (router.isFallback) {
        return <div>Loading...</div>
    }

    return (
        <div className={styles.container}>
            <div>
                <strong>Unary Static page</strong>
            </div>
            <div>last generation: <strong>{props.dateTime}</strong></div>
            <pre>{JSON.stringify(props.result, null, 2)}</pre>
        </div>
    )
}

export default memo(UnaryStatic)
