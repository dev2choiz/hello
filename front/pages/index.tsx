import type { GetServerSidePropsContext, NextPage } from 'next'
import UnaryComp from '@components/UnaryComp'
import { UnaryRequest, UnaryResponse } from '@protobuf/sandbox_pb'
import { SandboxClient } from '@protobuf/sandbox_pb_service'
import config from '@config/config'
import { grpc } from '@improbable-eng/grpc-web'
import { NodeHttpTransport } from '@improbable-eng/grpc-web-node-http-transport'

const Home: NextPage = props => {
    return <UnaryComp result={(props as any).result} />
}

export async function getServerSideProps(ctx: GetServerSidePropsContext) {
    const req = new UnaryRequest()
    const name = (ctx.query.name ?? '') as string
    req.setName(name)
    const opts = {} as grpc.RpcOptions
    if ('undefined' === typeof window) {
        opts.transport = NodeHttpTransport()
        opts.debug = false
    }

    const client = new SandboxClient(config.serverGrpcBaseUrl, opts)
    const result = await new Promise<UnaryResponse.AsObject | null>(resolve => {
        client.unary(req, (error, res) => {
            if (!!error) {
                console.error(error.message)
                resolve(null)
                return
            }
            resolve((res as UnaryResponse).toObject())
        })
    })
    return {
        props: {
            result,
        },
    }
}

export default Home
