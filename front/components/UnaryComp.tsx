import styles from './Style.module.css'
import { useContext } from 'react'
import IndexContext from '@/pageContexts/indexContext'

const Unary = () => {
    const ctx = useContext(IndexContext)
    return (
        <div className={styles.container}>
            <pre>{JSON.stringify(ctx.result, null, 2)}</pre>
        </div>
    )
}

export default Unary
