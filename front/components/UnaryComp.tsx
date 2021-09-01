import styles from './Style.module.css'
import { UnaryResponse } from '@protobuf/sandbox_pb'

type Props = {
    result: UnaryResponse.AsObject | null
}

const Unary = (props: Props) => {
    return (
        <div className={styles.container}>
            <pre>{JSON.stringify(props.result, null, 2)}</pre>
        </div>
    )
}

export default Unary
