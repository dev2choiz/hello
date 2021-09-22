import { memo, useContext } from 'react'
import { useRouter } from 'next/router'
import UnaryStaticContext from '@/pageContexts/unaryStaticContext'
import { CircularProgress, Container } from '@mui/material'
import { useTheme } from '@mui/material/styles'

const UnaryStatic = () => {
    const router = useRouter()
    const ctx = useContext(UnaryStaticContext)
    const theme = useTheme()

    if (router.isFallback) {
        return <Container sx={{ bgcolor: theme.palette.background.paper }}><CircularProgress /></Container>
    }

    return (
        <Container sx={{ bgcolor: theme.palette.background.paper }}>
            <div>
                <strong>Unary Static page</strong>
            </div>
            <div>last generation: <strong>{ctx.dateTime}</strong></div>
            <div>Generation interval: <strong>{ctx.revalidate}</strong> seconds</div>
            <pre>{JSON.stringify(ctx.result, null, 2)}</pre>
        </Container>
    )
}

export default memo(UnaryStatic)
