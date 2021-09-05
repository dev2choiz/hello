import { memo, useEffect, useRef, useState } from 'react'
import { ServerStreamRequest } from '@protobuf/sandbox_pb'
import { SandboxClient } from '@protobuf/sandbox_pb_service'
import { useRouter } from 'next/router'
import config from '@config/config'
import { Box, Container } from '@mui/material'
import { useTheme } from '@mui/material/styles'

const ServerStream = () => {
    const cancelGrpc = useRef<(() => void)>(() => {})
    const router = useRouter()
    const theme = useTheme()
    const nbResp = !!router.query.number ? parseInt(router.query.number[0] as string) : 15
    const [data, setData] = useState<Array<string>>([])
    const url = config.grpcBaseUrl

    useEffect(() => {
        const client = new SandboxClient(url)
        const req = new ServerStreamRequest()
        req.setNumber(nbResp)
        req.setMsPerResponse(1000)
        const res = client.serverStream(req)
        cancelGrpc.current = res.cancel
        res.on('data', msg => {
            setData(d => {
                const cd = [ ...d ]
                cd.push(msg.getMessage())
                return cd
            })
        })
        return () => {
            cancelGrpc.current()
        }
    }, [url, nbResp])

    return (
        <Container sx={{ bgcolor: theme.palette.background.paper }}>
            <div>nb requested: { nbResp }</div>
            <Box
                display="flex"
                flexWrap="wrap"
                p={1}
                m={1}
            >
                { data.map((d, i) => {
                    return <Box p={1} m={1} bgcolor={theme.palette.text.primary} key={i} style={ {
                        color: theme.palette.background.paper,
                        borderRadius: 4,
                    } }>
                        { d }
                    </Box>
                })}
            </Box>
        </Container>
    )
}

export default memo(ServerStream)
