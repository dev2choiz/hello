import { memo, useContext } from 'react'
import { useRouter } from 'next/router'
import { CircularProgress, Container } from '@mui/material'
import { useTheme } from '@mui/material/styles'
import HealthContext from '@/pageContexts/healthContext'

const Health = () => {
    const router = useRouter()
    const ctx = useContext(HealthContext)
    const theme = useTheme()

    if (router.isFallback) {
        return <Container sx={{ bgcolor: theme.palette.background.paper }}><CircularProgress /></Container>
    }

    return (
        <Container sx={{ bgcolor: theme.palette.background.paper }}>
            <div>
                <strong>Health page</strong>
            </div>
            <div>last generation: <strong>{ctx.dateTime}</strong></div>
            <div>regenerate each <strong>{ctx.revalidate}</strong> seconds</div>
            <pre>{JSON.stringify(ctx.result, null, 2)}</pre>
        </Container>
    )
}

export default memo(Health)
