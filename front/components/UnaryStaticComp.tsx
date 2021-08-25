import styles from './UnaryComp.module.css'
import { UnaryResponse } from '@protobuf/sandbox_pb'

type Props = {
    result: UnaryResponse.AsObject
}

const Unary = (props: Props) => {
    return (
        <div className={styles.container}>
            <div>
                <strong>Unary Static page</strong>
            </div>
            <pre>{JSON.stringify(props.result, null, 2)}</pre>
        </div>
    )
}

export default Unary
