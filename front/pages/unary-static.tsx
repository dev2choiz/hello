import type { GetStaticPropsContext, NextPage } from 'next'
import UnaryStaticComp from '@components/UnaryStaticComp'
import { UnaryRequest, UnaryResponse } from '@protobuf/sandbox_pb'
import {SandboxClient} from '@protobuf/sandbox_pb_service'
import config from '@config/config'
import {grpc} from '@improbable-eng/grpc-web'
import { NodeHttpTransport } from '@improbable-eng/grpc-web-node-http-transport'

const Home: NextPage = (props) => {
    return <UnaryStaticComp unaryResult={(props as any).unaryResult} />
}

export async function getStaticProps(ctx: GetStaticPropsContext) {
    const req = new UnaryRequest()
    req.setMessage('hello')
    const opts = {} as grpc.RpcOptions
    if ('undefined' === typeof window) {
        opts.transport = NodeHttpTransport()
    }

    const client = new SandboxClient(config.nodeGrpcBaseUrl, opts)
    const unaryResult = await new Promise<UnaryResponse.AsObject | null>((resolve) => {
        client.unary(req, (error, res) => {
            if (!!error) {
                console.log(error.message)
                resolve(null)
                return
            }
            resolve((res as UnaryResponse).toObject())
        })
    })
    return {
        props: {
            unaryResult,
        },
    }
}


export default Home
