import { memo, useContext } from 'react'
import styles from './Style.module.css'
import { useRouter } from 'next/router'
import UnaryStaticContext from '@/pageContexts/unaryStaticContext'
import { Box } from '@mui/material'
import { useTheme } from '@mui/material/styles'

const UnaryStatic = () => {
    const router = useRouter()
    const ctx = useContext(UnaryStaticContext)
    const theme = useTheme()

    if (router.isFallback) {
        return <div>Loading...</div>
    }

    return (
        <Box sx={{ bgcolor: theme.palette.primary.main }}>
            <div>
                <strong>Unary Static page</strong>
            </div>
            <div>last generation: <strong>{ctx.dateTime}</strong></div>
            <pre>{JSON.stringify(ctx.result, null, 2)}</pre>
        </Box>
    )
}

export default memo(UnaryStatic)
