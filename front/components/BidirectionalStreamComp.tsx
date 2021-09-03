import { useEffect, useState } from 'react'
import { SandboxClient } from '@protobuf/sandbox_pb_service'
import { UnaryRequest } from '@protobuf/sandbox_pb'
import { Box, Container } from '@mui/material'
import { useTheme } from '@mui/material/styles'

const BidirectionalStream = () => {
    const theme = useTheme()
    const [data, setData] = useState<string>('initial')
    const url = process.env.NEXT_PUBLIC_API_BASE_URL as string
    useEffect(() => {
        const client = new SandboxClient(url)
        const res = client.bidirectionalStream()
        res.on('data', msg => {
            setData(d => d + ' ' + msg.getResponse())
        })

        for (let i = 0; i < 20; i++) {
            const req = new UnaryRequest()
            req.setName('msg' + i)
            res.write(req)
        }

    }, [url])

    return <Container sx={{ bgcolor: theme.palette.background.paper }}>
        <pre>{JSON.stringify(data, null, 2)}</pre>
    </Container>
}

export default BidirectionalStream
