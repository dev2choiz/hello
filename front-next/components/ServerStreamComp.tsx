import { memo, useEffect, useRef, useState } from 'react'
import { ServerStreamRequest } from '@protobuf/sandbox_pb'
import { SandboxClient } from '@protobuf/sandbox_pb_service'
import { useRouter } from 'next/router'
import { Box, Container, LinearProgress, Skeleton } from '@mui/material'
import { useTheme } from '@mui/material/styles'
import getConfig from 'next/config'

const ServerStream = () => {
    const cancelGrpc = useRef<(() => void)>(() => {})
    const router = useRouter()
    const theme = useTheme()
    const nbResp = !!router.query.number ? parseInt(router.query.number[0] as string) : 15
    const [data, setData] = useState<Array<string>>([])
    const [progress, setProgress] = useState(0)
    const { publicRuntimeConfig } = getConfig()
    const url = publicRuntimeConfig.grpcBaseUrl

    useEffect(() => {
        setData([])
        const client = new SandboxClient(url)
        const req = new ServerStreamRequest()
        req.setNumber(nbResp)
        req.setMsPerResponse(1000)
        const res = client.serverStream(req)
        cancelGrpc.current = res.cancel
        res.on('data', msg => {
            setData(d => {
                const cd = [ ...d, msg.getMessage() ]

                setProgress(Math.floor((cd.length / nbResp) * 100))
                return cd
            })
        })
        return () => {
            cancelGrpc.current()
        }
    }, [url, nbResp])

    return (
        <Container sx={{ bgcolor: theme.palette.background.paper }}>
            <div>
                <strong>Server stream gRPC client-side call [SPA]</strong>
            </div>
            <div>nb responses: { nbResp }</div>
            <Box
                display="flex"
                flexWrap="wrap"
                p={1}
                m={1}
            >
                { 0 === data.length
                    ? displaySkeleton(nbResp)
                    : data.map((d, i) => {
                        return <Box p={1} m={1} bgcolor={theme.palette.text.primary} key={i} style={ {
                            color: theme.palette.background.paper,
                            borderRadius: 4,
                        } }>
                            { d }
                        </Box>
                    })
                }
            </Box>
            <LinearProgress variant="determinate" value={progress} />
        </Container>
    )
}

const displaySkeleton = (nb: number) => {
    return <>
        { Array(nb).fill(1).map((_, i) => {
            return <Skeleton
                variant="rectangular"
                width={30}
                height={30}
                key={i}
                style={{
                    margin: 2,
                }}
            />
        }) }
    </>
}

export default memo(ServerStream)
