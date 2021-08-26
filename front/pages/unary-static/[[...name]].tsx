import type { GetStaticProps, GetStaticPropsContext, GetStaticPaths, NextPage } from 'next'
import UnaryStaticComp, { UnaryStaticProps } from '@components/UnaryStaticComp'
import { UnaryRequest, UnaryResponse } from '@protobuf/sandbox_pb'
import { SandboxClient } from '@protobuf/sandbox_pb_service'
import config from '@config/config'
import { grpc } from '@improbable-eng/grpc-web'
import { NodeHttpTransport } from '@improbable-eng/grpc-web-node-http-transport'

const UnaryStaticPage: NextPage<UnaryStaticProps> = props => {
    return <UnaryStaticComp { ...(props) } />
}

export const getStaticPaths: GetStaticPaths = async() => {
    return {
        paths: config.unaryStaticParams.map(name => ({
            params: { name: [name] },
        })),
        fallback: true,
    }
}

const timeout = (t = 2000) => {
    console.log(`\nsleep of ${t}ms\n`)
    return new Promise(resolve => setTimeout(resolve, t))
}

export const getStaticProps: GetStaticProps = async(ctx: GetStaticPropsContext) => {
    let name = ''
    if (!!ctx.params?.name) {
        name = ctx.params?.name[0]
    }
    console.log(`\ngenerate unary-static for ${name}\n`)

    const opts = {} as grpc.RpcOptions
    if ('undefined' === typeof window) {
        opts.transport = NodeHttpTransport()
    }

    const req = new UnaryRequest()
    req.setName(name as string)
    const client = new SandboxClient(config.serverGrpcBaseUrl, opts)
    const result = await new Promise<UnaryResponse.AsObject | null>(resolve => {
        client.unary(req, (error, res) => {
            if (!!error) {
                console.log(error.message)
                resolve(null)
                return
            }
            resolve((res as UnaryResponse).toObject())
        })
    })

    //if (!config.unaryStaticParams.includes(name)) await timeout(4000) // simulate waiting

    return {
        props: { result, dateTime: new Date().toLocaleString() },
        revalidate: 15,
    }
}

export default UnaryStaticPage
