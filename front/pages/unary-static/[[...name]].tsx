import type { GetStaticPropsContext, NextPage } from 'next'
import UnaryStaticComp from '@components/UnaryStaticComp'
import { ServerStreamRequest, UnaryRequest, UnaryResponse } from '@protobuf/sandbox_pb'
import {SandboxClient} from '@protobuf/sandbox_pb_service'
import config from '@config/config'
import {grpc} from '@improbable-eng/grpc-web'
import { NodeHttpTransport } from '@improbable-eng/grpc-web-node-http-transport'
import { GetStaticPaths, GetStaticProps } from 'next'

const Home: NextPage = (props) => {
    return <UnaryStaticComp result={(props as any).result} />
}

export const getStaticPaths: GetStaticPaths = async(ctx) => {
    const names = ['rand', 'richard', 'fitz', 'belgarion', 'gerald', '']
    return {
        paths: names.map(name => ({
            params: { name: [name] },
        })),
        fallback: false,
    }
}

export async function getStaticProps(ctx: GetStaticPropsContext) {
    const req = new UnaryRequest()
    let name = ''
    if (!!ctx.params?.name) {
        name = ctx.params?.name[0]
    }

    req.setName(name as string)
    const opts = {} as grpc.RpcOptions
    if ('undefined' === typeof window) {
        opts.transport = NodeHttpTransport()
    }

    const client = new SandboxClient(config.nodeGrpcBaseUrl, opts)
    const result = await new Promise<UnaryResponse.AsObject | null>((resolve) => {
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
            result,
        },
    }
}

export default Home
