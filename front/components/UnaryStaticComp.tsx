import { memo, useContext } from 'react'
import styles from './Style.module.css'
import { useRouter } from 'next/router'
import UnaryStaticContext from '@/pageContexts/unaryStaticContext'

const UnaryStatic = () => {
    const router = useRouter()
    const ctx = useContext(UnaryStaticContext)

    if (router.isFallback) {
        return <div>Loading...</div>
    }

    return (
        <div className={styles.container}>
            <div>
                <strong>Unary Static page</strong>
            </div>
            <div>last generation: <strong>{ctx.dateTime}</strong></div>
            <pre>{JSON.stringify(ctx.result, null, 2)}</pre>
        </div>
    )
}

export default memo(UnaryStatic)
