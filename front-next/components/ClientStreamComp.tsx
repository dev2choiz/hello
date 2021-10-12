import { useEffect } from 'react'
import { UnaryRequest } from '@protobuf/sandbox_pb'
import { SandboxClient } from '@protobuf/sandbox_pb_service'
import { Container } from '@mui/material'
import { useTheme } from '@mui/material/styles'

const ClientStream = () => {
    const url = process.env.NEXT_PUBLIC_API_BASE_URL as string
    const theme = useTheme()

    useEffect(() => {
        const client = new SandboxClient(url)
        const reqStream = client.clientStream()
        for (let i = 1; i < 10; i++) {
            setTimeout(() => {
                if (9 === i) {
                    reqStream.end()
                    return
                }
                const req = new UnaryRequest()
                req.setName('msg' + i)
                reqStream.write(req)
            }, 1000 * i)
        }
    }, [url])

    return (
        <Container sx={{ bgcolor: theme.palette.background.paper }}>
            <pre>streaming to server</pre>
        </Container>
    )
}

export default ClientStream
