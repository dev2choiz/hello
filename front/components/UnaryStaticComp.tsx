import { memo } from 'react'
import { UnaryResponse } from '@protobuf/sandbox_pb'
import styles from './UnaryComp.module.css'

type Props = {
    result: UnaryResponse.AsObject
}

const UnaryStatic = (props: Props) => {
    return (
        <div className={styles.container}>
            <div>
                <strong>Unary Static page</strong>
            </div>
            <pre>{JSON.stringify(props.result, null, 2)}</pre>
        </div>
    )
}

export default memo(UnaryStatic)
