import { useContext } from 'react'
import IndexContext from '@/pageContexts/indexContext'
import { Box, Container } from '@mui/material'
import { useTheme } from '@mui/material/styles'

const Unary = () => {
    const ctx = useContext(IndexContext)
    const theme = useTheme()

    return (
        <Container sx={{ bgcolor: theme.palette.background.paper }}>
            <pre>{JSON.stringify(ctx.result, null, 2)}</pre>
        </Container>
    )
}

export default Unary
